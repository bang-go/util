package util

import (
	"errors"
	"io"
	"os"
	"path"
	"runtime"
)

type PathWrapper struct {
	File string
	Dir  string
	Line int
}

func Path() *PathWrapper {
	return PathWithCallerSkip(2)
}

func PathWithCallerSkip(skip int) *PathWrapper {
	_, file, line, _ := runtime.Caller(skip)
	return &PathWrapper{
		File: file,
		Dir:  path.Dir(file),
		Line: line,
	}
}

func FilePosition(file *os.File) (int64, error) {
	if file == nil {
		return 0, errors.New("null fd when retrieving file position")
	}
	return file.Seek(0, io.SeekCurrent)
}

func FileExist(name string) (b bool, err error) {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
	}
	// Propagates the error if the error is not FileNotExist error.
	return true, err
}

func CreateDirIfNotExist(dirname string) error {
	if _, err := os.Stat(dirname); err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(dirname, os.ModePerm)
		} else {
			return err
		}
	}
	return nil
}
