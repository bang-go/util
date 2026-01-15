package errcode

import (
	"errors"
	"fmt"
	"testing"
)

func TestError_Error(t *testing.T) {
	err := New(404, "Not Found")
	expected := "code: 404, msg: Not Found"
	if err.Error() != expected {
		t.Errorf("Expected %q, got %q", expected, err.Error())
	}
}

func TestHelpers(t *testing.T) {
	err := New(500, "Internal Server Error")

	if Code(err) != 500 {
		t.Errorf("Expected code 500, got %d", Code(err))
	}
	if Msg(err) != "Internal Server Error" {
		t.Errorf("Expected msg 'Internal Server Error', got %q", Msg(err))
	}

	// Test with standard error
	stdErr := errors.New("standard error")
	if Code(stdErr) != -1 {
		t.Errorf("Expected code -1 for std error, got %d", Code(stdErr))
	}
	if Msg(stdErr) != "standard error" {
		t.Errorf("Expected msg 'standard error', got %q", Msg(stdErr))
	}

	// Test with nil
	if Code(nil) != 0 {
		t.Errorf("Expected code 0 for nil, got %d", Code(nil))
	}
	if Msg(nil) != "" {
		t.Errorf("Expected empty msg for nil, got %q", Msg(nil))
	}
}

func TestIs(t *testing.T) {
	err1 := New(404, "foo")
	err2 := New(404, "bar")
	err3 := New(500, "foo")

	if !errors.Is(err1, err2) {
		t.Error("Same code should be Is=true")
	}
	if errors.Is(err1, err3) {
		t.Error("Different code should be Is=false")
	}
}

func TestIsCode(t *testing.T) {
	err := New(404, "foo")
	wrapped := fmt.Errorf("wrap: %w", err)

	if !IsCode(wrapped, 404) {
		t.Error("IsCode failed on wrapped error")
	}
	if IsCode(wrapped, 500) {
		t.Error("IsCode matched wrong code")
	}
	if IsCode(nil, 404) {
		t.Error("IsCode(nil) should be false")
	}
}

func TestTypedNil(t *testing.T) {
	var e *Error // nil pointer of type *Error
	var err error = e

	if Code(err) != -1 {
		t.Errorf("Expected code -1 for typed nil, got %d", Code(err))
	}
	if Msg(err) != "" {
		t.Errorf("Expected empty msg for typed nil, got %q", Msg(err))
	}

	// Test Is method safety
	if e.Is(errors.New("foo")) {
		t.Error("nil *Error.Is should return false")
	}
}
