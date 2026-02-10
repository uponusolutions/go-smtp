package tester

import (
	"embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/*
var embedFS embed.FS

func TestCompare(t *testing.T) {
	files, err := getAllFilenames(&embedFS, "testdata/a")
	require.NoError(t, err)
	require.Equal(t, []string{"testdata/a/b.txt", "testdata/a/c.txt"}, files)

	files, err = getAllFilenames(&embedFS, "testdata/d")
	require.NoError(t, err)
	require.Equal(t, []string{"testdata/d/e.txt", "testdata/d/f/g.txt"}, files)

	_, err = getAllFilenames(&embedFS, "testdata/e")
	require.Error(t, err)
}
