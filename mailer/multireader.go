package mailer

import "io"

// ReaderLen extends io.Reader with Len.
type ReaderLen interface {
	io.Reader
	Len() int
}

type multiReader struct {
	io.Reader
	readers []ReaderLen
}

func (mr *multiReader) Len() int {
	length := 0
	for _, r := range mr.readers {
		length += r.Len()
	}
	return length
}

// MultiReader is essentially io.MultiReader but requires and provides Len.
func MultiReader(readers ...ReaderLen) ReaderLen {
	ioReaders := make([]io.Reader, len(readers))
	for i, r := range readers {
		ioReaders[i] = r
	}

	log := &multiReader{
		Reader:  io.MultiReader(ioReaders...),
		readers: readers,
	}

	return log
}
