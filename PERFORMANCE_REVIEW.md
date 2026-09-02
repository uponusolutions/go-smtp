# Performance-Review go-smtp

Datum: 2026-09-02
Stand: Commit `2f9d8a6` (main)
Methode: Code-Review der Hot Paths plus Benchmark-Lauf mit CPU- und Allokations-Profil.

## Zusammenfassung

Die Bibliothek ist bereits stark auf Performance optimiert (Zero-Copy-Peek im
DotReader, Flush-Heuristik für Pipelining im Server, wiederverwendeter
Chunking-Buffer im Client). Trotzdem gibt es klares Potential:

1. Rund ein Viertel aller Allokationen im Send-Pfad stammt aus Debug-Logging,
   das nie ausgegeben wird.
2. Der Client nutzt kein PIPELINING, obwohl der Server es advertised – bei
   echten Netzwerk-Latenzen der größte Hebel.
3. Nebenbefund: ein echter Korrektheitsbug in `server/conn.go` (`handleData`),
   der zu Protokoll-Desync führen kann.

## Messbasis

```
go test -bench='Benchmark/SmallWithChunkingSameConnection$' -benchtime=2s
cpu: AMD Ryzen AI 9 HX 370
Benchmark/SmallWithChunkingSameConnection-24    36722    66437 ns/op    53.19 MB/s
```

- CPU-Profil: 55 % der Samples in Syscalls (Netzwerk-Reads/-Writes,
  `bufio.Flush` mit 43 % kumuliert) – auf Loopback limitieren die Roundtrips,
  nicht die CPU.
- Alloc-Profil (alloc_objects): größter Einzelposten ist
  `server.(*Conn).writeResponse` mit 18,85 % aller Objekte, davon 677k von der
  Debug-Log-Zeile allein.

## Korrektheitsbug: Variable-Shadowing in handleData

`server/conn.go:804` – in der `rstart`-Closure wird das äußere `r` durch `:=`
verschattet:

```go
var r io.Reader
rstart := func() io.Reader {
    if r != nil { return r }
    c.writeResponse(354, ...)
    r := textsmtp.NewDotReader(...)  // shadowing! äußeres r bleibt nil
    return r
}
```

Folgen:

- Der Drain `io.Copy(io.Discard, r)` nach erfolgreichem `session.Data`
  (`conn.go:826`) läuft nie, ebenso wenig der Drain im Fehlerpfad
  (`conn.go:819`). Liest eine Session den Reader nicht bis EOF, bleiben
  Mail-Daten im Buffer und werden anschließend als SMTP-Kommandos
  interpretiert (Protokoll-Desync).
- Ein zweiter `rstart()`-Aufruf sendet ein zweites 354 und erzeugt einen
  zweiten Reader.

Fix: `r = textsmtp.NewDotReader(...)` (Zuweisung statt Deklaration).

## Performance-Findings (nach Impact sortiert)

### 1. Debug-Logging in den Hot Paths (~27 % aller Allokationen)

`server/conn.go:909` (`writeResponse`) und `server/conn.go:968` (`readLine`)
rufen bei jeder Zeile `c.logger().DebugContext(...)` auf. Auch bei
deaktiviertem Debug-Level werden die slog-Attribute eager gebaut –
`slog.Any("enhCode", enhCode)` und `slog.Any("text", ...)` boxen ihre Werte,
und `c.logger()` ruft jedes Mal `session.Logger(ctx)` auf.

Profil: 677k Objekte an `conn.go:909`, 251k an `conn.go:968`.

Empfehlung: Logger einmal pro Connection cachen und die Debug-Aufrufe mit
`logger.Enabled(ctx, slog.LevelDebug)` guarden. Praktisch gratis, größter
Einzelgewinn.

### 2. Client nutzt PIPELINING nicht

