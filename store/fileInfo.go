package store

import (
	"os"
	"time"
)

// FileInfo is an implementation to store os.FileInfo. Warning: Sys is always nil
type FileInfo struct {
	NameField    string      `json:"name"`
	SizeField    int64       `json:"size"`
	ModeField    os.FileMode `json:"mode"`
	ModTimeField time.Time   `json:"modTime"`
}

func (f FileInfo) Name() string {
	return f.NameField
}

func (f FileInfo) Size() int64 {
	return f.SizeField
}

func (f FileInfo) Mode() os.FileMode {
	return f.ModeField
}

func (f FileInfo) ModTime() time.Time {
	return f.ModTimeField
}

func (f FileInfo) IsDir() bool {
	return f.ModeField.IsDir()
}

func (f FileInfo) Sys() interface{} {
	return nil
}

func CopyToFileInfo(info os.FileInfo) FileInfo {
	return FileInfo{
		NameField:    info.Name(),
		SizeField:    info.Size(),
		ModeField:    info.Mode(),
		ModTimeField: info.ModTime(),
	}
}
