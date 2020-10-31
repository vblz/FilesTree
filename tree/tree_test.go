package tree

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func Test_traverser_Traverse_NotExisting(t *testing.T) {
	_, err := NewTraverser(nil).Traverse("/not_existing_path")
	assert.Error(t, err)
}

func Test_traverser_Traverse_EmptyDir(t *testing.T) {
	dirPath, err := ioutil.TempDir("", "")
	require.NoError(t, err)

	defer os.RemoveAll(dirPath)

	res, err := NewTraverser(nil).Traverse(dirPath)
	require.NoError(t, err)
	assert.Len(t, res, 1)

	dirInfo, ok := res[dirPath]
	require.True(t, ok)
	assert.Equal(t, path.Base(dirPath), dirInfo.Name())
	assert.True(t, dirInfo.IsDir())
}

func Test_traverser_Traverse(t *testing.T) {
	dirPath, err := ioutil.TempDir("", "")
	require.NoError(t, err)

	defer os.RemoveAll(dirPath)

	const rootFileSize = 1256
	const rootFileName = "testFileName"
	var rootFilePath = path.Join(dirPath, rootFileName)
	err = ioutil.WriteFile(rootFilePath, bytes.Repeat([]byte{11}, rootFileSize), 666)
	require.NoError(t, err)

	const subDirName = "subDirTestName"
	var subDirPath = path.Join(dirPath, subDirName)
	const subDirFileSize = 14
	const subDirFileName = "name_of_subdir_file"
	var subDirFilePath = path.Join(subDirPath, subDirFileName)
	err = os.Mkdir(subDirPath, 0777)
	require.NoError(t, err)
	err = ioutil.WriteFile(subDirFilePath, bytes.Repeat([]byte{4}, subDirFileSize), 666)
	require.NoError(t, err)

	res, err := NewTraverser(nil).Traverse(dirPath)
	require.NoError(t, err)
	assert.Len(t, res, 4)

	dirInfo, ok := res[dirPath]
	require.True(t, ok)
	assert.Equal(t, path.Base(dirPath), dirInfo.Name())
	assert.True(t, dirInfo.IsDir())

	subDirInfo, ok := res[subDirPath]
	require.True(t, ok)
	assert.Equal(t, subDirName, subDirInfo.Name())
	assert.True(t, subDirInfo.IsDir())

	rootFileInfo, ok := res[rootFilePath]
	require.True(t, ok)
	assert.Equal(t, rootFileName, rootFileInfo.Name())
	assert.Equal(t, int64(rootFileSize), rootFileInfo.Size())
	assert.False(t, rootFileInfo.IsDir())

	subDirFileInfo, ok := res[subDirFilePath]
	require.True(t, ok)
	assert.Equal(t, subDirFileName, subDirFileInfo.Name())
	assert.Equal(t, int64(subDirFileSize), subDirFileInfo.Size())
	assert.False(t, subDirFileInfo.IsDir())
}
