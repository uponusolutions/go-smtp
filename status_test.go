package smtp

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// RFC 3463: status-code = class "." subject "." detail, subject and detail have 1*3digit.
func TestEnhancedCodeToPart(t *testing.T) {
	assert.Nil(t, NoEnhancedCode.ToPart())
	assert.Equal(t, "2.0.0 ", string(EnhancedCode{2, 0, 0}.ToPart()))
	assert.Equal(t, "5.7.26 ", string(EnhancedCode{5, 7, 26}.ToPart()))
	assert.Equal(t, "4.123.999 ", string(EnhancedCode{4, 123, 999}.ToPart()))
}

// RFC 2034: the same enhanced code at the beginning of the text of each line.
func TestStatusWriteToMultiDigitEnhancedCode(t *testing.T) {
	var out bytes.Buffer
	w := bufio.NewWriter(&out)

	s := NewStatusM(550, EnhancedCode{5, 7, 26}, []string{"first", "last"})
	n, err := s.WriteTo(w)
	require.NoError(t, err)
	require.NoError(t, w.Flush())

	assert.Equal(t, "550-5.7.26 first\r\n550 5.7.26 last\r\n", out.String())
	assert.Equal(t, int64(out.Len()), n)
}
