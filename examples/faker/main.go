package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/smtp"
	"os"
	"os/signal"

	"github.com/uponusolutions/go-smtp/tester"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	s := tester.Standard()

	listen, err := s.Listen(ctx)
	if err != nil {
		slog.Error("error listen server", slog.Any("error", err))
	}

	addr := listen.Addr().String()

	go func() {
		if err := s.Serve(ctx, listen); err != nil {
			slog.Error("smtp server response %s", slog.Any("error", err))
		}
	}()

	defer func() {
		if err := s.Close(ctx); err != nil {
			slog.Error("error closing server", slog.Any("error", err))
		}
	}()

	// Send email.
	from := "alice@i.com"
	to := []string{"bob@e.com", "mal@b.com"}
	msg := []byte("Test\r\n")
	if err := smtp.SendMail(addr, nil, from, to, msg); err != nil {
		panic(err)
	}

	// Lookup email.
	m, found := tester.GetBackend(s).Load(from, to)
	fmt.Printf("Found %t, mail %+v\n", found, m)
}
