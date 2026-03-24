package errcode

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

func TestError_Error(t *testing.T) {
	err := New(404, "Not Found")
	expected := "code=404 message=Not Found"
	if err.Error() != expected {
		t.Errorf("Expected %q, got %q", expected, err.Error())
	}
}

func TestHelpers(t *testing.T) {
	err := New(500, "Internal Server Error")

	code, ok := Code(err)
	if !ok || code != 500 {
		t.Errorf("Expected code 500, got (%d, %v)", code, ok)
	}
	message, ok := Message(err)
	if !ok || message != "Internal Server Error" {
		t.Errorf("Expected message 'Internal Server Error', got (%q, %v)", message, ok)
	}

	// Test with standard error
	stdErr := errors.New("standard error")
	if _, ok := Code(stdErr); ok {
		t.Error("Expected Code(stdErr) to return ok=false")
	}
	if _, ok := Message(stdErr); ok {
		t.Error("Expected Message(stdErr) to return ok=false")
	}

	// Test with nil
	if _, ok := Code(nil); ok {
		t.Error("Expected Code(nil) to return ok=false")
	}
	if _, ok := Message(nil); ok {
		t.Error("Expected Message(nil) to return ok=false")
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

func TestHasCode(t *testing.T) {
	err := New(404, "foo")
	wrapped := fmt.Errorf("wrap: %w", err)

	if !HasCode(wrapped, 404) {
		t.Error("HasCode failed on wrapped error")
	}
	if HasCode(wrapped, 500) {
		t.Error("HasCode matched wrong code")
	}
	if HasCode(nil, 404) {
		t.Error("HasCode(nil) should be false")
	}
}

func TestTypedNil(t *testing.T) {
	var e *Error // nil pointer of type *Error
	var err error = e

	if _, ok := Code(err); ok {
		t.Error("Expected Code(typed nil) to return ok=false")
	}
	if _, ok := Message(err); ok {
		t.Error("Expected Message(typed nil) to return ok=false")
	}
	if err.Error() != "<nil>" {
		t.Errorf("Expected <nil> for typed nil Error(), got %q", err.Error())
	}

	// Test Is method safety
	if e.Is(errors.New("foo")) {
		t.Error("nil *Error.Is should return false")
	}
	if HasCode(err, 404) {
		t.Error("HasCode(typed nil) should return false")
	}

	var target *Error
	var targetErr error = target
	if e := New(404, "foo"); errors.Is(e, targetErr) {
		t.Error("errors.Is with typed nil target should be false")
	}
}

func TestAs(t *testing.T) {
	err := fmt.Errorf("wrap: %w", New(401, "unauthorized"))
	got, ok := As(err)
	if !ok {
		t.Fatal("As() should find errcode.Error in wrapped chain")
	}
	if got.Code != 401 || got.Message != "unauthorized" {
		t.Fatalf("As() = %+v", got)
	}
}

func TestJSON(t *testing.T) {
	data, err := json.Marshal(New(403, "forbidden"))
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}
	if string(data) != `{"code":403,"message":"forbidden"}` {
		t.Fatalf("Marshal() = %s", data)
	}
}
