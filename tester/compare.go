package tester

import (
	"bytes"
	"embed"
	"io"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func getAllFilenames(fs *embed.FS, path string) (out []string, err error) {
	if len(path) == 0 {
		path = "."
	}
	entries, err := fs.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		fp := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			res, err := getAllFilenames(fs, fp)
			if err != nil {
				return nil, err
			}
			out = append(out, res...)
			continue
		}
		out = append(out, fp)
	}
	return out, err
}

func chunkSlice(slice []byte, chunkSize int) [][]byte {
	var chunks [][]byte
	for len(slice) != 0 {
		// necessary check to avoid slicing beyond
		// slice capacity
		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}

func checkExpectedBufferAgainsActual(t *testing.T, b []byte, expected func(io.Writer) io.WriteCloser, actual func(io.Writer) io.WriteCloser) {
	var buf bytes.Buffer
	var err error

	f := expected(&buf)
	_, err = f.Write(b)
	require.NoError(t, err)
	require.NoError(t, f.Close())

	size := 1

	for size < 4048 && len(b) >= size {
		bsplitted := chunkSlice(b, size)

		var buf1 bytes.Buffer
		f := actual(&buf1)

		for _, b := range bsplitted {
			_, err = f.Write(b)
			require.NoError(t, err)
		}

		require.NoError(t, f.Close())

		require.Equal(t, buf, buf1)

		size++
	}
}

// WriterCompareTest reads all files out of fs[path] and compares the expected func against the actual func.
// To simulate differences of Write calls of different sizes it slices the files in increasing sizes up to 4048.
func WriterCompareTest(t *testing.T, fs *embed.FS, path string, expected func(io.Writer) io.WriteCloser, actual func(io.Writer) io.WriteCloser) {
	files, err := getAllFilenames(fs, path)
	require.NoError(t, err)

	for _, file := range files {
		dat, err := fs.ReadFile(file)
		require.NoError(t, err)
		checkExpectedBufferAgainsActual(t, []byte(dat), expected, actual)
	}
}

func checkRaderExpectedAgainsActual(t *testing.T, b []byte, expected func(io.Reader) ([]byte, error), actual func(io.Reader) ([]byte, error)) {
	pr, pw := io.Pipe()

	go func() {
		_, err := pw.Write(b)
		require.NoError(t, err)
		err = pw.Close()
		require.NoError(t, err)
	}()
	buf, err := expected(pr)
	require.ErrorIs(t, io.ErrUnexpectedEOF, err)

	size := 1
	for size < 4048 && len(b) >= size {
		bsplitted := chunkSlice(b, size)

		pr, pw = io.Pipe()

		writeInGoroutine(t, bsplitted, pw)

		buf1, err := actual(pr)
		require.ErrorIs(t, io.ErrUnexpectedEOF, err)
		// print(string(buf), string(buf1))
		require.Equal(t, buf, buf1)

		size++
	}
}

func writeInGoroutine(t *testing.T, bsplitted [][]byte, pw *io.PipeWriter) {
	go func() {
		var err error
		for _, b := range bsplitted {
			_, err = pw.Write(b)
			require.NoError(t, err)
		}
		err = pw.Close()
		require.NoError(t, err)
	}()
}

// ReaderCompareTest reads all files out of fs[path] and compares the result of the expected func against the actual func.
// To simulate differences of Read calls with different sizes it slices the files in increasing sizes up to 4048.
func ReaderCompareTest(t *testing.T, fs *embed.FS, path string, expected func(io.Reader) ([]byte, error), actual func(io.Reader) ([]byte, error)) {
	files, err := getAllFilenames(fs, path)
	require.NoError(t, err)

	for _, file := range files {
		dat, err := fs.ReadFile(file)
		require.NoError(t, err)
		checkRaderExpectedAgainsActual(t, []byte(dat), expected, actual)
	}
}
