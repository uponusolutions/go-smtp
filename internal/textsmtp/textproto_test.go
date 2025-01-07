// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package textsmtp

import (
	"bytes"
	"io"
	"net/textproto"
	"testing"

	"github.com/uponusolutions/go-smtp/internal/faker"
)

func reader(in string, out *bytes.Buffer) *Conn {
	return NewConn(faker.NewConn(in, out))
}

func TestPrintfLine(t *testing.T) {
	buf := &bytes.Buffer{}
	w := reader("", buf)
	err := w.PrintfLine("foo %d", 123)
	if s := buf.String(); s != "foo 123\r\n" || err != nil {
		t.Fatalf("s=%q; err=%s", s, err)
	}
}

func TestReadLine(t *testing.T) {
	r := reader("line1\nline2\n", &bytes.Buffer{})
	s, err := r.ReadLine()
	if s != "line1" || err != nil {
		t.Fatalf("Line 1: %s, %v", s, err)
	}
	s, err = r.ReadLine()
	if s != "line2" || err != nil {
		t.Fatalf("Line 2: %s, %v", s, err)
	}
	s, err = r.ReadLine()
	if s != "" || err != io.EOF {
		t.Fatalf("EOF: %s, %v", s, err)
	}
}

func TestReadCodeLine(t *testing.T) {
	r := reader("123 hi\n234 bye\n345 no way\n", &bytes.Buffer{})
	code, msg, err := r.ReadCodeLine(0)
	if code != 123 || msg != "hi" || err != nil {
		t.Fatalf("Line 1: %d, %s, %v", code, msg, err)
	}
	code, msg, err = r.ReadCodeLine(23)
	if code != 234 || msg != "bye" || err != nil {
		t.Fatalf("Line 2: %d, %s, %v", code, msg, err)
	}
	code, msg, err = r.ReadCodeLine(346)
	if code != 345 || msg != "no way" || err == nil {
		t.Fatalf("Line 3: %d, %s, %v", code, msg, err)
	}
	if e, ok := err.(*textproto.Error); !ok || e.Code != code || e.Msg != msg {
		t.Fatalf("Line 3: wrong error %v\n", err)
	}
	code, msg, err = r.ReadCodeLine(1)
	if code != 0 || msg != "" || err != io.EOF {
		t.Fatalf("EOF: %d, %s, %v", code, msg, err)
	}
}

type readResponseTest struct {
	in       string
	inCode   int
	wantCode int
	wantMsg  string
}

var readResponseTests = []readResponseTest{
	{"230-Anonymous access granted, restrictions apply\n" +
		"Read the file README.txt,\n" +
		"230  please",
		23,
		230,
		"Anonymous access granted, restrictions apply\nRead the file README.txt,\n please",
	},

	{"230 Anonymous access granted, restrictions apply\n",
		23,
		230,
		"Anonymous access granted, restrictions apply",
	},

	{"400-A\n400-B\n400 C",
		4,
		400,
		"A\nB\nC",
	},

	{"400-A\r\n400-B\r\n400 C\r\n",
		4,
		400,
		"A\nB\nC",
	},
}

// See https://www.ietf.org/rfc/rfc959.txt page 36.
func TestRFC959Lines(t *testing.T) {
	for i, tt := range readResponseTests {
		r := reader(tt.in+"\nFOLLOWING DATA", &bytes.Buffer{})
		code, msg, err := r.ReadResponse(tt.inCode)
		if err != nil {
			t.Errorf("#%d: ReadResponse: %v", i, err)
			continue
		}
		if code != tt.wantCode {
			t.Errorf("#%d: code=%d, want %d", i, code, tt.wantCode)
		}
		if msg != tt.wantMsg {
			t.Errorf("#%d: msg=%q, want %q", i, msg, tt.wantMsg)
		}
	}
}

// Test that multi-line errors are appropriately and fully read. Issue 10230.
func TestReadMultiLineError(t *testing.T) {
	r := reader("550-5.1.1 The email account that you tried to reach does not exist. Please try\n"+
		"550-5.1.1 double-checking the recipient's email address for typos or\n"+
		"550-5.1.1 unnecessary spaces. Learn more at\n"+
		"Unexpected but legal text!\n"+
		"550 5.1.1 https://support.google.com/mail/answer/6596 h20si25154304pfd.166 - gsmtp\n", &bytes.Buffer{})

	wantMsg := "5.1.1 The email account that you tried to reach does not exist. Please try\n" +
		"5.1.1 double-checking the recipient's email address for typos or\n" +
		"5.1.1 unnecessary spaces. Learn more at\n" +
		"Unexpected but legal text!\n" +
		"5.1.1 https://support.google.com/mail/answer/6596 h20si25154304pfd.166 - gsmtp"

	code, msg, err := r.ReadResponse(250)
	if err == nil {
		t.Errorf("ReadResponse: no error, want error")
	}
	if code != 550 {
		t.Errorf("ReadResponse: code=%d, want %d", code, 550)
	}
	if msg != wantMsg {
		t.Errorf("ReadResponse: msg=%q, want %q", msg, wantMsg)
	}
	if err != nil && err.Error() != "550 "+wantMsg {
		t.Errorf("ReadResponse: error=%q, want %q", err.Error(), "550 "+wantMsg)
	}
}
