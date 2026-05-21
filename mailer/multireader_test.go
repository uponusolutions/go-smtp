package mailer_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uponusolutions/go-smtp/mailer"
)

func TestClient_Multireader(t *testing.T) {
	a := bytes.NewBuffer([]byte{'a', 'b'})
	b := bytes.NewBuffer([]byte{'c'})
	c := bytes.NewBuffer([]byte{'d'})

	res := mailer.MultiReader(a, b, c)

	require.Equal(t, 4, res.Len())

	bytes := make([]byte, 1)
	i, err := res.Read(bytes)
	require.Equal(t, i, 1)
	require.Equal(t, bytes, []byte{'a'})
	require.NoError(t, err)
	require.Equal(t, 3, res.Len())

	bytes = make([]byte, 2)
	i, err = res.Read(bytes)
	require.NoError(t, err)
	require.Equal(t, i, 1)
	require.Equal(t, 2, res.Len())
	i, err = res.Read(bytes[1:])
	require.NoError(t, err)
	require.Equal(t, i, 1)
	require.Equal(t, bytes, []byte{'b', 'c'})
	require.Equal(t, 1, res.Len())

	bytes = make([]byte, 1)
	i, err = res.Read(bytes)
	require.Equal(t, i, 1)
	require.Equal(t, bytes, []byte{'d'})
	require.NoError(t, err)
	require.Equal(t, 0, res.Len())
}

// Test taken from https://github.com/golang/go/blob/master/src/io/multi_test.go
func TestMultiReader(t *testing.T) {
	var mr io.Reader
	var buf []byte
	nread := 0
	withFooBar := func(tests func()) {
		r1 := strings.NewReader("foo ")
		r2 := strings.NewReader("")
		r3 := strings.NewReader("bar")
		mr = mailer.MultiReader(r1, r2, r3)
		buf = make([]byte, 20)
		tests()
	}
	expectRead := func(size int, expected string, eerr error) {
		nread++
		n, gerr := mr.Read(buf[0:size])
		if n != len(expected) {
			t.Errorf("#%d, expected %d bytes; got %d",
				nread, len(expected), n)
		}
		got := string(buf[0:n])
		if got != expected {
			t.Errorf("#%d, expected %q; got %q",
				nread, expected, got)
		}
		if gerr != eerr {
			t.Errorf("#%d, expected error %v; got %v",
				nread, eerr, gerr)
		}
		buf = buf[n:]
	}
	withFooBar(func() {
		expectRead(2, "fo", nil)
		expectRead(5, "o ", nil)
		expectRead(5, "bar", nil)
		expectRead(5, "", io.EOF)
	})
	withFooBar(func() {
		expectRead(4, "foo ", nil)
		expectRead(1, "b", nil)
		expectRead(3, "ar", nil)
		expectRead(1, "", io.EOF)
	})
	withFooBar(func() {
		expectRead(5, "foo ", nil)
	})
}

// Test taken from https://github.com/golang/go/blob/master/src/io/multi_test.go
func TestMultiReaderAsWriterTo(t *testing.T) {
	mr := mailer.MultiReader(
		strings.NewReader("foo "),
		mailer.MultiReader(
			strings.NewReader(""),
			strings.NewReader("bar"),
		),
	)
	mrAsWriterTo, ok := mr.(io.WriterTo)
	if !ok {
		t.Fatal("expected cast to WriterTo to succeed")
	}
	sink := &strings.Builder{}
	n, err := mrAsWriterTo.WriteTo(sink)
	if err != nil {
		t.Fatalf("expected no error; got %v", err)
	}
	if n != 7 {
		t.Errorf("expected read 7 bytes; got %d", n)
	}
	if result := sink.String(); result != "foo bar" {
		t.Errorf(`expected "foo bar"; got %q`, result)
	}
}
