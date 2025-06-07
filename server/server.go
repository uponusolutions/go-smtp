package server

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"runtime/debug"
	"time"

	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/internal/textsmtp"
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
				if maxDelay := 1 * time.Second; tempDelay > maxDelay {
					tempDelay = maxDelay
				}
				s.logger.ErrorContext(
					ctx,
					"accept error, retrying",
					slog.Any("err", err),
					slog.Any("temp_delay", tempDelay),
				)
				time.Sleep(tempDelay)
				continue
			}
			return err
		}

		s.wg.Add(1)
		go s.handleConn(ctx, c)
	}
}

func (s *Server) handleConn(ctx context.Context, conn net.Conn) {
	ctx, cancel := context.WithCancel(ctx)

	c := &Conn{
		ctx:    ctx,
		server: s,
		conn:   conn,
		text:   textsmtp.NewTextproto(conn, s.readerSize, s.writerSize, s.maxLineLength),
	}

	s.locker.Lock()
	s.conns[c] = struct{}{}
	s.locker.Unlock()

	var err error

	defer func() {
		if err := recover(); err != nil {
			c.writeResponse(421, smtp.EnhancedCode{4, 0, 0}, "Internal server error")
			stack := debug.Stack()
			c.logger().ErrorContext(
				c.ctx,
				"panic serving",
				slog.Any("err", err),
				slog.Any("stack", string(stack)),
			)
			c.Close(errors.New("recovered from panic inside handleConn"))
		}

		s.locker.Lock()
		delete(s.conns, c)
		s.locker.Unlock()

		s.wg.Done()

		cancel()
	}()

	sctx, session, err := s.backend.NewSession(ctx, c)
	if err != nil {
		c.Close(fmt.Errorf("couldn't create connection wrapper: %w", err))
		return
	}

	// update ctx and set session
	c.ctx = sctx
	c.session = session

	c.logger().InfoContext(c.ctx, "connection is opened")

	// explicit tls handshake call
	if tlsConn, ok := c.conn.(*tls.Conn); ok {
		if d := s.readTimeout; d != 0 {
			_ = c.conn.SetReadDeadline(time.Now().Add(d))
		}
		if d := s.writeTimeout; d != 0 {
			_ = c.conn.SetWriteDeadline(time.Now().Add(d))
		}
		if err := tlsConn.Handshake(); err != nil {
			c.handleError(err)
			return
		}
	}

	// run always returns an error when finished
	c.handleError(c.run())
}

// Listen listens on the network address s.Addr
// to handle requests on incoming connections.
//
// If s.Addr is blank and LMTP is disabled, ":smtp" is used.
func (s *Server) Listen() (net.Listener, error) {
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
		return nil, err
	}

	return l, nil
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
		// directly close underlying connection
		_ = conn.conn.Close()
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
