package smtpreader

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

func reader(in string, out *bytes.Buffer) *Receiver {
	return NewReceiver(tester.NewFakeConn(in, out), 4096, 0, 0)
}

func TestReadLine(t *testing.T) {
	r := reader("line1\nline2\n", &bytes.Buffer{})
	s, err := r.ReadFullLine()
	if s != "line1" || err != nil {
		t.Fatalf("Line 1: %s, %v", s, err)
	}
	s, err = r.ReadFullLine()
	if s != "line2" || err != nil {
		t.Fatalf("Line 2: %s, %v", s, err)
	}
	s, err = r.ReadFullLine()
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
		// RFC 2034: the class must match the reply code and be 2, 4 or 5
		"123-1.2.3 test",
		123,
		true,
		smtp.NoEnhancedCode,
		"1.2.3 test",
		"",
	},
	{
		"523-5.2.3 test",
		523,
		true,
		smtp.EnhancedCode{5, 2, 3},
		"test",
		"",
	},
	{
		// RFC 3463: subject and detail have up to 3 digits
		"550-5.7.26 test",
		550,
		true,
		smtp.EnhancedCode{5, 7, 26},
		"test",
		"",
	},
	{
		"451 4.123.999 test",
		451,
		false,
		smtp.EnhancedCode{4, 123, 999},
		"test",
		"",
	},
	{
		"550 4.7.1 class mismatch",
		550,
		false,
		smtp.NoEnhancedCode,
		"4.7.1 class mismatch",
		"",
	},
	{
		"550 5.1234.1 too many digits",
		550,
		false,
		smtp.NoEnhancedCode,
		"5.1234.1 too many digits",
		"",
	},
	{
		"550 5.07.1 leading zero",
		550,
		false,
		smtp.NoEnhancedCode,
		"5.07.1 leading zero",
		"",
	},
	{
		"550 5.+7.1 sign",
		550,
		false,
		smtp.NoEnhancedCode,
		"5.+7.1 sign",
		"",
	},
	{
		"550 55.7.1 class with two digits",
		550,
		false,
		smtp.NoEnhancedCode,
		"55.7.1 class with two digits",
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
	{
		"550-5.7.26 test",
		"550",
		smtp.EnhancedCode{5, 7, 26}.ToPart(),
		true,
		"test",
		"",
	},
	{
		// RFC 2034: the same enhanced code on every line
		"550 5.7.2 test",
		"550",
		smtp.EnhancedCode{5, 7, 26}.ToPart(),
		false,
		"",
		"invalid response",
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
	require.Equal(t, "550 5.1.1 "+expectedText, status.Error())
}

// TestReadMultiLineMultiDigitEnhancedCode reads a Gmail style reply whose
// enhanced code has a two digit detail. The whole reply has to be consumed,
// otherwise the next reply is read off by one.
func TestReadMultiLineMultiDigitEnhancedCode(t *testing.T) {
	r := reader("550-5.7.26 Unauthenticated email from example.com is not accepted due to\r\n"+
		"550-5.7.26 domain's DMARC policy. Please contact the administrator of\r\n"+
		"550 5.7.26 example.com domain if this was a legitimate mail. - gsmtp\r\n"+
		"250 2.0.0 OK\r\n", &bytes.Buffer{})

	err := r.ReadResponseValid(250)
	var status *smtp.Status
	require.ErrorAs(t, err, &status)
	require.Equal(t, 550, status.Code)
	require.Equal(t, smtp.EnhancedCode{5, 7, 26}, status.EnhancedCode)
	require.Len(t, status.Lines, 3)
	require.Equal(t, "example.com domain if this was a legitimate mail. - gsmtp", status.Lines[2])

	require.NoError(t, r.ReadResponseValid(250), "next reply must be read in sync")
}

func TestReadResponseTooManyLines(t *testing.T) {
	const maxResponseLines = 16
	limited := func(in string) *Receiver {
		return NewReceiver(tester.NewFakeConn(in, &bytes.Buffer{}), 4096, 0, maxResponseLines)
	}

	in := strings.Repeat("250-ext\r\n", maxResponseLines) + "250 last\r\n"
	_, err := limited(in).ReadResponse()
	require.ErrorIs(t, err, ErrTooManyResponseLines)

	in = strings.Repeat("250-ext\r\n", maxResponseLines-1) + "250 last\r\n"
	status, err := limited(in).ReadResponse()
	require.NoError(t, err)
	require.Len(t, status.Lines, maxResponseLines)

	// ReadResponseValid is limited as well
	in = strings.Repeat("250-ext\r\n", maxResponseLines) + "250 last\r\n"
	require.ErrorIs(t, limited(in).ReadResponseValid(250), ErrTooManyResponseLines)

	// zero disables the limit
	in = strings.Repeat("250-ext\r\n", 10*maxResponseLines) + "250 last\r\n"
	status, err = reader(in, &bytes.Buffer{}).ReadResponse()
	require.NoError(t, err)
	require.Len(t, status.Lines, 10*maxResponseLines+1)
}
