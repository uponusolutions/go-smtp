package textsmtp

import (
	"bytes"
	"io"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uponusolutions/go-smtp"
	"github.com/uponusolutions/go-smtp/tester"
)

func reader(in string, out *bytes.Buffer) *Textproto {
	return NewTextproto(tester.NewFakeConn(in, out), 4096, 4096, 0)
}

func TestPrintfLine(t *testing.T) {
	buf := &bytes.Buffer{}
	w := reader("", buf)
	err := w.PrintfLineAndFlush("foo %d", 123)
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

type parseFirstCodeLineTest struct {
	codeLine        string
	wantCode        int
	wantContinued   bool
	wantEnhCode     smtp.EnhancedCode
	wantMsg         string
	wantErrContains string
}

var parseFirstCodeLineTests = []parseFirstCodeLineTest{
	{
		"123 test",
		123,
		false,
		smtp.NoEnhancedCode,
		"test",
		"",
	},
	{
		"123-test",
		123,
		true,
		smtp.NoEnhancedCode,
		"test",
		"",
	},
	{
		"123-1.2 test",
		123,
		true,
		smtp.NoEnhancedCode,
		"1.2 test",
		"",
	},
	{
		"123-1.2.3 test",
		123,
		true,
		smtp.EnhancedCode{1, 2, 3},
		"test",
		"",
	},
	{
		"323-1.2.3 test",
		323,
		true,
		smtp.NoEnhancedCode,
		"1.2.3 test",
		"",
	},
}

func TestParseFirstCodeLine(t *testing.T) {
	for _, d := range parseFirstCodeLineTests {
		t.Run(d.codeLine, func(t *testing.T) {
			status, continued, err := parseFirstCodeLine(d.codeLine)
			require.Equal(t, d.wantCode, status.Code)
			require.Equal(t, d.wantContinued, continued)
			require.Equal(t, d.wantEnhCode, status.EnhancedCode)
			require.Equal(t, []string{d.wantMsg}, status.Lines)
			if d.wantErrContains == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, d.wantErrContains)
			}
		})
	}
}

type parseExtraCodeLineTest struct {
	line             string
	codeString       string
	enhancedCodePart []byte
	wantContinued    bool
	wantMsg          string
	wantErrContains  string
}

var parseExtraCodeLineTests = []parseExtraCodeLineTest{
	{
		"123-1.1.0 test",
		"123",
		[]byte("1.1.0 "),
		true,
		"test",
		"",
	},
	{
		"123 1.2.0 test",
		"123",
		[]byte("1.2.0 "),
		false,
		"test",
		"",
	},
	{
		"123 test",
		"123",
		nil,
		false,
		"test",
		"",
	},
}

func TestParseExtraCodeLine(t *testing.T) {
	for i, d := range parseExtraCodeLineTests {
		t.Run(strconv.Itoa(i)+": "+d.line, func(t *testing.T) {
			continued, msg, err := parseExtraCodeLine(d.line, d.codeString, d.enhancedCodePart)
			require.Equal(t, d.wantContinued, continued)
			require.Equal(t, d.wantMsg, msg)
			if d.wantErrContains == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, d.wantErrContains)
			}
		})
	}
}

type readResponseTest struct {
	in       string
	wantCode int
	wantMsg  string
}

var readResponseTests = []readResponseTest{
	/*
		 	// invalid by RFC 5321
			{
				"230-Anonymous access granted, restrictions apply\n" +
					"Read the file README.txt,\n" +
					"230  please",
				230,
				"Anonymous access granted, restrictions apply\nRead the file README.txt,\n please",
			},
	*/
	{
		"230 Anonymous access granted, restrictions apply\n",
		230,
		"Anonymous access granted, restrictions apply",
	},

	{
		"400-A\n400-B\n400 C",
		400,
		"A\nB\nC",
	},

	{
		"400-A\r\n400-B\r\n400 C\r\n",
		400,
		"A\nB\nC",
	},
}

// See https://www.ietf.org/rfc/rfc959.txt page 36.
func TestRFC959Lines(t *testing.T) {
	for i, tt := range readResponseTests {
		r := reader(tt.in+"\nFOLLOWING DATA", &bytes.Buffer{})
		status, err := r.ReadResponse()
		if err != nil {
			t.Errorf("#%d: ReadResponse: %v", i, err)
			continue
		}
		if status.Code != tt.wantCode {
			t.Errorf("#%d: code=%d, want %d", i, status.Code, tt.wantCode)
		}
		if strings.Join(status.Lines, "\n") != tt.wantMsg {
			t.Errorf("#%d: msg=%q, want %q", i, strings.Join(status.Lines, "\n"), tt.wantMsg)
		}
	}
}

// Test that multi-line errors are appropriately and fully read.
func TestReadMultiLineError(t *testing.T) {
	r := reader("550-5.1.1 The email account that you tried to reach does not exist. Please try\n"+
		"550-5.1.1 double-checking the recipient's email address for typos or\n"+
		"550-5.1.1 unnecessary spaces. Learn more at\n"+
		"Unexpected but legal text!\n"+
		"550 5.1.1 https://support.google.com/mail/answer/6596 h20si25154304pfd.166 - gsmtp\n", &bytes.Buffer{})

	err := r.ReadResponseValid(250)
	if err == nil {
		t.Error("ReadResponse: no error, want error")
	}

	require.ErrorContains(t, err, "invalid response")
}

// Test that multi-line errors are appropriately and fully read.
func TestReadMultiLine(t *testing.T) {
	r := reader("550-5.1.1 The email account that you tried to reach does not exist. Please try\n"+
		"550-5.1.1 double-checking the recipient's email address for typos or\n"+
		"550-5.1.1 unnecessary spaces. Learn more at\n"+
		"550 5.1.1 https://support.google.com/mail/answer/6596 h20si25154304pfd.166 - gsmtp\n", &bytes.Buffer{})

	expectedText := "The email account that you tried to reach does not exist. Please try\n" +
		"double-checking the recipient's email address for typos or\n" +
		"unnecessary spaces. Learn more at\n" +
		"https://support.google.com/mail/answer/6596 h20si25154304pfd.166 - gsmtp"

	err := r.ReadResponseValid(250)
	if err == nil {
		t.Error("ReadResponse: no error, want error")
	}

	status, ok := err.(*smtp.Status)
	require.True(t, ok)

	require.Equal(t, expectedText, status.Text())
	require.Equal(t, "SMTP error 550 5.1.1: "+expectedText, status.Error())
}
