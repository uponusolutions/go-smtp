package textsmtp

import (
	"bufio"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uponusolutions/go-smtp"
)

func TestBdatReaderArgErrors(t *testing.T) {
	t.Run("NoArguments", func(t *testing.T) {
		_, _, err := BdatArg("")
		require.ErrorContains(t, err, "501")
	})

	t.Run("TooManyArguments", func(t *testing.T) {
		_, _, err := BdatArg("5 b c")
		require.ErrorContains(t, err, "501")
	})

	t.Run("InvalidArguments", func(t *testing.T) {
		_, _, err := BdatArg("5 FIRST")
		require.ErrorContains(t, err, "501")

		_, _, err = BdatArg("-1")
		require.ErrorContains(t, err, "501")
	})

	t.Run("InvalidArgAtStart", func(t *testing.T) {
		input := "BDAT 3\r\nhapBDAT 0 LAST\r\n"
		readerInput := bufio.NewReader(strings.NewReader(input))

		_, prefix, err := readerInput.ReadLine()
		require.NoError(t, err)
		require.False(t, prefix)

		_, _, err = BdatArg("4 FIRST")
		require.ErrorContains(t, err, "501")
	})

	t.Run("InvalidArgNextCommand", func(t *testing.T) {
		input := "BDAT 3\r\nhapBDAT 0 LAST\r\n"
		readerInput := bufio.NewReader(strings.NewReader(input))

		byteArg, prefix, err := readerInput.ReadLine()
		arg := string(byteArg)
		require.NoError(t, err)
		require.False(t, prefix)
		size, last, err := BdatArg(arg[5:])
		require.NoError(t, err)
		reader := NewBdatReader(size, last, 0, readerInput, func() (string, string, error) {
			return "BDAT", "4 FIRST", nil
		})
		_, err = io.ReadAll(reader)
		require.ErrorContains(t, err, "501")

		size, last, err = BdatArg(arg[5:])
		require.NoError(t, err)
		reader = NewBdatReader(size, last, 0, readerInput, func() (string, string, error) {
			return "STRANGE", "4 LAST", nil
		})
		_, err = io.ReadAll(reader)
		require.ErrorContains(t, err, "501")
	})
}

func TestBdatReader(t *testing.T) {
	t.Run("LastWithoutData", func(t *testing.T) {
		input := "BDAT 3\r\nhapBDAT 0 LAST\r\n"
		readerInput := bufio.NewReader(strings.NewReader(input))

		byteArg, prefix, err := readerInput.ReadLine()
		arg := string(byteArg)
		require.NoError(t, err)
		require.False(t, prefix)

		size, last, err := BdatArg(arg[5:])
		require.NoError(t, err)
		reader := NewBdatReader(size, last, 0, readerInput, func() (string, string, error) {
			byteArg, prefix, err := readerInput.ReadLine()
			arg := string(byteArg)

			require.NoError(t, err)
			require.False(t, prefix)
			return arg[0:4], arg[5:], nil
		})
		res, err := io.ReadAll(reader)
		require.NoError(t, err)

		require.Equal(t, "hap", string(res))
	})

	t.Run("NoChunkingLimit", func(t *testing.T) {
		input := "BDAT 3\r\nhapBDAT 2 LAST\r\npy"
		readerInput := bufio.NewReader(strings.NewReader(input))

		byteArg, prefix, err := readerInput.ReadLine()
		arg := string(byteArg)
		require.NoError(t, err)
		require.False(t, prefix)

		size, last, err := BdatArg(arg[5:])
		require.NoError(t, err)
		reader := NewBdatReader(size, last, 0, readerInput, func() (string, string, error) {
			byteArg, prefix, err := readerInput.ReadLine()
			arg := string(byteArg)

			require.NoError(t, err)
			require.False(t, prefix)
			return arg[0:4], arg[5:], nil
		})
		res, err := io.ReadAll(reader)
		require.NoError(t, err)

		require.Equal(t, "happy", string(res))
	})

	t.Run("MaxMessageSizeExceeded", func(t *testing.T) {
		input := "BDAT 3\r\nhapBDAT 2 LAST\r\npy"
		readerInput := bufio.NewReader(strings.NewReader(input))

		byteArg, prefix, err := readerInput.ReadLine()
		arg := string(byteArg)
		require.NoError(t, err)
		require.False(t, prefix)

		size, last, err := BdatArg(arg[5:])
		require.NoError(t, err)
		reader := NewBdatReader(size, last, 4, readerInput, func() (string, string, error) {
			byteArg, prefix, err := readerInput.ReadLine()
			arg := string(byteArg)

			require.NoError(t, err)
			require.False(t, prefix)
			return arg[0:4], arg[5:], nil
		})
		_, err = io.ReadAll(reader)
		require.ErrorIs(t, err, smtp.ErrDataTooLarge)
	})

	t.Run("ToFewBytes", func(t *testing.T) {
		input := "BDAT 3\r\nhapBDAT 2 LAST\r\np"
		readerInput := bufio.NewReader(strings.NewReader(input))

		byteArg, prefix, err := readerInput.ReadLine()
		arg := string(byteArg)
		require.NoError(t, err)
		require.False(t, prefix)

		size, last, err := BdatArg(arg[5:])
		require.NoError(t, err)
		reader := NewBdatReader(size, last, 6, readerInput, func() (string, string, error) {
			byteArg, prefix, err := readerInput.ReadLine()
			arg := string(byteArg)

			require.NoError(t, err)
			require.False(t, prefix)
			return arg[0:4], arg[5:], nil
		})
		_, err = io.ReadAll(reader)
		require.ErrorIs(t, err, smtp.ErrConnection)
	})

	t.Run("NextCommandErrorEOF", func(t *testing.T) {
		input := "BDAT 3\r\nhapBDAT 2 LAST\r\npy"
		readerInput := bufio.NewReader(strings.NewReader(input))

		byteArg, prefix, err := readerInput.ReadLine()
		arg := string(byteArg)
		require.NoError(t, err)
		require.False(t, prefix)

		size, last, err := BdatArg(arg[5:])
		require.NoError(t, err)
		reader := NewBdatReader(size, last, 4, readerInput, func() (string, string, error) {
			return "", "", io.EOF
		})
		_, err = io.ReadAll(reader)
		require.ErrorIs(t, smtp.ErrConnection, err)
	})

	t.Run("NextCommandErrorRandomSmtpError", func(t *testing.T) {
		input := "BDAT 3\r\nhapBDAT 2 LAST\r\npy"
		readerInput := bufio.NewReader(strings.NewReader(input))

		byteArg, prefix, err := readerInput.ReadLine()
		arg := string(byteArg)
		require.NoError(t, err)
		require.False(t, prefix)

		size, last, err := BdatArg(arg[5:])
		require.NoError(t, err)
		reader := NewBdatReader(size, last, 4, readerInput, func() (string, string, error) {
			return "", "", smtp.ErrAuthFailed
		})
		require.NoError(t, err)
		_, err = io.ReadAll(reader)
		require.ErrorIs(t, smtp.ErrAuthFailed, err)
	})
}
