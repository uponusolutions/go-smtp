# go-smtp

![coverage](https://raw.githubusercontent.com/uponusolutions/go-smtp/badges/.badges/main/coverage.svg)
[![reference](https://pkg.go.dev/badge/github.com/uponusolutions/go-smtp.svg)](https://pkg.go.dev/github.com/uponusolutions/go-smtp)
[![report](https://goreportcard.com/badge/github.com/uponusolutions/go-smtp)](https://goreportcard.com/report/github.com/uponusolutions/go-smtp)

An ESMTP client and server library written in Go.

## Documentation

  - [Mailer](https://pkg.go.dev/github.com/uponusolutions/go-smtp/mailer) - High-level SMTP client
  - [Client](https://pkg.go.dev/github.com/uponusolutions/go-smtp/client) - Low-level SMTP client
  - [Server](https://pkg.go.dev/github.com/uponusolutions/go-smtp/server) - SMTP server
  - [Resolve](https://pkg.go.dev/github.com/uponusolutions/go-smtp/resolve) - MX-Record resolve
  - [Shared](https://pkg.go.dev/github.com/uponusolutions/go-smtp) - Shared definitions e.g. SMTP status codes
  - [Tester](https://pkg.go.dev/github.com/uponusolutions/go-smtp/tester) - Testing utilities e.g. server with mail map

See [examples](https://github.com/uponusolutions/go-smtp/tree/main/examples) for more information.

## Pages

 - [Coverage](https://uponusolutions.github.io/go-smtp/coverage/)
 - [Benchmark](https://uponusolutions.github.io/go-smtp/dev/bench/)

## Features

* ESMTP client & server implementing [RFC 5321]
* Support for additional SMTP extensions such as [AUTH] and [PIPELINING]
* UTF-8 support for subject and message

## Relationship with emersion/go-smtp

This library is a fork of the great emersion/go-smtp. We try to modernize the api, the performance and supporting non standard extensions like XOORG. If you don't have performance issues or want something very stable, use upstream.

## Relationship with net/smtp

The Go standard library provides a SMTP client implementation in `net/smtp`.
However `net/smtp` is frozen: it's not getting any new features. go-smtp
provides a server implementation and a number of client improvements.

## Licence

MIT

- [RFC 5321](https://tools.ietf.org/html/rfc5321)
- [AUTH](https://tools.ietf.org/html/rfc4954)
- [PIPELINING](https://tools.ietf.org/html/rfc2920)
