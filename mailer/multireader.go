package mailer

import (
	"bytes"
	"io"
)

// ReaderWriteToLen extends io.Reader with Len and WriteTo.
type ReaderWriteToLen interface {
	io.Reader
	io.WriterTo
	Len() int
}

// Modified from https://github.com/golang/go/blob/master/src/io/multi_test.go
type multiReader struct {
	readers []ReaderWriteToLen
}

type eofReader struct{}

func (eofReader) Read([]byte) (int, error) {
	return 0, io.EOF
}

func (eofReader) Len() int {
	return 0
}

func (eofReader) WriteTo(_ io.Writer) (sum int64, err error) {
	return 0, nil
}

func (mr *multiReader) Read(p []byte) (n int, err error) {
	for len(mr.readers) > 0 {
		// Optimization to flatten nested multiReaders (Issue 13558).
		if len(mr.readers) == 1 {
			if r, ok := mr.readers[0].(*multiReader); ok {
				mr.readers = r.readers
				continue
			}
		}
		n, err = mr.readers[0].Read(p)
		if err == io.EOF {
			// Use eofReader instead of nil to avoid nil panic
			// after performing flatten (Issue 18232).
			mr.readers[0] = eofReader{} // permit earlier GC
			mr.readers = mr.readers[1:]
		}
		if n > 0 || err != io.EOF {
			if err == io.EOF && len(mr.readers) > 0 {
				// Don't return EOF yet. More readers remain.
				err = nil
			}
			return n, err
		}
	}
	return 0, io.EOF
}

func (mr *multiReader) WriteTo(w io.Writer) (sum int64, err error) {
	n := int64(0)
	for _, r := range mr.readers {
		rn, err := r.WriteTo(w)
		n += rn
		if err != nil {
			return n, err
		}
	}
	return n, nil
}

func (mr *multiReader) Len() int {
	length := 0
	for _, r := range mr.readers {
		length += r.Len()
	}
	return length
}

// MultiReaderFactory is a utility struct to ease MultiReader creation.
type MultiReaderFactory struct {
	readers []ReaderWriteToLen
}

// AddReader adds the next readers to the factory.
// The new [MultiReaderFactory] takes ownership of the Readers,
func (mrf *MultiReaderFactory) AddReader(readers ...ReaderWriteToLen) {
	mrf.readers = append(mrf.readers, readers...)
}

// AddBytes adds the multiple byte slices to the factory.
// The new [MultiReaderFactory] takes ownership of the bytes,
func (mrf *MultiReaderFactory) AddBytes(arrbytes ...[]byte) {
	for _, reader := range arrbytes {
		mrf.readers = append(mrf.readers, bytes.NewBuffer(reader))
	}
}

// Create creates a MultiReader from all added readers.
func (mrf *MultiReaderFactory) Create() ReaderWriteToLen {
	if len(mrf.readers) == 0 {
		return eofReader{}
	}

	var reader ReaderWriteToLen

	if len(mrf.readers) == 1 {
		reader = mrf.readers[0]
	} else {
		reader = MultiReader(mrf.readers...)
	}

	// reset state
	mrf.readers = nil

	return reader
}

// MultiReader returns a Reader that's the logical concatenation of
// the provided input readers. They're read sequentially. Once all
// inputs have returned EOF, Read will return EOF.  If any of the readers
// return a non-nil, non-EOF error, Read will return that error.
// The new [Reader] takes ownership of the Readers,
func MultiReader(readers ...ReaderWriteToLen) ReaderWriteToLen {
	return &multiReader{readers}
}
