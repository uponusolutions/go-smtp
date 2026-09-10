window.BENCHMARK_DATA = {
  "lastUpdate": 1789051271421,
  "repoUrl": "https://github.com/uponusolutions/go-smtp",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "committer": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "distinct": true,
          "id": "21f6161974c03d3aa8e92cec0d1f4dcfbec83487",
          "message": "Merge branch 'emersion-master'",
          "timestamp": "2026-09-07T00:44:47+02:00",
          "tree_id": "c6d5dd74eb1458edf471a01d02dcfe045ffa5370",
          "url": "https://github.com/uponusolutions/go-smtp/commit/21f6161974c03d3aa8e92cec0d1f4dcfbec83487"
        },
        "date": 1788771800005,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 342083,
            "unit": "ns/op",
            "extra": "35332 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 426871,
            "unit": "ns/op",
            "extra": "28236 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 99083,
            "unit": "ns/op",
            "extra": "120146 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 98425,
            "unit": "ns/op",
            "extra": "120728 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 364826,
            "unit": "ns/op",
            "extra": "32835 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 367681,
            "unit": "ns/op",
            "extra": "32612 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 122807,
            "unit": "ns/op",
            "extra": "97845 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 124640,
            "unit": "ns/op",
            "extra": "95845 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 18408269,
            "unit": "ns/op",
            "extra": "661 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 18060078,
            "unit": "ns/op",
            "extra": "661 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 17042112,
            "unit": "ns/op",
            "extra": "703 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 16189252,
            "unit": "ns/op",
            "extra": "738 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 28085106,
            "unit": "ns/op",
            "extra": "417 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 28270390,
            "unit": "ns/op",
            "extra": "403 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 805120063,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 30188968,
            "unit": "ns/op",
            "extra": "344 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 24300774,
            "unit": "ns/op",
            "extra": "493 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1196696,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 24303661,
            "unit": "ns/op",
            "extra": "493 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1195483,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1158970317,
            "unit": "ns/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 49919993,
            "unit": "ns/op",
            "extra": "240 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1172480372,
            "unit": "ns/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 62017582,
            "unit": "ns/op",
            "extra": "193 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "committer": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "distinct": true,
          "id": "f02688e05e0e7f067059c1372c15446d9163f620",
          "message": "feat: remove redundant upper case and improve starttls handling",
          "timestamp": "2026-09-07T11:59:15+02:00",
          "tree_id": "41ca33b3e65f218d7c64bfda9ca7dec9b28a4136",
          "url": "https://github.com/uponusolutions/go-smtp/commit/f02688e05e0e7f067059c1372c15446d9163f620"
        },
        "date": 1788775459660,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 344465,
            "unit": "ns/op",
            "extra": "34398 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 418903,
            "unit": "ns/op",
            "extra": "28329 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 87058,
            "unit": "ns/op",
            "extra": "137946 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 85712,
            "unit": "ns/op",
            "extra": "140560 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 356625,
            "unit": "ns/op",
            "extra": "33630 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 357402,
            "unit": "ns/op",
            "extra": "33338 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 108245,
            "unit": "ns/op",
            "extra": "114649 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 113846,
            "unit": "ns/op",
            "extra": "106302 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 16382346,
            "unit": "ns/op",
            "extra": "738 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 16917868,
            "unit": "ns/op",
            "extra": "688 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 15098475,
            "unit": "ns/op",
            "extra": "781 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 15324049,
            "unit": "ns/op",
            "extra": "794 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 24128177,
            "unit": "ns/op",
            "extra": "499 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 24262691,
            "unit": "ns/op",
            "extra": "502 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 25941927,
            "unit": "ns/op",
            "extra": "410 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 27091355,
            "unit": "ns/op",
            "extra": "411 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 22978291,
            "unit": "ns/op",
            "extra": "528 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1083917,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 23489616,
            "unit": "ns/op",
            "extra": "522 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1076306,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1049173889,
            "unit": "ns/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 45264369,
            "unit": "ns/op",
            "extra": "266 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1054297169,
            "unit": "ns/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 54800909,
            "unit": "ns/op",
            "extra": "217 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "committer": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "distinct": true,
          "id": "e7a5913b16d0a860827e07e1fbf096b6a345ae44",
          "message": "fix: Args parsing now splits only by whitespace and returns an array\ninstead of a map, reverse path removes leading whitespaces from S,\nimproved error handling of Args calls",
          "timestamp": "2026-09-07T12:42:22+02:00",
          "tree_id": "c44283a9e7065ec6661723f62acaa5793b1b455f",
          "url": "https://github.com/uponusolutions/go-smtp/commit/e7a5913b16d0a860827e07e1fbf096b6a345ae44"
        },
        "date": 1788778048086,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 264228,
            "unit": "ns/op",
            "extra": "45412 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 325674,
            "unit": "ns/op",
            "extra": "36915 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 75667,
            "unit": "ns/op",
            "extra": "155265 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 75350,
            "unit": "ns/op",
            "extra": "158689 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 281649,
            "unit": "ns/op",
            "extra": "42723 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 283576,
            "unit": "ns/op",
            "extra": "42296 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 94089,
            "unit": "ns/op",
            "extra": "127352 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 95455,
            "unit": "ns/op",
            "extra": "125744 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 14197231,
            "unit": "ns/op",
            "extra": "817 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 15391279,
            "unit": "ns/op",
            "extra": "772 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 13671575,
            "unit": "ns/op",
            "extra": "870 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 13165309,
            "unit": "ns/op",
            "extra": "892 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 21407760,
            "unit": "ns/op",
            "extra": "544 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 21562695,
            "unit": "ns/op",
            "extra": "553 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 23144853,
            "unit": "ns/op",
            "extra": "454 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 26087052,
            "unit": "ns/op",
            "extra": "493 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 18811650,
            "unit": "ns/op",
            "extra": "636 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 926396,
            "unit": "ns/op",
            "extra": "12946 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 19469944,
            "unit": "ns/op",
            "extra": "626 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 925553,
            "unit": "ns/op",
            "extra": "12963 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 897794906,
            "unit": "ns/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 38811415,
            "unit": "ns/op",
            "extra": "308 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 860109365,
            "unit": "ns/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 50988882,
            "unit": "ns/op",
            "extra": "235 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "committer": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "distinct": true,
          "id": "2f71f54082d5d8457f19097d490f0700fa8598c3",
          "message": "feat: order and rewrite Fprintf to WriteString",
          "timestamp": "2026-09-07T14:33:35+02:00",
          "tree_id": "f55eb71804b1714bcd8424fe562320145c40633f",
          "url": "https://github.com/uponusolutions/go-smtp/commit/2f71f54082d5d8457f19097d490f0700fa8598c3"
        },
        "date": 1788784722529,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 370332,
            "unit": "ns/op",
            "extra": "32562 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 477618,
            "unit": "ns/op",
            "extra": "24864 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 95775,
            "unit": "ns/op",
            "extra": "126028 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 94633,
            "unit": "ns/op",
            "extra": "127270 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 390603,
            "unit": "ns/op",
            "extra": "30681 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 396460,
            "unit": "ns/op",
            "extra": "30500 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 118387,
            "unit": "ns/op",
            "extra": "98367 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 121880,
            "unit": "ns/op",
            "extra": "98846 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 18197019,
            "unit": "ns/op",
            "extra": "643 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 18713306,
            "unit": "ns/op",
            "extra": "621 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 17827760,
            "unit": "ns/op",
            "extra": "666 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 16515726,
            "unit": "ns/op",
            "extra": "706 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 26674617,
            "unit": "ns/op",
            "extra": "447 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 27269717,
            "unit": "ns/op",
            "extra": "428 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 30616627,
            "unit": "ns/op",
            "extra": "357 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 33636462,
            "unit": "ns/op",
            "extra": "344 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 24460876,
            "unit": "ns/op",
            "extra": "484 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1144335,
            "unit": "ns/op",
            "extra": "9285 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 24327444,
            "unit": "ns/op",
            "extra": "494 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1145218,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1117465418,
            "unit": "ns/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 47685456,
            "unit": "ns/op",
            "extra": "250 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1127960743,
            "unit": "ns/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 58605731,
            "unit": "ns/op",
            "extra": "204 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "committer": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "distinct": true,
          "id": "7809f9ab524e3b3cafe0a64276f9942e0c652a36",
          "message": "feat: status multiline support",
          "timestamp": "2026-09-07T17:31:56+02:00",
          "tree_id": "84fdfb203de1be166031386edf5531e8d8f28855",
          "url": "https://github.com/uponusolutions/go-smtp/commit/7809f9ab524e3b3cafe0a64276f9942e0c652a36"
        },
        "date": 1788795418124,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 357604,
            "unit": "ns/op",
            "extra": "33337 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 450485,
            "unit": "ns/op",
            "extra": "26689 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 92194,
            "unit": "ns/op",
            "extra": "122708 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 91316,
            "unit": "ns/op",
            "extra": "131832 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 374624,
            "unit": "ns/op",
            "extra": "32174 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 378933,
            "unit": "ns/op",
            "extra": "31592 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 112941,
            "unit": "ns/op",
            "extra": "106678 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 116588,
            "unit": "ns/op",
            "extra": "102255 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 18806745,
            "unit": "ns/op",
            "extra": "625 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 18828501,
            "unit": "ns/op",
            "extra": "636 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 18208396,
            "unit": "ns/op",
            "extra": "652 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 16768133,
            "unit": "ns/op",
            "extra": "706 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 27393209,
            "unit": "ns/op",
            "extra": "438 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 28861380,
            "unit": "ns/op",
            "extra": "417 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 31517224,
            "unit": "ns/op",
            "extra": "327 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 28691520,
            "unit": "ns/op",
            "extra": "382 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 24349943,
            "unit": "ns/op",
            "extra": "493 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1139289,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 24355265,
            "unit": "ns/op",
            "extra": "492 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1138967,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1110116829,
            "unit": "ns/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 47949258,
            "unit": "ns/op",
            "extra": "249 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1128918820,
            "unit": "ns/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 59166864,
            "unit": "ns/op",
            "extra": "202 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "committer": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "distinct": true,
          "id": "92b75d35398f867227b1ce54c11a72d9fd7d44b5",
          "message": "feat: improve documentation",
          "timestamp": "2026-09-07T17:43:21+02:00",
          "tree_id": "580d62cf180067fbda189268eb3544a55a391bfa",
          "url": "https://github.com/uponusolutions/go-smtp/commit/92b75d35398f867227b1ce54c11a72d9fd7d44b5"
        },
        "date": 1788796102866,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 197188,
            "unit": "ns/op",
            "extra": "60956 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 338715,
            "unit": "ns/op",
            "extra": "35227 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 37290,
            "unit": "ns/op",
            "extra": "315394 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 37141,
            "unit": "ns/op",
            "extra": "320228 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 208907,
            "unit": "ns/op",
            "extra": "57226 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 211717,
            "unit": "ns/op",
            "extra": "56623 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 46816,
            "unit": "ns/op",
            "extra": "255205 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 49415,
            "unit": "ns/op",
            "extra": "243196 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 12232322,
            "unit": "ns/op",
            "extra": "1002 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 13783220,
            "unit": "ns/op",
            "extra": "856 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 11633908,
            "unit": "ns/op",
            "extra": "1029 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 11882402,
            "unit": "ns/op",
            "extra": "1003 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 17966027,
            "unit": "ns/op",
            "extra": "562 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 17744186,
            "unit": "ns/op",
            "extra": "626 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 17874530,
            "unit": "ns/op",
            "extra": "668 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 20859546,
            "unit": "ns/op",
            "extra": "546 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 15300789,
            "unit": "ns/op",
            "extra": "784 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1082141,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 15293531,
            "unit": "ns/op",
            "extra": "784 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1085244,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 778722710,
            "unit": "ns/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 38535385,
            "unit": "ns/op",
            "extra": "310 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 796966663,
            "unit": "ns/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 47783954,
            "unit": "ns/op",
            "extra": "252 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "committer": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "distinct": true,
          "id": "a0e417f3d5f339bb22c7c3b01dbb7fa6ae324659",
          "message": "feat: improve code structure",
          "timestamp": "2026-09-07T22:26:19+02:00",
          "tree_id": "b54bbd853048eb4745b9b1145ab3634343e7df79",
          "url": "https://github.com/uponusolutions/go-smtp/commit/a0e417f3d5f339bb22c7c3b01dbb7fa6ae324659"
        },
        "date": 1788813085243,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 352633,
            "unit": "ns/op",
            "extra": "34242 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 455796,
            "unit": "ns/op",
            "extra": "26173 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 90654,
            "unit": "ns/op",
            "extra": "129926 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 90160,
            "unit": "ns/op",
            "extra": "131575 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 371681,
            "unit": "ns/op",
            "extra": "32294 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 375114,
            "unit": "ns/op",
            "extra": "31941 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 112423,
            "unit": "ns/op",
            "extra": "110022 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 116420,
            "unit": "ns/op",
            "extra": "102499 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 17371139,
            "unit": "ns/op",
            "extra": "692 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 18243416,
            "unit": "ns/op",
            "extra": "669 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 16842242,
            "unit": "ns/op",
            "extra": "733 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 15756164,
            "unit": "ns/op",
            "extra": "759 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 26120385,
            "unit": "ns/op",
            "extra": "458 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 26760964,
            "unit": "ns/op",
            "extra": "458 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 27527952,
            "unit": "ns/op",
            "extra": "421 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 27244050,
            "unit": "ns/op",
            "extra": "416 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 24212735,
            "unit": "ns/op",
            "extra": "495 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1134765,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 24276018,
            "unit": "ns/op",
            "extra": "493 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1132848,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1111169305,
            "unit": "ns/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 47874828,
            "unit": "ns/op",
            "extra": "249 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1123215024,
            "unit": "ns/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 58441548,
            "unit": "ns/op",
            "extra": "205 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "committer": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "distinct": true,
          "id": "7f3e4e29e4734d5a73623a423fee0d4228f46614",
          "message": "feat: Status checks and naming",
          "timestamp": "2026-09-08T12:11:23+02:00",
          "tree_id": "0994ccc30c63be61f5661d2badd83dcbd60b2031",
          "url": "https://github.com/uponusolutions/go-smtp/commit/7f3e4e29e4734d5a73623a423fee0d4228f46614"
        },
        "date": 1788862710257,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 202278,
            "unit": "ns/op",
            "extra": "61250 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 333947,
            "unit": "ns/op",
            "extra": "35572 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 40671,
            "unit": "ns/op",
            "extra": "292376 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 40213,
            "unit": "ns/op",
            "extra": "294750 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 210353,
            "unit": "ns/op",
            "extra": "57399 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 215672,
            "unit": "ns/op",
            "extra": "55464 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 50155,
            "unit": "ns/op",
            "extra": "238801 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 53155,
            "unit": "ns/op",
            "extra": "224209 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 15373181,
            "unit": "ns/op",
            "extra": "794 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 17280699,
            "unit": "ns/op",
            "extra": "687 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 14551934,
            "unit": "ns/op",
            "extra": "796 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 14952291,
            "unit": "ns/op",
            "extra": "822 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 21657701,
            "unit": "ns/op",
            "extra": "501 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 20714656,
            "unit": "ns/op",
            "extra": "572 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 1090234449,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 311706307,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 15996085,
            "unit": "ns/op",
            "extra": "747 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1103498,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 16147436,
            "unit": "ns/op",
            "extra": "750 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1118320,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 813281345,
            "unit": "ns/op",
            "extra": "14 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 44269769,
            "unit": "ns/op",
            "extra": "270 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 854381176,
            "unit": "ns/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 59147375,
            "unit": "ns/op",
            "extra": "201 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "committer": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "distinct": true,
          "id": "aa4d326dd3995f2293ab2b3dbcb4d5d66ad0cf8a",
          "message": "feat: remove unecessary textproto.Pipeline",
          "timestamp": "2026-09-08T16:59:54+02:00",
          "tree_id": "063e3669e8a0237381955a5610b1da55a2747157",
          "url": "https://github.com/uponusolutions/go-smtp/commit/aa4d326dd3995f2293ab2b3dbcb4d5d66ad0cf8a"
        },
        "date": 1788880796143,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 227886,
            "unit": "ns/op",
            "extra": "52713 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 346146,
            "unit": "ns/op",
            "extra": "34512 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 52025,
            "unit": "ns/op",
            "extra": "228930 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 51659,
            "unit": "ns/op",
            "extra": "231129 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 242334,
            "unit": "ns/op",
            "extra": "49496 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 246695,
            "unit": "ns/op",
            "extra": "48600 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 64868,
            "unit": "ns/op",
            "extra": "185757 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 69021,
            "unit": "ns/op",
            "extra": "171716 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 17985157,
            "unit": "ns/op",
            "extra": "694 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 20640439,
            "unit": "ns/op",
            "extra": "574 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 16349826,
            "unit": "ns/op",
            "extra": "712 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 16958089,
            "unit": "ns/op",
            "extra": "718 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 24082938,
            "unit": "ns/op",
            "extra": "504 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 24436797,
            "unit": "ns/op",
            "extra": "495 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 27012724,
            "unit": "ns/op",
            "extra": "379 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 36620413,
            "unit": "ns/op",
            "extra": "277 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 19439383,
            "unit": "ns/op",
            "extra": "613 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1241640,
            "unit": "ns/op",
            "extra": "9530 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 19493248,
            "unit": "ns/op",
            "extra": "602 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1242612,
            "unit": "ns/op",
            "extra": "9571 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1031082456,
            "unit": "ns/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 48583126,
            "unit": "ns/op",
            "extra": "246 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1052275479,
            "unit": "ns/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 64764870,
            "unit": "ns/op",
            "extra": "186 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "committer": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "distinct": true,
          "id": "71573837c9b8002aaf7672066192762c8795ca8c",
          "message": "feat: move server in tester to tester/testserver, cleanup textproto",
          "timestamp": "2026-09-08T17:15:04+02:00",
          "tree_id": "a1c5815c5d7a409d875f6c8f16214c0d63405ee5",
          "url": "https://github.com/uponusolutions/go-smtp/commit/71573837c9b8002aaf7672066192762c8795ca8c"
        },
        "date": 1788880870806,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 356678,
            "unit": "ns/op",
            "extra": "33610 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 442453,
            "unit": "ns/op",
            "extra": "27088 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 90981,
            "unit": "ns/op",
            "extra": "130179 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 91280,
            "unit": "ns/op",
            "extra": "132336 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 377017,
            "unit": "ns/op",
            "extra": "31575 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 381429,
            "unit": "ns/op",
            "extra": "31551 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 114395,
            "unit": "ns/op",
            "extra": "105337 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 116972,
            "unit": "ns/op",
            "extra": "102068 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 18072904,
            "unit": "ns/op",
            "extra": "669 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 19217346,
            "unit": "ns/op",
            "extra": "630 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 17211238,
            "unit": "ns/op",
            "extra": "698 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 16641468,
            "unit": "ns/op",
            "extra": "727 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 26937532,
            "unit": "ns/op",
            "extra": "451 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 27537676,
            "unit": "ns/op",
            "extra": "448 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 323437992,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 168384311,
            "unit": "ns/op",
            "extra": "319 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 20427792,
            "unit": "ns/op",
            "extra": "590 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1131825,
            "unit": "ns/op",
            "extra": "9555 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 20382700,
            "unit": "ns/op",
            "extra": "586 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1131656,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1113522132,
            "unit": "ns/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 47933180,
            "unit": "ns/op",
            "extra": "249 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1130566178,
            "unit": "ns/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 58985140,
            "unit": "ns/op",
            "extra": "202 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "committer": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "distinct": true,
          "id": "e80477d641871ea89687f87c378f831a68faecd7",
          "message": "feat: rework Status and rework textsmtp ReadResponse for our purpose",
          "timestamp": "2026-09-09T02:02:40+02:00",
          "tree_id": "9439a593cc45c345e06d9df6a714f1c7343f3e46",
          "url": "https://github.com/uponusolutions/go-smtp/commit/e80477d641871ea89687f87c378f831a68faecd7"
        },
        "date": 1788912465654,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 371254,
            "unit": "ns/op",
            "extra": "32890 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 487989,
            "unit": "ns/op",
            "extra": "24622 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 94958,
            "unit": "ns/op",
            "extra": "125542 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 95025,
            "unit": "ns/op",
            "extra": "125697 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 388186,
            "unit": "ns/op",
            "extra": "30962 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 392465,
            "unit": "ns/op",
            "extra": "30426 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 117475,
            "unit": "ns/op",
            "extra": "102038 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 122244,
            "unit": "ns/op",
            "extra": "98275 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 17639322,
            "unit": "ns/op",
            "extra": "691 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 18984499,
            "unit": "ns/op",
            "extra": "639 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 17169389,
            "unit": "ns/op",
            "extra": "696 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 15973330,
            "unit": "ns/op",
            "extra": "741 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 26453943,
            "unit": "ns/op",
            "extra": "441 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 27797042,
            "unit": "ns/op",
            "extra": "453 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 27337346,
            "unit": "ns/op",
            "extra": "409 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 27502289,
            "unit": "ns/op",
            "extra": "436 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 24468424,
            "unit": "ns/op",
            "extra": "486 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1136717,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 24415862,
            "unit": "ns/op",
            "extra": "490 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1133593,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1115217634,
            "unit": "ns/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 48093823,
            "unit": "ns/op",
            "extra": "249 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1124925876,
            "unit": "ns/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 58978205,
            "unit": "ns/op",
            "extra": "201 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "committer": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "distinct": true,
          "id": "2d4e67951232e5ba1a80f04ac826ad8293ed4be2",
          "message": "fix: test case",
          "timestamp": "2026-09-09T02:06:59+02:00",
          "tree_id": "2464fdca8f1a61461f6290d7bab63feda1eac555",
          "url": "https://github.com/uponusolutions/go-smtp/commit/2d4e67951232e5ba1a80f04ac826ad8293ed4be2"
        },
        "date": 1788912722116,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 368744,
            "unit": "ns/op",
            "extra": "32000 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 510352,
            "unit": "ns/op",
            "extra": "23493 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 95637,
            "unit": "ns/op",
            "extra": "123974 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 95188,
            "unit": "ns/op",
            "extra": "124593 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 388773,
            "unit": "ns/op",
            "extra": "30998 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 390422,
            "unit": "ns/op",
            "extra": "30602 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 116800,
            "unit": "ns/op",
            "extra": "103167 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 120383,
            "unit": "ns/op",
            "extra": "99855 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 17029772,
            "unit": "ns/op",
            "extra": "708 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 17529239,
            "unit": "ns/op",
            "extra": "673 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 16120072,
            "unit": "ns/op",
            "extra": "727 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 14969499,
            "unit": "ns/op",
            "extra": "781 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 26526170,
            "unit": "ns/op",
            "extra": "459 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 27436336,
            "unit": "ns/op",
            "extra": "456 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 27599612,
            "unit": "ns/op",
            "extra": "424 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 27329262,
            "unit": "ns/op",
            "extra": "426 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 24196998,
            "unit": "ns/op",
            "extra": "495 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1138783,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 24219706,
            "unit": "ns/op",
            "extra": "494 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1135614,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1109470539,
            "unit": "ns/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 47956936,
            "unit": "ns/op",
            "extra": "249 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1123661912,
            "unit": "ns/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 58652377,
            "unit": "ns/op",
            "extra": "205 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "committer": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "distinct": true,
          "id": "88f37904782faa51b1fedba98280da81f83f4c97",
          "message": "feat: small improvements",
          "timestamp": "2026-09-09T02:23:55+02:00",
          "tree_id": "1a2a4213ae32c1e48d327adde8ba5df401b39c56",
          "url": "https://github.com/uponusolutions/go-smtp/commit/88f37904782faa51b1fedba98280da81f83f4c97"
        },
        "date": 1788913743728,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 301565,
            "unit": "ns/op",
            "extra": "38684 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 490796,
            "unit": "ns/op",
            "extra": "25388 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 79916,
            "unit": "ns/op",
            "extra": "148496 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 78575,
            "unit": "ns/op",
            "extra": "153186 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 318468,
            "unit": "ns/op",
            "extra": "36502 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 299618,
            "unit": "ns/op",
            "extra": "39142 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 97207,
            "unit": "ns/op",
            "extra": "123736 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 101604,
            "unit": "ns/op",
            "extra": "116720 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 13931331,
            "unit": "ns/op",
            "extra": "889 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 14080372,
            "unit": "ns/op",
            "extra": "831 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 13259675,
            "unit": "ns/op",
            "extra": "830 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 11418169,
            "unit": "ns/op",
            "extra": "1030 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 20862885,
            "unit": "ns/op",
            "extra": "559 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 20753038,
            "unit": "ns/op",
            "extra": "573 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 24181637,
            "unit": "ns/op",
            "extra": "456 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 22462713,
            "unit": "ns/op",
            "extra": "522 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 18821360,
            "unit": "ns/op",
            "extra": "637 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 917867,
            "unit": "ns/op",
            "extra": "13069 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 19024339,
            "unit": "ns/op",
            "extra": "633 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 917155,
            "unit": "ns/op",
            "extra": "13083 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 899780514,
            "unit": "ns/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 38723261,
            "unit": "ns/op",
            "extra": "308 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 909967571,
            "unit": "ns/op",
            "extra": "13 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 50360046,
            "unit": "ns/op",
            "extra": "237 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "committer": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "distinct": true,
          "id": "a6eba6f610abaa59ffdde53bbaee5772af65c719",
          "message": "feat: rework minimumDeliverByTime to minimumDeliverInSeconds",
          "timestamp": "2026-09-09T11:14:10+02:00",
          "tree_id": "f636415a4351618aeb2c1f78628e63bd7e58f8bb",
          "url": "https://github.com/uponusolutions/go-smtp/commit/a6eba6f610abaa59ffdde53bbaee5772af65c719"
        },
        "date": 1788945552650,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 339498,
            "unit": "ns/op",
            "extra": "35731 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 419825,
            "unit": "ns/op",
            "extra": "28552 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 98651,
            "unit": "ns/op",
            "extra": "121539 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 98248,
            "unit": "ns/op",
            "extra": "122323 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 363400,
            "unit": "ns/op",
            "extra": "33247 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 367187,
            "unit": "ns/op",
            "extra": "32235 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 122713,
            "unit": "ns/op",
            "extra": "97789 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 124494,
            "unit": "ns/op",
            "extra": "96598 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 18327302,
            "unit": "ns/op",
            "extra": "636 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 17826062,
            "unit": "ns/op",
            "extra": "670 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 17299966,
            "unit": "ns/op",
            "extra": "696 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 16112284,
            "unit": "ns/op",
            "extra": "726 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 27872614,
            "unit": "ns/op",
            "extra": "429 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 28487147,
            "unit": "ns/op",
            "extra": "434 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 32199502,
            "unit": "ns/op",
            "extra": "318 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 32980292,
            "unit": "ns/op",
            "extra": "351 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 24308490,
            "unit": "ns/op",
            "extra": "493 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1201997,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 24289353,
            "unit": "ns/op",
            "extra": "492 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1195924,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1256944264,
            "unit": "ns/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 50100588,
            "unit": "ns/op",
            "extra": "238 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1271561853,
            "unit": "ns/op",
            "extra": "8 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 62635748,
            "unit": "ns/op",
            "extra": "190 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "committer": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "distinct": true,
          "id": "b69f90431e7b7bc3e58cf83dc2e42123ae8b317a",
          "message": "feat: rework structure",
          "timestamp": "2026-09-09T12:10:01+02:00",
          "tree_id": "44a619dedab7caffa2f140d8a1e31dc9c232d804",
          "url": "https://github.com/uponusolutions/go-smtp/commit/b69f90431e7b7bc3e58cf83dc2e42123ae8b317a"
        },
        "date": 1788948904382,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 365084,
            "unit": "ns/op",
            "extra": "33186 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 520949,
            "unit": "ns/op",
            "extra": "23228 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 93074,
            "unit": "ns/op",
            "extra": "127856 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 92293,
            "unit": "ns/op",
            "extra": "129243 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 382888,
            "unit": "ns/op",
            "extra": "31472 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 389331,
            "unit": "ns/op",
            "extra": "30886 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 115406,
            "unit": "ns/op",
            "extra": "103930 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 118746,
            "unit": "ns/op",
            "extra": "100581 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 18565054,
            "unit": "ns/op",
            "extra": "699 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 18816330,
            "unit": "ns/op",
            "extra": "642 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 17891946,
            "unit": "ns/op",
            "extra": "678 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 15664429,
            "unit": "ns/op",
            "extra": "772 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 26696690,
            "unit": "ns/op",
            "extra": "450 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 28124005,
            "unit": "ns/op",
            "extra": "399 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 28552741,
            "unit": "ns/op",
            "extra": "362 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 29118430,
            "unit": "ns/op",
            "extra": "374 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 20339923,
            "unit": "ns/op",
            "extra": "590 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1149088,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 20359174,
            "unit": "ns/op",
            "extra": "590 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1146070,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1108393678,
            "unit": "ns/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 48394064,
            "unit": "ns/op",
            "extra": "247 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1127804223,
            "unit": "ns/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 59362275,
            "unit": "ns/op",
            "extra": "202 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "committer": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "distinct": true,
          "id": "12c9785a92b5a0925436d6b920a45a02cec19a86",
          "message": "feat: improve mailer function signatures",
          "timestamp": "2026-09-09T13:56:13+02:00",
          "tree_id": "df6075c86039e26547ff5e0f754649b09fc83fd8",
          "url": "https://github.com/uponusolutions/go-smtp/commit/12c9785a92b5a0925436d6b920a45a02cec19a86"
        },
        "date": 1788955289174,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 359062,
            "unit": "ns/op",
            "extra": "33590 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 460554,
            "unit": "ns/op",
            "extra": "26878 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 93426,
            "unit": "ns/op",
            "extra": "127666 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 93097,
            "unit": "ns/op",
            "extra": "128784 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 382315,
            "unit": "ns/op",
            "extra": "31274 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 384914,
            "unit": "ns/op",
            "extra": "31026 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 116138,
            "unit": "ns/op",
            "extra": "103488 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 120502,
            "unit": "ns/op",
            "extra": "99433 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 18069886,
            "unit": "ns/op",
            "extra": "687 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 18581502,
            "unit": "ns/op",
            "extra": "640 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 17479549,
            "unit": "ns/op",
            "extra": "704 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 16078039,
            "unit": "ns/op",
            "extra": "784 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 27066675,
            "unit": "ns/op",
            "extra": "448 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 26904267,
            "unit": "ns/op",
            "extra": "442 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 212662252,
            "unit": "ns/op",
            "extra": "100 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 27248129,
            "unit": "ns/op",
            "extra": "421 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 20221170,
            "unit": "ns/op",
            "extra": "594 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1151403,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 20257064,
            "unit": "ns/op",
            "extra": "592 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1150108,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1109353438,
            "unit": "ns/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 48103796,
            "unit": "ns/op",
            "extra": "249 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1147118095,
            "unit": "ns/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 58390762,
            "unit": "ns/op",
            "extra": "205 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "committer": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "distinct": true,
          "id": "21f3d248d03b3a118723a5f2270fab2464624ae6",
          "message": "fix: defer large SASL initial response instead of inlining it\n\nhttps://github.com/emersion/go-smtp/pull/307/changes",
          "timestamp": "2026-09-09T15:17:32+02:00",
          "tree_id": "32a8dbeac1dfd6a665379fefd2b1569d1242c4ef",
          "url": "https://github.com/uponusolutions/go-smtp/commit/21f3d248d03b3a118723a5f2270fab2464624ae6"
        },
        "date": 1788960158373,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 363751,
            "unit": "ns/op",
            "extra": "32972 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 472302,
            "unit": "ns/op",
            "extra": "24906 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 92118,
            "unit": "ns/op",
            "extra": "130182 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 91785,
            "unit": "ns/op",
            "extra": "130032 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 377393,
            "unit": "ns/op",
            "extra": "31828 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 381301,
            "unit": "ns/op",
            "extra": "31388 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 114860,
            "unit": "ns/op",
            "extra": "104546 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 117888,
            "unit": "ns/op",
            "extra": "99940 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 16441367,
            "unit": "ns/op",
            "extra": "728 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 17509574,
            "unit": "ns/op",
            "extra": "687 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 15256597,
            "unit": "ns/op",
            "extra": "808 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 14917837,
            "unit": "ns/op",
            "extra": "812 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 25493500,
            "unit": "ns/op",
            "extra": "482 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 25848495,
            "unit": "ns/op",
            "extra": "447 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 25553604,
            "unit": "ns/op",
            "extra": "409 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 25222169,
            "unit": "ns/op",
            "extra": "448 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 20256700,
            "unit": "ns/op",
            "extra": "591 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1142689,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 20349651,
            "unit": "ns/op",
            "extra": "579 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1142883,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1109709187,
            "unit": "ns/op",
            "extra": "10 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 48305211,
            "unit": "ns/op",
            "extra": "247 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1122217674,
            "unit": "ns/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 59542613,
            "unit": "ns/op",
            "extra": "201 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "marcel.link@uponu.com",
            "name": "Marcel Link",
            "username": "ml1nk"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "f1e7ad4bdf3c54552429f6bd8270584efeea115d",
          "message": "Merge pull request #10 from uponusolutions/fix/ms-1560\n\nDotreader stuff first line and more test cases",
          "timestamp": "2026-09-10T16:36:08+02:00",
          "tree_id": "28f3b0ccfa30ff48d941d3878da0989fd3d4d3ca",
          "url": "https://github.com/uponusolutions/go-smtp/commit/f1e7ad4bdf3c54552429f6bd8270584efeea115d"
        },
        "date": 1789051270700,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSmallWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 330742,
            "unit": "ns/op",
            "extra": "36595 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 432403,
            "unit": "ns/op",
            "extra": "28102 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 97655,
            "unit": "ns/op",
            "extra": "122647 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 97114,
            "unit": "ns/op",
            "extra": "122701 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 359723,
            "unit": "ns/op",
            "extra": "33409 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 364630,
            "unit": "ns/op",
            "extra": "32767 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 121163,
            "unit": "ns/op",
            "extra": "98533 times\n4 procs"
          },
          {
            "name": "BenchmarkSmallWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 125260,
            "unit": "ns/op",
            "extra": "96577 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunking (github.com/uponusolutions/go-smtp)",
            "value": 18148163,
            "unit": "ns/op",
            "extra": "614 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 17851873,
            "unit": "ns/op",
            "extra": "673 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 16436012,
            "unit": "ns/op",
            "extra": "716 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 16462431,
            "unit": "ns/op",
            "extra": "736 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunking (github.com/uponusolutions/go-smtp)",
            "value": 28511628,
            "unit": "ns/op",
            "extra": "403 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 29559792,
            "unit": "ns/op",
            "extra": "415 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnection (github.com/uponusolutions/go-smtp)",
            "value": 29429224,
            "unit": "ns/op",
            "extra": "351 times\n4 procs"
          },
          {
            "name": "BenchmarkLargeWithoutChunkingSameConnectionSimpleReader (github.com/uponusolutions/go-smtp)",
            "value": 30770019,
            "unit": "ns/op",
            "extra": "400 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 19935505,
            "unit": "ns/op",
            "extra": "602 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1200365,
            "unit": "ns/op",
            "extra": "9951 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 20045146,
            "unit": "ns/op",
            "extra": "603 times\n4 procs"
          },
          {
            "name": "BenchmarkDotReaderOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1201238,
            "unit": "ns/op",
            "extra": "9976 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacy (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1162267773,
            "unit": "ns/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimized (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 49833161,
            "unit": "ns/op",
            "extra": "240 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterLegacySimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 1167053953,
            "unit": "ns/op",
            "extra": "9 times\n4 procs"
          },
          {
            "name": "BenchmarkDotWriterOptimizedSimpleReader (github.com/uponusolutions/go-smtp/internal/textsmtp)",
            "value": 62362236,
            "unit": "ns/op",
            "extra": "192 times\n4 procs"
          }
        ]
      }
    ]
  }
}