package client

import "errors"

var (
	// ErrPipeliningNotEnabled means pipeline is not enabled
	ErrPipeliningNotEnabled = errors.New("pipelining isn't enabled")
	// ErrPipeliningGroupEnded means you first need to consume all responses
	ErrPipeliningGroupEnded = errors.New("pipelining group ended, please consume responses")
	// ErrPipeliningNothingPending means nothing is pending
	ErrPipeliningNothingPending = errors.New("pipelining has no pending responses")
	// ErrPipeliningNoPendingRequired means something is pending but it is required that there is none pending
	ErrPipeliningNoPendingRequired = errors.New("pipelining does not work with this command," +
		" please consume all pending responses")
)