Pro Mail fallen client-seitig ~4–5 Roundtrips an (MAIL → warten, je RCPT →
warten, DATA/BDAT → warten), obwohl der Server PIPELINING advertised und die
server-seitige Flush-Heuristik (`conn.go:948`) dafür bereits existiert.

Empfehlung: MAIL + alle RCPTs + DATA/BDAT in einem Flush senden und die
Antworten gesammelt lesen. Halbiert die Latenz pro Mail; bei echten
Netzwerk-RTTs (nicht Loopback) mit Abstand der größte Hebel. Größerer Umbau
in `client/client.go` und `mailer/mailer.go` (`prepare`), aber lohnend – und
der einzige Punkt, der die 66 µs/op strukturell drückt.

### 3. parse.Cmd uppercased die ganze Zeile

`internal/parse/parse.go:23`:

```go
case strings.HasPrefix(strings.ToUpper(line), "STARTTLS"):
```

Allokiert bei jeder Zeile mit Kleinbuchstaben eine komplette Kopie – auch bei
langen `MAIL FROM:`-Zeilen, nur um ein 8-Zeichen-Präfix zu prüfen.
`strings.ToUpper` ist im CPU-Profil sichtbar (1,55 %).

Empfehlung: das eigene `CutPrefixFold` bzw. `strings.EqualFold` auf die ersten
8 Bytes verwenden. Nebenbei: `strings.ToUpper(cmd)` in `server/conn.go:101`
ist redundant, `parse.Cmd` liefert bereits Uppercase.

### 4. Antworten laufen durch fmt-Reflection

`Textproto.PrintfLine` nutzt `fmt.Fprintf`; `conn.go:942`
(`"%d %v.%v.%v %v"`) erzeugt 163k Objekte im Profil. Im Client wird doppelt
formatiert: erst per `strings.Builder` gebaut, dann durch
`c.cmd(250, "%s", sb.String())` nochmal durch fmt geschleust.

Empfehlung: Statuszeilen manuell mit `strconv.AppendInt` / direkten
bufio-Writes bauen; auf `Textproto` eine `WriteLine(string)`-Methode ergänzen
und im Client statt `"%s"`-Format verwenden.

### 5. sb.Grow(2048) pro Kommando im Client

`client/client.go` – `Mail`, `Rcpt` und `Verify` reservieren pro Kommando
2 KB (`Client.Rcpt` + `Client.Mail` zusammen ~9 % der Objekte).

Empfehlung: realistische Größe (~128–256 B) reservieren oder einen Buffer am
Client wiederverwenden.

### 6. parse.Args: Map plus Split pro MAIL/RCPT

`internal/parse/parse.go:52` allokiert pro MAIL/RCPT eine Map und pro Argument
ein Slice via `strings.Split` (~5 % der Objekte). Typische Kommandos haben
0–2 Argumente.

Empfehlung: `strings.Cut` statt `Split`; Map durch kleines Slice oder
Callback-Iteration ersetzen.

### 7. Kleinere Punkte mit gutem Aufwand/Nutzen-Verhältnis

- `smtp.Timeout` (`timeout.go:9`) allokiert pro Client-Kommando eine Closure
  (~4 % der Objekte). Die zwei `SetDeadline`-Aufrufe inline machen.
- `fmt.Sprintf("Roger, accepting mail from <%v>", from)` bei jedem MAIL/RCPT
  (`conn.go:573`, `conn.go:643`): Adresse nicht zurück-echoen (RFC verlangt
  das nicht) und statische `smtp.Status`-Objekte verwenden wie bei
  `smtp.Noop`.
- `NewBdatWriter` wird pro Mail neu allokiert (~2 %) – ließe sich wie der
  `chunkingBuffer` am Client wiederverwenden.
- `Textproto.ReadResponse` (`textproto.go:143`): `message += "\n" + ...` ist
  O(n²) – nur bei mehrzeiligen Antworten (EHLO) relevant, ein
  `strings.Builder` kostet aber nichts.

