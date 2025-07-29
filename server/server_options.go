package server

import (
	"crypto/tls"
	"errors"
	"log/slog"
	"net"
	"sync"
	"time"
)

// ErrServerClosed occurs if a server is already closed.
var ErrServerClosed = errors.New("smtp: server already closed")

// Server implements a SMTP server.
type Server struct {
	// The type of network, "tcp" or "unix".
	network string
	// TCP or Unix address to listen on.
	addr string
	// The server TLS configuration.
	tlsConfig *tls.Config

	hostname string

	maxRecipients int
	// Max line length for every command except data and bdat.
	maxLineLength int
	// Maximum size when receiving data and bdat.
	maxMessageBytes int64
	// Reader buffer size.
	readerSize int
	// Writer buffer size.
	writerSize int

	readTimeout  time.Duration
	writeTimeout time.Duration

	implicitTLS bool

	// Enforces usage of implicit tls or starttls before accepting commands except NOOP, EHLO, STARTTLS, or QUIT.
	enforceSecureConnection bool

	// Enforces usage of authentication.
	enforceAuthentication bool

	// Advertise SMTPUTF8 (RFC 6531) capability.
	// Should be used only if backend supports it.
	enableSMTPUTF8 bool

	// Advertise REQUIRETLS (RFC 8689) capability.
	// Should be used only if backend supports it.
	enableREQUIRETLS bool

	// Advertise CHUNKING (RFC 1830) capability.
	enableCHUNKING bool

	// Advertise BINARYMIME (RFC 3030) capability.
	// Should be used only if backend supports it.
	enableBINARYMIME bool

	// Advertise DSN (RFC 3461) capability.
	// Should be used only if backend supports it.
	enableDSN bool

	// Advertise XOORG capability.
	// Should be used only if backend supports it.
	enableXOORG bool

	// The server backend.
	backend Backend

	logger *slog.Logger

	wg   sync.WaitGroup
	done chan struct{}

	locker    sync.Mutex
	listeners []net.Listener
	conns     map[*Conn]struct{}
}

// Backend returns the servers Backend.
func (s *Server) Backend() Backend {
	return s.backend
}

// Option is an option for the server.
type Option func(*Server)

// New creates a new SMTP server.
func New(opts ...Option) *Server {
	s := &Server{
		done:     make(chan struct{}, 1),
		conns:    make(map[*Conn]struct{}),
		hostname: "localhost",
	}

	for _, o := range opts {
		o(s)
	}

	if s.logger == nil {
		s.logger = slog.Default()
	}

	return s
}

// WithLogger sets the backend.
func WithLogger(logger *slog.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

// WithBackend sets the backend.
func WithBackend(backend Backend) Option {
	return func(s *Server) {
		s.backend = backend
	}
}

// WithNetwork sets the network.
func WithNetwork(network string) Option {
	return func(s *Server) {
		s.network = network
	}
}

// WithReadTimeout sets the read timeout.
func WithReadTimeout(readTimeout time.Duration) Option {
	return func(s *Server) {
		s.readTimeout = readTimeout
	}
}

// WithWriteTimeout sets the write timeout.
func WithWriteTimeout(writeTimeout time.Duration) Option {
	return func(s *Server) {
		s.writeTimeout = writeTimeout
	}
}

// WithMaxMessageBytes sets the max message size.
func WithMaxMessageBytes(maxMessageBytes int64) Option {
	return func(s *Server) {
		s.maxMessageBytes = maxMessageBytes
	}
}

// WithMaxLineLength sets the max length per line.
func WithMaxLineLength(maxLineLength int) Option {
	return func(s *Server) {
		s.maxLineLength = maxLineLength
	}
}

// WithMaxRecipients sets the max recipients per mail.
func WithMaxRecipients(maxRecipients int) Option {
	return func(s *Server) {
		s.maxRecipients = maxRecipients
	}
}

// WithAddr sets addr.
func WithAddr(addr string) Option {
	return func(s *Server) {
		s.addr = addr
	}
}

// WithEnableXOORG enables xoorg.
func WithEnableXOORG(enableXOORG bool) Option {
	return func(s *Server) {
		s.enableXOORG = enableXOORG
	}
}

// WithEnableBINARYMIME sets EnableBINARYMIME.
func WithEnableBINARYMIME(enableBINARYMIME bool) Option {
	return func(s *Server) {
		s.enableBINARYMIME = enableBINARYMIME
	}
}

// WithEnableREQUIRETLS sets EnableREQUIRETLS.
func WithEnableREQUIRETLS(enableREQUIRETLS bool) Option {
	return func(s *Server) {
		s.enableREQUIRETLS = enableREQUIRETLS
	}
}

// WithEnableCHUNKING sets EnableCHUNKING.
func WithEnableCHUNKING(enableCHUNKING bool) Option {
	return func(s *Server) {
		s.enableCHUNKING = enableCHUNKING
	}
}

// WithEnableSMTPUTF8 sets EnableSMTPUTF8.
func WithEnableSMTPUTF8(enableSMTPUTF8 bool) Option {
	return func(s *Server) {
		s.enableSMTPUTF8 = enableSMTPUTF8
	}
}

// WithEnableDSN sets EnableDSN.
func WithEnableDSN(enableDSN bool) Option {
	return func(s *Server) {
		s.enableDSN = enableDSN
	}
}

// WithImplicitTLS sets implicitTLS.
func WithImplicitTLS(implicitTLS bool) Option {
	return func(s *Server) {
		s.implicitTLS = implicitTLS
	}
}

// WithHostname sets the domain.
func WithHostname(hostname string) Option {
	return func(s *Server) {
		s.hostname = hostname
	}
}

// WithTLSConfig sets certificate.
func WithTLSConfig(tlsConfig *tls.Config) Option {
	return func(s *Server) {
		s.tlsConfig = tlsConfig
	}
}

// WithEnforceSecureConnection enforces implicit TLS or STARTTLS.
func WithEnforceSecureConnection(enforceSecureConnection bool) Option {
	return func(s *Server) {
		s.enforceSecureConnection = enforceSecureConnection
	}
}

// WithEnforceAuthentication enforces authentication before mail usage.
func WithEnforceAuthentication(enforceAuthentication bool) Option {
	return func(s *Server) {
		s.enforceAuthentication = enforceAuthentication
	}
}

// WithReaderSize sets ReaderSize.
func WithReaderSize(readerSize int) Option {
	return func(s *Server) {
		s.readerSize = readerSize
	}
}

// WithWriterSize sets WriterSize.
func WithWriterSize(writerSize int) Option {
	return func(s *Server) {
		s.writerSize = writerSize
	}
}
