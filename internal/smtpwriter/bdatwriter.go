package smtpwriter

import (
	"bufio"
	"errors"
	"io"
	"strconv"
)

var (
	ending   = []byte("BDAT 0 LAST\r\n")
	prefix   = []byte("BDAT ")
	lastcrlf = []byte(" LAST\r\n")
	crlf     = []byte("\r\n")
)

// NewBdat returns a writer that can be used to write bdat commands to w.
// The caller should close the BdatWriter before the next call to a method on w.
func NewBdat(maxChunkSize int, writer *bufio.Writer, read func() error, size int) io.WriteCloser {
	return &bdat{
		w:                  writer,
		read:               read,
		maxChunkSize:       maxChunkSize,
		remainingSize:      size,
		knownSize:          size > 0,
		remainingChunkSize: 0,
	}
}

type bdat struct {
	w    *bufio.Writer
	read func() error
	// maximum bdat chunk size
	maxChunkSize int
	// size still expected to be written
	remainingSize int
	// if true then size is known, so remainingSize is respected
	knownSize bool
	// if larger then 0 bdat command is send and some bytes are left to write
	// only used if knownSize is true
	remainingChunkSize int
	// if Write was called at least once this flag is true.
	started bool
}

// Write writes bytes as multiple bdat commands split by max chunk size.
func (d *bdat) Write(b []byte) (n int, err error) {
	d.started = true

	var p int

	// something left to write
	if d.remainingChunkSize > 0 {
		p, err = d.writeBdat(b[:min(len(b), d.remainingChunkSize)], d.remainingChunkSize == d.remainingSize)
		n += p
		b = b[p:]
		if err != nil {
			return n, err
		}

		if len(b) == 0 {
			return n, nil
		}
	}

	// more then max chunk size received
	for d.maxChunkSize > 0 && len(b) > d.maxChunkSize {
		p, err = d.writeBdat(b[:d.maxChunkSize], false)
		n += p
		b = b[p:]

		if err != nil {
			return n, err
		}
	}

	if d.knownSize && d.remainingSize < len(b) {
		return n, errors.New("got more bytes than expected, check length")
	}

	p, err = d.writeBdat(b, d.knownSize && (d.maxChunkSize == 0 || d.remainingSize <= d.maxChunkSize))
	n += p

	return n, err
}

// write writes b until everything is written
func (d *bdat) write(b []byte) (n int, err error) {
	var p int
	for n < len(b) {
		p, err = d.w.Write(b[n:])
		n += p
		if err != nil {
			return n, err
		}
	}

	if d.remainingChunkSize > 0 {
		d.remainingChunkSize -= n
	}
	if d.knownSize {
		d.remainingSize -= n
	}
	return n, nil
}

// writeBdat writes b as bdat command and checks return
// b must be smaller or equal maxChunkSize
// if size is known we always use max chunk size
func (d *bdat) writeBdat(b []byte, last bool) (n int, err error) {
	if d.remainingChunkSize == 0 {
		size := len(b)
		// if size is known we can create nice chunks
		if d.knownSize && (d.maxChunkSize == 0 || size < d.maxChunkSize) {
			if d.maxChunkSize == 0 {
				size = d.remainingSize
			} else {
				size = min(d.maxChunkSize, d.remainingSize)
			}
			d.remainingChunkSize = size
		}
		if err = d.bdat(size, last); err != nil {
			return n, err
		}
	} else if len(b) > d.remainingChunkSize {
		// just finish the current chunk
		b = b[:d.remainingChunkSize]
	}

	n, err = d.write(b)
	if err != nil {
		return n, err
	}

	// Are there still bytes missing to finish bdat chunk?
	if d.remainingChunkSize > 0 {
		return n, nil
	}

	if err = d.w.Flush(); err != nil {
		return n, err
	}

	// last read is done outside of bdat reader
	if last {
		return n, nil
	}

	return n, d.read()
}

// write BDAT <SIZE> \r\n
func (d *bdat) bdat(size int, last bool) (err error) {
	if _, err = d.w.Write(prefix); err != nil {
		return err
	}

	if _, err = d.w.Write([]byte(strconv.Itoa(size))); err != nil {
		return err
	}

	if last {
		if _, err = d.w.Write(lastcrlf); err != nil {
			return err
		}
		return nil
	}

	if _, err = d.w.Write(crlf); err != nil {
		return err
	}

	return nil
}

func (d *bdat) Close() error {
	// The close came to early, more bytes expected
	if d.knownSize && d.remainingSize > 0 {
		if _, err := d.Write(make([]byte, d.remainingSize)); err != nil {
			return err
		}
	}

	// if size is not known a bdat 0 last is neccesary
	if !d.knownSize || !d.started {
		if _, err := d.w.Write(ending); err != nil {
			return err
		}
	}

	return d.w.Flush()
}
