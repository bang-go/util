package util

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

// ErrNilFile indicates that a file helper received a nil file handle.
var ErrNilFile = errors.New("util: nil file")

const ensureDirPerm = 0o755

// CallerInfo describes a caller source location.
type CallerInfo struct {
	File string
	Dir  string
	Line int
}

// Caller returns the caller source location.
func Caller() CallerInfo {
	return CallerSkip(2)
}

// CallerSkip returns caller source location for the given runtime.Caller skip.
func CallerSkip(skip int) CallerInfo {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return CallerInfo{}
	}

	return CallerInfo{
		File: file,
		Dir:  filepath.Dir(file),
		Line: line,
	}
}

// FilePosition returns the current offset of file.
func FilePosition(file *os.File) (int64, error) {
	if file == nil {
		return 0, ErrNilFile
	}
	return file.Seek(0, io.SeekCurrent)
}

// FileExists reports whether the named path exists.
// It returns false, nil when the path does not exist.
func FileExists(name string) (bool, error) {
	if _, err := os.Stat(name); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// EnsureDir creates dirname and its parents when they do not exist.
func EnsureDir(dirname string) error {
	return os.MkdirAll(dirname, ensureDirPerm)
}
