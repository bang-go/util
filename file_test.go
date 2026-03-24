package util_test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/bang-go/util"
)

func TestPath(t *testing.T) {
	p := helperPath()
	if p.File == "" || p.Dir == "" || p.Line == 0 {
		t.Fatalf("Caller() returned incomplete data: %+v", p)
	}
	if filepath.Base(p.File) != "file_test.go" {
		t.Fatalf("Caller().File = %q", p.File)
	}
	if p.Dir != filepath.Dir(p.File) {
		t.Fatalf("Caller().Dir = %q, want %q", p.Dir, filepath.Dir(p.File))
	}
}

func helperPath() util.CallerInfo {
	return util.Caller()
}

func TestFilePosition(t *testing.T) {
	if _, err := util.FilePosition(nil); !errors.Is(err, util.ErrNilFile) {
		t.Fatalf("FilePosition(nil) error = %v", err)
	}

	file, err := os.CreateTemp(t.TempDir(), "position-*")
	if err != nil {
		t.Fatalf("CreateTemp() error = %v", err)
	}
	defer file.Close()

	if _, err := file.WriteString("hello"); err != nil {
		t.Fatalf("WriteString() error = %v", err)
	}

	pos, err := util.FilePosition(file)
	if err != nil {
		t.Fatalf("FilePosition() error = %v", err)
	}
	if pos != 5 {
		t.Fatalf("FilePosition() = %d, want 5", pos)
	}
}

func TestFileExists(t *testing.T) {
	existingFile := filepath.Join(t.TempDir(), "existing.txt")
	if err := os.WriteFile(existingFile, []byte("ok"), 0o644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	exists, err := util.FileExists(existingFile)
	if err != nil || !exists {
		t.Fatalf("FileExists(existing) = (%v, %v)", exists, err)
	}

	exists, err = util.FileExists(filepath.Join(t.TempDir(), "missing.txt"))
	if err != nil || exists {
		t.Fatalf("FileExists(missing) = (%v, %v)", exists, err)
	}

}

func TestEnsureDir(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "a", "b", "c")
	if err := util.EnsureDir(dir); err != nil {
		t.Fatalf("EnsureDir() error = %v", err)
	}
	if info, err := os.Stat(dir); err != nil || !info.IsDir() {
		t.Fatalf("Stat() = (%v, %v)", info, err)
	}
	if err := util.EnsureDir(dir); err != nil {
		t.Fatalf("EnsureDir(existing) error = %v", err)
	}
}
