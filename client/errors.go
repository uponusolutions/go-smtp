package client

import "errors"

var (
	// ErrPipeliningNotEnabled means pipeline is not enabled
	ErrPipeliningNotEnabled = errors.New("pipelining isn't enabled")
	// ErrPipeliningGroupConcluded means you first need to consume all responses
	ErrPipeliningGroupConcluded = errors.New("pipelining group concluded, please consume responses")
	// ErrPipeliningNothingPending means nothing is pending
	ErrPipeliningNothingPending = errors.New("pipelining has no pending responses")
	// ErrPipeliningNoPendingRequired means something is pending but it is required that there is none pending
	ErrPipeliningNoPendingRequired = errors.New("pipelining does not work with this command," +
		" please consume all pending responses")
	// ErrPipeliningCongestion means you wrote to many bytes and the group was forcefully concluded.
	ErrPipeliningCongestion = errors.New("pipelining group concluded because of congestion, please consume responses")
)
