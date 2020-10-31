package store

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

var testData = map[string]os.FileInfo{
	"a":          FileInfo{NameField: "a field name", SizeField: 184, ModeField: os.ModeDir},
	"/full/path": FileInfo{NameField: "test name", SizeField: 22000},
}

func TestStore(t *testing.T) {
	dbFileName := getTempFileName(t)
	defer os.Remove(dbFileName)

	err := Write(dbFileName, testData)
	require.NoError(t, err)

	readed, err := Read(dbFileName)
	require.NoError(t, err)

	for k, v := range testData {
		readedFileInfo, ok := readed[k]
		require.True(t, ok)

		assert.Equal(t, v, readedFileInfo)
	}
}

func getTempFileName(t *testing.T) string {
	tempFile, err := ioutil.TempFile("", "")
	require.NoError(t, err)
	err = tempFile.Close()
	require.NoError(t, err)
	err = os.Remove(tempFile.Name())
	require.NoError(t, err)

	return tempFile.Name()
}
