// Package resolvermx implements a resolver to get prioritized server addresses for recipients.
package resolve

import (
	"context"
	"errors"
	"fmt"
	"net"
	"slices"
	"strings"
)

// LookupMX describes the functions needed for a struct to be used as a resolver.
type LookupMX interface {
	LookupMX(ctx context.Context, name string) ([]*net.MX, error)
}

// Resolver is
type Resolver struct {
	resolver LookupMX
}

// New creates a new resolver. If nil is given the net.DefaultResolver is used.
func New(resolver LookupMX) Resolver {
	if resolver == nil {
		resolver = net.DefaultResolver
	}
	return Resolver{
		resolver: resolver,
	}
}

// Server contains prioritized server addresses for specific recipients.
type Server struct {
	Addresses [][]string
	Rcpts     []string
}

// Failure contains recipients where no server could be resolved and the corresponding error.
type Failure struct {
	Error error
	Rcpts []string
}

// Result contains Servers and Fails of an recipients server resolve.
type Result struct {
	Servers  []Server
	Failures []Failure
}

func (r *Result) Error() error {
	fails := make([]error, len(r.Failures))
	for i, fail := range r.Failures {
		fails[i] = fmt.Errorf("error for %s: %w", fail.Rcpts, fail.Error)
	}
	return errors.Join(fails...)
}

func (r *Result) addError(rcpt string, err error) int {
	r.Failures = append(r.Failures, Failure{
		Rcpts: []string{rcpt},
		Error: err,
	})
	return len(r.Failures) - 1
}

func (r *Result) addErrorRcpt(index int, rcpt string) int {
	r.Failures[index].Rcpts = append(r.Failures[index].Rcpts, rcpt)
	return len(r.Failures) - 1
}

func (r *Result) addServer(rcpt string, addresses [][]string) int {
	for i, server := range r.Servers {
		if sliceEqual(addresses, server.Addresses) {
			r.Servers[i].Rcpts = append(r.Servers[i].Rcpts, rcpt)
			return i
		}
	}

	r.Servers = append(r.Servers, Server{Rcpts: []string{rcpt}, Addresses: addresses})

	return len(r.Servers) - 1
}

func (r *Result) addServerRcpt(i int, rcpt string) {
	r.Servers[i].Rcpts = append(r.Servers[i].Rcpts, rcpt)
}

type cache struct {
	index int
	err   error
}

// Recipients resolves the smtp servers for specific recipients and groups them.
// If an error returns, then there was an error resolving the MX-Record for the domains of the recipients.
// For example if the network is offline or the dns server isn't responding.
// Failures inside the result are permanent errors like no mx record could be found or
// the domain could not be extracted.
func (r *Resolver) Recipients(ctx context.Context, rcpts []string) (Result, error) {
	res := Result{}
	rcptToDomain := map[string]string{}
	domainToServer := map[string]cache{}

	for _, rcpt := range rcpts {
		domainIndex := strings.LastIndex(rcpt, "@")
		if domainIndex == -1 {
			res.addError(rcpt, fmt.Errorf("couldn't extract domain part of %s", rcpt))
			continue
		}
		domain := rcpt[domainIndex+1:]

		rcptToDomain[rcpt] = domain
		c, ok := domainToServer[domain]
		if ok {
			if c.err != nil {
				res.addErrorRcpt(c.index, rcpt)
			} else {
				res.addServerRcpt(c.index, rcpt)
			}
			continue
		}

		addresses, err := r.Lookup(ctx, domain)
		if err != nil {
			return res, err
		}

		if len(addresses) == 0 {
			err = fmt.Errorf("domain resolve failed, no mx record found for %s", domain)
			domainToServer[domain] = cache{
				index: res.addError(rcpt, err),
				err:   err,
			}
			continue
		}
		domainToServer[domain] = cache{
			index: res.addServer(rcpt, addresses),
		}
	}

	return res, nil
}

// Lookup returns prioritized server addresses for a specific domains.
// It returns nil,nil if no mx record is found and an error if the dns request didn't worked.
func (r *Resolver) Lookup(ctx context.Context, domain string) ([][]string, error) {
	// LookupMX returns the DNS MX records for the given domain name sorted by preference.
	// => We can assume it is sorted and just need
	mxs, err := r.resolver.LookupMX(ctx, domain)
	if err != nil {
		netErr := &net.DNSError{}
		if errors.As(err, &netErr) && netErr.IsNotFound {
			return nil, nil
		}
		return nil, err
	}

	res := [][]string{}

	var prio uint16
	for i := range mxs {
		host := net.JoinHostPort(strings.TrimSuffix(mxs[i].Host, "."), "25")
		if mxs[i].Pref > prio || i == 0 {
			prio = mxs[i].Pref
			res = append(res, []string{
				host,
			})
			continue
		}
		res[len(res)-1] = append(
			res[len(res)-1],
			host,
		)
	}

	return res, nil
}

func sliceEqual(a [][]string, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
	}

	for i := range a {
		if !slices.Equal(a[i], b[i]) {
			return false
		}
	}

	return true
}
