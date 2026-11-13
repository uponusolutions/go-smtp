package server

import (
	"crypto/tls"
	"errors"
	"log/slog"
	"net"
	"sync"
	"time"

	"github.com/uponusolutions/go-smtp"
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
	dataTimeout  time.Duration

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

	// Advertise RRVS (RFC 7293) capability.
	// Should be used only if backend supports it.
	enableRRVS bool

	// Advertise DELIVERBY (RFC 2852) capability.
	// Should be used only if backend supports it.
	enableDELIVERBY bool
	// The minimum time, with seconds precision, that a client
	// may specify in the BY argument with return mode.
	// A zero value indicates no set minimum.
	// Only use if DELIVERBY is enabled.
	// 99999999999 is the maximum and 1 the minimum value allowed.
	minimumDeliverInSeconds int

	// Advertise MT-PRIORITY (RFC 6710) capability.
	// Should only be used if backend supports it.
	enableMTPRIORITY bool
	// The priority profile mapping as defined
	// in RFC 6710 section 10.2.
	//
	// Default value of NONE to advertise no specific profile.
	mtPriorityProfile smtp.PriorityProfile

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
//
// The server is bounded by default (line length, message size, recipients
// and read/write timeouts). Each limit can be disabled explicitly by setting
// it to zero with the corresponding option.
func New(opts ...Option) *Server {
	s := &Server{
		done:     make(chan struct{}, 1),
		conns:    make(map[*Conn]struct{}),
		hostname: "localhost",

		// RFC 4954 extends the maximum line length of the AUTH command (and
		// the SASL responses) to 12288 octets, which also covers commands with
		// all usual ESMTP parameters (RFC 5321 allows 512 octets plus extensions).
		maxLineLength: 12288,

		// 32 MiB
		maxMessageBytes: 32 << 20,

		// RFC 5321 requires at least 100 recipients to be accepted.
		maxRecipients: 1000,

		// RFC 5321 (Section 4.5.3.2.7) recommends at least 5 minutes.
		readTimeout: 5 * time.Minute,

		// Whole DATA or BDAT transfer, RFC 5321 (Section 4.5.3.2.6) lets a
		// client wait 10 minutes for the reply after the data, so the transfer
		// itself gets the same. With 32 MiB this requires about 55 KB/s.
		dataTimeout: 10 * time.Minute,

		writeTimeout: 5 * time.Minute,
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

// WithReadTimeout sets the read timeout for every command line.
// DATA and BDAT transfers use the data timeout instead, see WithDataTimeout.
// Defaults to 5 minutes, zero disables the timeout.
func WithReadTimeout(readTimeout time.Duration) Option {
	return func(s *Server) {
		s.readTimeout = readTimeout
	}
}

// WithDataTimeout sets the timeout for a whole DATA or BDAT transfer.
// For BDAT it covers all chunks until the LAST chunk, including the BDAT
// command lines in between. It also applies when a BDAT chunk is discarded.
// Defaults to 10 minutes, zero disables the timeout (no deadline during the transfer).
func WithDataTimeout(dataTimeout time.Duration) Option {
	return func(s *Server) {
		s.dataTimeout = dataTimeout
	}
}

// WithWriteTimeout sets the write timeout.
// Defaults to 5 minutes, zero disables the timeout.
func WithWriteTimeout(writeTimeout time.Duration) Option {
	return func(s *Server) {
		s.writeTimeout = writeTimeout
	}
}

// WithMaxMessageBytes sets the max message size.
// Defaults to 32 MiB, zero (or a negative value) disables the limit.
func WithMaxMessageBytes(maxMessageBytes int64) Option {
	return func(s *Server) {
		s.maxMessageBytes = maxMessageBytes
	}
}

// WithMaxLineLength sets the max length per line.
// Defaults to 12288, zero (or a negative value) disables the limit.
func WithMaxLineLength(maxLineLength int) Option {
	return func(s *Server) {
		s.maxLineLength = maxLineLength
	}
}

// WithMaxRecipients sets the max recipients per mail.
// Defaults to 1000, zero (or a negative value) disables the limit.
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

// WithEnableRRVS advertises RRVS (RFC 7293) capability.
// Should be used only if backend supports it.
func WithEnableRRVS(enableRRVS bool) Option {
	return func(s *Server) {
		s.enableRRVS = enableRRVS
	}
}

// WithEnableDELIVERBY advertises DELIVERBY (RFC 2852) capability.
// Should be used only if backend supports it.
func WithEnableDELIVERBY(enableDELIVERBY bool) Option {
	return func(s *Server) {
		s.enableDELIVERBY = enableDELIVERBY
	}
}

// WithMinimumDeliverInSeconds defines the minimum time, with seconds precision, that a client
// may specify in the BY argument with return mode.
// A zero value indicates no set minimum.
// Only use if DELIVERBY is enabled.
func WithMinimumDeliverInSeconds(minimumDeliverInSeconds int) Option {
	return func(s *Server) {
		s.minimumDeliverInSeconds = minimumDeliverInSeconds
	}
}

// WithEnableMTPRIORITY advertises MT-PRIORITY (RFC 6710) capability.
// Should only be used if backend supports it.
func WithEnableMTPRIORITY(enableMTPRIORITY bool) Option {
	return func(s *Server) {
		s.enableMTPRIORITY = enableMTPRIORITY
	}
}

// WithMtPriorityProfile sets the priority profile mapping as defined
// in RFC 6710 section 10.2.
//
// Default value of NONE to advertise no specific profile.
func WithMtPriorityProfile(mtPriorityProfile smtp.PriorityProfile) Option {
	return func(s *Server) {
		s.mtPriorityProfile = mtPriorityProfile
	}
}
