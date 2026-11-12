package smtpreader

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
		reader := NewBdat(size, last, 0, readerInput, func() (string, string, error) {
			return "BDAT", "4 FIRST", nil
		})
		_, err = io.ReadAll(reader)
		require.ErrorContains(t, err, "501")

		size, last, err = BdatArg(arg[5:])
		require.NoError(t, err)
		reader = NewBdat(size, last, 0, readerInput, func() (string, string, error) {
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
		reader := NewBdat(size, last, 0, readerInput, func() (string, string, error) {
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
		reader := NewBdat(size, last, 0, readerInput, func() (string, string, error) {
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
		reader := NewBdat(size, last, 4, readerInput, func() (string, string, error) {
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
		reader := NewBdat(size, last, 6, readerInput, func() (string, string, error) {
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
		reader := NewBdat(size, last, 4, readerInput, func() (string, string, error) {
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
		reader := NewBdat(size, last, 4, readerInput, func() (string, string, error) {
			return "", "", smtp.ErrAuthFailed
		})
		require.NoError(t, err)
		_, err = io.ReadAll(reader)
		require.ErrorIs(t, smtp.ErrAuthFailed, err)
	})
}

// The server relies on the limit error being sticky to detect a breach the
// backend swallowed: every further Read must fail without consuming input.
func TestBdatReaderLimitSticky(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		read  string
		rest  string
	}{
		{name: "FirstChunk", input: "BDAT 5 LAST\r\nhappyNOOP\r\n", read: "", rest: "happyNOOP\r\n"},
		{name: "SecondChunk", input: "BDAT 3\r\nhapBDAT 2 LAST\r\npyNOOP\r\n", read: "hap", rest: "pyNOOP\r\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			readerInput := bufio.NewReader(strings.NewReader(tc.input))
			line, _, err := readerInput.ReadLine()
			require.NoError(t, err)
			size, last, err := BdatArg(string(line[5:]))
			require.NoError(t, err)

			reader := NewBdat(size, last, 4, readerInput, func() (string, string, error) {
				line, _, err := readerInput.ReadLine()
				if err != nil {
					return "", "", err
				}
				return string(line[0:4]), string(line[5:]), nil
			})

			res, err := io.ReadAll(reader)
			require.ErrorIs(t, err, smtp.ErrDataTooLarge)
			require.Equal(t, tc.read, string(res))

			for range 3 {
				n, err := reader.Read(make([]byte, 10))
				require.Zero(t, n)
				require.ErrorIs(t, err, smtp.ErrDataTooLarge)
			}

			// the chunk is still pending on the input
			rest, err := io.ReadAll(readerInput)
			require.NoError(t, err)
			require.Equal(t, tc.rest, string(rest))
		})
	}
}
