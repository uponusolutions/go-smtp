package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/smtp"
	"os"
	"os/signal"
	"time"

	"github.com/uponusolutions/go-smtp/tester"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	s := tester.Standard()

	go func() {
		if err := s.ListenAndServe(ctx); err != nil {
			slog.Error("smtp server response %s", slog.Any("error", err))
		}
	}()

	defer func() {
		if err := s.Close(ctx); err != nil {
			slog.Error("error closing server", slog.Any("error", err))
		}
	}()

	// Wait a second to let the server come up.
	time.Sleep(time.Second)

	// Send email.
	from := "alice@i.com"
	to := []string{"bob@e.com", "mal@b.com"}
	msg := []byte("Test\r\n")
	if err := smtp.SendMail(s.Address(), nil, from, to, msg); err != nil {
		panic(err)
	}

	// Lookup email.
	m, found := tester.GetBackend(s).Load(from, to)
	fmt.Printf("Found %t, mail %+v\n", found, m)
}
