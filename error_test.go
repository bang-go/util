package util_test

import (
	"bytes"
	"errors"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/bang-go/util"
)

func TestRunAndRecover(t *testing.T) {
	originalWriter := log.Writer()
	originalFlags := log.Flags()
	log.SetFlags(0)

	var buf bytes.Buffer
	log.SetOutput(&buf)
	t.Cleanup(func() {
		log.SetOutput(originalWriter)
		log.SetFlags(originalFlags)
	})

	util.RunAndRecover(func() {
		panic("boom")
	})

	if !strings.Contains(buf.String(), "panic recovered: boom") {
		t.Fatalf("RunAndRecover() log = %q", buf.String())
	}
}

func TestRunAndRecoverNilFunc(t *testing.T) {
	util.RunAndRecover(nil)
}

func TestRetry(t *testing.T) {
	t.Run("success after retries", func(t *testing.T) {
		calls := 0
		err := util.Retry(3, func() error {
			calls++
			if calls < 3 {
				return os.ErrExist
			}
			return nil
		})
		if err != nil {
			t.Fatalf("Retry() error = %v", err)
		}
		if calls != 3 {
			t.Fatalf("Retry() calls = %d, want 3", calls)
		}
	})

	t.Run("returns final error", func(t *testing.T) {
		calls := 0
		wantErr := errors.New("failed")
		err := util.RetryWithInterval(2, time.Millisecond, func() error {
			calls++
			return wantErr
		})
		if !errors.Is(err, wantErr) {
			t.Fatalf("RetryWithInterval() error = %v, want %v", err, wantErr)
		}
		if calls != 2 {
			t.Fatalf("RetryWithInterval() calls = %d, want 2", calls)
		}
	})

	t.Run("nil retry func", func(t *testing.T) {
		err := util.Retry(1, nil)
		if !errors.Is(err, util.ErrNilRetryFunc) {
			t.Fatalf("Retry(nil) error = %v", err)
		}
	})

	t.Run("zero retry does not call", func(t *testing.T) {
		calls := 0
		err := util.Retry(0, func() error {
			calls++
			return nil
		})
		if err != nil {
			t.Fatalf("Retry(0) error = %v", err)
		}
		if calls != 0 {
			t.Fatalf("Retry(0) calls = %d, want 0", calls)
		}
	})
}

func TestFirstError(t *testing.T) {
	wantErr := errors.New("boom")
	got := util.FirstError(nil, wantErr, errors.New("later"))
	if !errors.Is(got, wantErr) {
		t.Fatalf("FirstError() = %v, want %v", got, wantErr)
	}

	if got := util.FirstError(nil, nil); got != nil {
		t.Fatalf("FirstError(all nil) = %v, want nil", got)
	}
}
