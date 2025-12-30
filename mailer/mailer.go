// Package mailer implements a mailer to send mail by smtp.
// It uses the client library and provides high level functionality.
package mailer

import (
	"context"
	"errors"
	"io"
	"net"

	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/client"
	"github.com/uponusolutions/go-smtp/resolvemx"
)

// Mailer implements a smtp client with .
type Mailer struct {
	client *client.Client
	cfg    additionalConfig
}

// New returns a new smtp client.
// When not set via options a default tls.Config and used an StartTLS is preferred but not enforced.
// You need to set at least the server address to get a working mailer.
func New(opts ...Option) *Mailer {
	cfg := DefaultConfig()

	for _, o := range opts {
		o(&cfg)
	}

	return NewFromConfig(cfg)
}

// NewFromConfig returns a new smtp client from existing config.
func NewFromConfig(cfg Config) *Mailer {
	return &Mailer{
		cfg:    cfg.extra,
		client: client.NewFromConfig(cfg.client),
	}
}

// Connect connects to one of the available smtp server.
// When server supports auth and clients SaslClient is set, auth is called.
// Security is enforced like configured (Plain, TLS, StartTLS or PreferStartTLS)
// If an error occures, the connection is closed if open.
func (c *Mailer) Connect(ctx context.Context) error {
	var err error

	for i := 0; i < len(c.cfg.serverAddresses); i++ {
		for p := 0; p < len(c.cfg.serverAddresses[i]); p++ {
			// use c.serverAddressIndex
			address := c.cfg.serverAddresses[i][(p+c.cfg.serverAddressIndex)%len(c.cfg.serverAddresses[i])]
			err = c.connectAddress(ctx, address)
			if err == nil {
				return nil
			}
		}
	}

	return err
}

// Connect connects to the SMTP server (addr).
// When server supports auth and clients SaslClient is set, auth is called.
// Security is enforced like configured (Plain, TLS, StartTLS or PreferStartTLS)
// If an error occures, the connection is closed if open.
func (c *Mailer) connectAddress(ctx context.Context, addr string) error {
	var err error

	switch c.cfg.security {
	case SecurityTLS:
		err = c.client.DialTLS(ctx, c.cfg.tlsConfig, addr)
	case SecurityPlain, SecurityStartTLS, SecurityPreferStartTLS:
		fallthrough
	default:
		err = c.client.Dial(ctx, addr)
	}

	if err != nil {
		return err
	}

	if c.cfg.security == SecurityStartTLS || c.cfg.security == SecurityPreferStartTLS {
		if ok, _ := c.client.Extension("STARTTLS"); !ok {
			if c.cfg.security == SecurityStartTLS {
				_ = c.client.Quit()
				return errors.New("smtp: server doesn't support STARTTLS")
			}
		} else {
			serverName, _, _ := net.SplitHostPort(addr)

			err = c.client.StartTLS(c.cfg.tlsConfig, serverName)
			if err != nil {
				return err
			}
		}
	}

	return c.auth()
}

func (c *Mailer) auth() error {
	// Authenticate if authentication is possible and sasl client available.
	if ok, _ := c.client.Extension("AUTH"); ok && c.cfg.saslClient != nil {
		if err := c.client.Auth(c.cfg.saslClient); err != nil {
			_ = c.client.Quit()
			return err
		}
	}
	return nil
}

// Len defines the Len method existing in some structs to get the length of the internal []byte (e.g. bytes.Buffer)
type Len interface {
	Len() int
}

func (c *Mailer) prepare(
	ctx context.Context,
	from string,
	mailOptions *client.MailOptions,
	rcpt []string,
	rcptsOptions []*smtp.RcptOptions,
	size int,
) (*client.DataCloser, []resolvemx.Failure, error) {
	if !c.client.Connected() {
		err := c.Connect(ctx)
		if err != nil {
			return nil, nil, err
		}
	}

	if len(rcpt) < 1 {
		return nil, nil, errors.New("no recipients")
	}

	if mailOptions == nil && size > 0 {
		mailOptions = &client.MailOptions{
			Size: int64(size),
		}
	}

	// MAIL FROM:
	if err := c.client.Mail(from, mailOptions); err != nil {
		return nil, nil, err
	}

	failures := []resolvemx.Failure{}

	// RCPT TO:
	for i, addr := range rcpt {
		var rcptsOption *smtp.RcptOptions
		if len(rcptsOptions) > i {
			rcptsOption = rcptsOptions[i]
		}

		if err := c.client.Rcpt(addr, rcptsOption); err != nil {
			smtpErr := &smtp.Status{}

			// continue sending if code is 550 Requested action not taken
			if c.cfg.noPartialSend || !errors.As(err, &smtpErr) || smtpErr.Code != 550 {
				return nil, nil, err
			}

			failures = append(failures, resolvemx.Failure{
				Rcpts: []string{addr},
				Error: err,
			})
		}
	}

	// DATA
	w, err := c.client.Content(size)
	if err != nil {
		return nil, nil, err
	}
	return w, failures, nil
}

