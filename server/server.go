package server

import (
	"context"
	"crypto/tls"
	"log/slog"
	"net"
	"runtime/debug"
	"time"

	"github.com/uponusolutions/go-smtp"
)

// Serve accepts incoming connections on the Listener l.
func (s *Server) Serve(ctx context.Context, l net.Listener) error {
	s.locker.Lock()
	s.listeners = append(s.listeners, l)
	s.locker.Unlock()

	var tempDelay time.Duration // how long to sleep on accept failure

	for {
		c, err := l.Accept()
		if err != nil {
			select {
			case <-s.done:
				// we called Close()
				return nil
			default:
			}
			if ne, ok := err.(net.Error); ok && ne.Timeout() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				s.logger.ErrorContext(ctx, "accept error, retrying", slog.Any("error", err), slog.Any("temp_delay", tempDelay))
				time.Sleep(tempDelay)
				continue
			}
			return err
		}

		s.wg.Add(1)
		go s.handleConn(ctx, c)
	}
}

func (s *Server) sessionLogger(c *Conn) *slog.Logger {
	var l *slog.Logger
	if c != nil {
		l = c.logger()
	} else {
		l = s.logger
	}
	return l
}

func (s *Server) handleConn(ctx context.Context, conn net.Conn) {
	s.locker.Lock()
	s.conns[conn] = struct{}{}
	s.locker.Unlock()

	ctx, cancel := context.WithCancel(ctx)

	var c *Conn
	var err error

	defer func() {
		if err := recover(); err != nil {
			c.writeResponse(421, smtp.EnhancedCode{4, 0, 0}, "Internal server error")
			stack := debug.Stack()
			s.sessionLogger(c).ErrorContext(
				ctx,
				"panic serving",
				slog.Any("remoteAddr", c.conn.RemoteAddr()),
				slog.Any("err", err),
				slog.Any("stack", stack),
			)
		}

		s.locker.Lock()
		delete(s.conns, conn)
		s.locker.Unlock()

		s.wg.Done()

		cancel()
		_ = conn.Close()
	}()

	c, err = newConn(ctx, conn, s)
	if err != nil {
		s.sessionLogger(c).ErrorContext(ctx, "couldn't create connection wrapper", slog.Any("error", err))
		return
	}

	c.run()
}

// ListenAndServe listens on the network address s.Addr and then calls Serve
// to handle requests on incoming connections.
//
// If s.Addr is blank and LMTP is disabled, ":smtp" is used.
func (s *Server) ListenAndServe(ctx context.Context) error {
	network := s.network
	if network == "" {
		network = "tcp"
	}

	addr := s.addr
	if addr == "" {
		addr = ":smtp"
	}

	var l net.Listener
	var err error

	if s.implicitTLS {
		l, err = tls.Listen(network, addr, s.tlsConfig)
	} else {
		l, err = net.Listen(network, addr)
	}

	if err != nil {
		return err
	}

	return s.Serve(ctx, l)
}

// Close immediately closes all active listeners and connections.
//
// Close returns any error returned from closing the server's underlying
// listener(s).
func (s *Server) Close() error {
	select {
	case <-s.done:
		return ErrServerClosed
	default:
		close(s.done)
	}

	var err error
	s.locker.Lock()
	for _, l := range s.listeners {
		if lerr := l.Close(); lerr != nil && err == nil {
			err = lerr
		}
	}

	for conn := range s.conns {
		conn.Close()
	}
	s.locker.Unlock()

	return err
}

// Shutdown gracefully shuts down the server without interrupting any
// active connections. Shutdown works by first closing all open
// listeners and then waiting indefinitely for connections to return to
// idle and then shut down.
// If the provided context expires before the shutdown is complete,
// Shutdown returns the context's error, otherwise it returns any
// error returned from closing the Server's underlying Listener(s).
func (s *Server) Shutdown(ctx context.Context) error {
	select {
	case <-s.done:
		return ErrServerClosed
	default:
		close(s.done)
	}

	var err error
	s.locker.Lock()
	for _, l := range s.listeners {
		if lerr := l.Close(); lerr != nil && err == nil {
			err = lerr
		}
	}
	s.locker.Unlock()

	connDone := make(chan struct{})
	go func() {
		defer close(connDone)
		s.wg.Wait()
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-connDone:
		return err
	}
}
