package tree

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type ForEach func(path string, size int64)

type traverser struct {
	forEach ForEach
}

// NewTraverser creates a new traverser with specified func applied to each file, nil is allowed as func.
func NewTraverser(forEach ForEach) traverser {
	return traverser{
		forEach: forEach,
	}
}

// Traverse the given path and return a map of full filenames and FileInfos of each of them
func (t traverser) Traverse(path string) (map[string]os.FileInfo, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("getting stat error: %w", err)
	}

	return t.traverseFileInfo(stat, path)
}

func (t traverser) traverseFileInfo(fileInfo os.FileInfo, fileInfoPath string) (map[string]os.FileInfo, error) {
	if t.forEach != nil {
		t.forEach(fileInfoPath, fileInfo.Size())
	}

	result := make(map[string]os.FileInfo)
	result[fileInfoPath] = fileInfo

	if fileInfo.IsDir() {
		fileInfos, err := ioutil.ReadDir(fileInfoPath)
		if err != nil {
			return nil, fmt.Errorf("reading dir error: %w", err)
		}

		for _, v := range fileInfos {
			content, err := t.traverseFileInfo(v, path.Join(fileInfoPath, v.Name()))
			if err != nil {
				return nil, err
			}
			appendMap(result, content)
		}
	}

	return result, nil
}

func appendMap(to map[string]os.FileInfo, from map[string]os.FileInfo) {
	for k, v := range from {
		to[k] = v
	}
}