// Send send an email from
// address from, to addresses to, with message stream in.
//
// It will use an existing connection if possible or create a new one otherwise.
//
// The addresses in the rcpts parameter are the SMTP RCPT addresses.
//
// The in parameter should be a stream of an RFC 822-style email with headers
// first, a blank line, and then the message body. The lines of in
// should be CRLF terminated. The in headers should usually include
// fields such as "From", "To", "Subject", and "Cc".  Sending "Bcc"
// messages is accomplished by including an email address in the to
// parameter but not including it in the in headers.
func (c *Mailer) Send(ctx context.Context, from string, rcpt []string, in io.Reader) (code int, msg string, failures []resolvemx.Failure, err error) {
	return c.SendAdvanced(ctx, from, nil, rcpt, nil, in)
}

// SendAdvanced send an email from
// address from, to addresses to, with message stream in.
//
// It will use an existing connection if possible or create a new one otherwise.
//
// The addresses in the rcpts parameter are the SMTP RCPT addresses.
// If mailOptions isn't set, default values are used (e.g. size determined from in)
// If rcptsOptions isn't set for some rcpts (e.g. len(rcpts) > len(rcptsOptions)),
// default values are used for these recipients.
//
// The in parameter should be a stream of an RFC 822-style email with headers
// first, a blank line, and then the message body. The lines of in
// should be CRLF terminated. The in headers should usually include
// fields such as "From", "To", "Subject", and "Cc".  Sending "Bcc"
// messages is accomplished by including an email address in the to
// parameter but not including it in the in headers.
func (c *Mailer) SendAdvanced(
	ctx context.Context,
	from string,
	mailOptions *client.MailOptions,
	rcpts []string,
	rcptsOptions []*smtp.RcptOptions,
	in io.Reader,
) (code int, msg string, failures []resolvemx.Failure, err error) {
	size := 0
	if wt, ok := in.(Len); ok {
		size = wt.Len()
	}

	w, failures, err := c.prepare(ctx, from, mailOptions, rcpts, rcptsOptions, size)
	if err != nil {
		return 0, "", failures, err
	}

	_, err = io.Copy(w.Writer(), in)
	if err != nil {
		// if err isn't smtp.Status we are in an unknown state, close connection
		if _, ok := err.(*smtp.Status); !ok {
			err = errors.Join(err, c.client.Close())
		}
		return 0, "", failures, err
	}

	code, msg, err = w.CloseWithResponse()

	// if err isn't smtp.Status we are in an unknown state, close connection
	if _, ok := err.(*smtp.Status); err != nil && !ok {
		err = errors.Join(err, c.client.Close())
	}

	return code, msg, failures, err
}

// Verify checks the validity of an email address on the server.
// If Verify returns nil, the address is valid. A non-nil return
// does not necessarily indicate an invalid address. Many servers
// will not verify addresses for security reasons.
//
// If server returns an error, it will be of type *smtp.
func (c *Mailer) Verify(addr string, opts *client.VrfyOptions) error {
	return c.client.Verify(addr, opts)
}

// Disconnect ends current connection gracefully, if any exists.
func (c *Mailer) Disconnect() error {
	return c.client.Quit()
}

// Terminate ends current connection forcefully.
func (c *Mailer) Terminate() error {
	return c.client.Close()
}

// Connected returns the current server name.
func (c *Mailer) Connected() bool {
	return c.client.Connected()
}

// ServerAddress returns the current server address.
func (c *Mailer) ServerAddress() string {
	return c.client.ServerAddress()
}

// ServerName returns the current server name.
func (c *Mailer) ServerName() string {
	return c.client.ServerName()
}

// Report contains responses and failures after sending a mail.
type Report struct {
	Responses []Response
	Failures  []resolvemx.Failure
}

// Response contains the response of a smtp server for specific recipients.
type Response struct {
	Code  int
	Msg   string
	Rcpts []string
}

// Send just sends a mail.
func Send(ctx context.Context, from string, rcpts []string, in func() io.Reader, opts ...Option) (res Report, err error) {
	r := resolvemx.New(nil)

	config := NewConfig(opts...)

	var mx resolvemx.Result

	if len(config.extra.serverAddresses) > 0 {
		mx = resolvemx.Result{
			Servers: []resolvemx.Server{
				{
					Rcpts:     rcpts,
					Addresses: config.extra.serverAddresses,
				},
			},
		}
	} else {
		mx, err = r.Recipients(context.Background(), rcpts)
		if err != nil {
			return Report{}, err
		}
	}

	res.Failures = mx.Failures

	for _, server := range mx.Servers {
		code, msg, failures, err := send(ctx, server, from, config, in())

		if err != nil {
			res.Failures = append(res.Failures, resolvemx.Failure{
				Rcpts: server.Rcpts,
				Error: err,
			})
		} else {
			if len(failures) > 0 {
				rcpts := []string{}

			outer:
				for _, rcpt := range server.Rcpts {
					for _, fail := range failures {
						for _, ircpt := range fail.Rcpts {
							if rcpt == ircpt {
								continue outer
							}
						}
					}
					rcpts = append(rcpts, rcpt)
				}
				server.Rcpts = rcpts
				res.Failures = append(res.Failures, failures...)
			}

			res.Responses = append(res.Responses, Response{
				Code:  code,
				Msg:   msg,
				Rcpts: server.Rcpts,
			})
		}
	}

	return res, nil
}

func send(ctx context.Context, server resolvemx.Server, from string, config Config, in io.Reader) (code int, msg string, failures []resolvemx.Failure, err error) {
	config.extra.serverAddresses = server.Addresses
	client := NewFromConfig(config)
	defer func() { _ = client.Disconnect() }()
	return client.Send(ctx, from, server.Rcpts, in)
}
