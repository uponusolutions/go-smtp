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

func checkExpectedAgainsActual(t *testing.T, b []byte, expected func(*bytes.Buffer) io.WriteCloser, actual func(*bytes.Buffer) io.WriteCloser) {
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

// BufferCompareTest reads all files out of fs[path] and comapres the expected func against the actual func.
// To simulate differences of Write calls of differnt sizes it slices the files in increasing sizes up to 4048.
func BufferCompareTest(t *testing.T, fs *embed.FS, path string, expected func(*bytes.Buffer) io.WriteCloser, actual func(*bytes.Buffer) io.WriteCloser) {
	files, err := getAllFilenames(fs, path)
	require.NoError(t, err)

	for _, file := range files {
		dat, err := fs.ReadFile(file)
		require.NoError(t, err)
		checkExpectedAgainsActual(t, []byte(dat), expected, actual)
	}
}