## Robustheits-Kandidat: dotReader-Panic bei abgerissener Verbindung

`internal/textsmtp/dotreader.go:91` – im `stateCR`-Zweig wird fest auf `c[3]`
und `c[4]` zugegriffen. Kappt der Peer die Verbindung genau an dieser Stelle,
kann `Peek` weniger als 5 Bytes liefern und der Zugriff panict. Der Server
fängt das per recover ab, sauber ist es aber nicht. Einen gezielten Test
schreiben und die Länge prüfen.

## Empfohlene Reihenfolge

| Prio | Maßnahme | Aufwand | Wirkung |
|------|----------|---------|---------|
| 1 | Shadowing-Bug in `handleData` fixen | trivial | Korrektheit |
| 2 | Debug-Logging guarden + Logger cachen | klein | ~27 % weniger Allokationen |
| 3 | `parse.Cmd`-ToUpper, `parse.Args`, `sb.Grow`, `Timeout`-Closure, statische Status | klein | zusammen grob die Hälfte der Allokationen im Send-Pfad |
| 4 | fmt durch manuelles Formatieren ersetzen, `WriteLine` | mittel | CPU + Allokationen |
| 5 | Client-PIPELINING | groß | Roundtrips pro Mail ~halbiert; einziger struktureller Hebel auf die 66 µs/op |
| 6 | dotReader-Grenzfall testen/absichern | klein | Robustheit |

Nach den Punkten 1–4 limitieren auf Loopback weiterhin die Syscalls; erst
Punkt 5 ändert das Bild strukturell.

## Update: Umsetzung und Ergebnisse (2026-09-02)

Alle Punkte 1–6 der Tabelle wurden umgesetzt (Ausnahme: BdatWriter-Reuse,
bewusst ausgelassen – ~2 % Objekte, unverhältnismäßiges Risiko). Der Client
sendet MAIL/RCPT jetzt über `Client.MailAndRcpt` gepipelint (RFC 2920) mit
sequenziellem Fallback, der Mailer nutzt das in `prepare`.

Beim Umsetzen zusätzlich gefunden und behoben:

- `dotReader.Read` blockierte nach erreichtem Endmarker in `Peek(5)` statt
  sofort `io.EOF` zu liefern. Latent, weil der Drain in `handleData` wegen
  des Shadowing-Bugs nie lief – nach dessen Fix deadlockte der Server.
  Früher EOF-Return ergänzt.
- Data-Race in `tester/compare.go`: die erste Writer-Goroutine las die per
  Closure geteilte Variable `pw`, während die Schleife sie neu zuweist.
  `pw` wird jetzt als Parameter übergeben (wie in `writeInGoroutine`).

Benchstat (je 5 Läufe, gleiche Maschine):

| Benchmark (SameConnection)  | vorher   | nachher  | Delta   |
|-----------------------------|----------|----------|---------|
| Small mit CHUNKING          | 69,9 µs  | 33,0 µs  | −52,8 % |
| Small ohne CHUNKING (DATA)  | 84,8 µs  | 47,4 µs  | −44,1 % |
| Large mit CHUNKING          | 5,92 ms  | 5,24 ms  | −11,5 % |
| Large ohne CHUNKING (DATA)  | 20,3 ms  | 21,1 ms  | ±0 (n.s.) |

Allokationen pro Mail (Small): 111 → 55 (−50 %), Bytes 17,7 KiB → 10,2 KiB
(−42 %). Bei großen Mails ohne Chunking dominiert die Dot-Kodierung der
Nutzdaten, dort ändert sich erwartungsgemäß nichts.

Hinweis: `TestReadMultiLineError` schlägt lokal auch auf unverändertem main
fehl – `net/textproto.Error.Error()` quotet die Message in neueren
Go-1.25-Patches. Betrifft nur die Test-Assertion, nicht das
Laufzeitverhalten (der Code nutzt `Code`/`Msg`, nie den Error-String).
