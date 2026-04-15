package util_test

import (
	"testing"

	"github.com/bang-go/util"
)

func TestClonePtr(t *testing.T) {
	t.Parallel()

	t.Run("nil", func(t *testing.T) {
		if util.ClonePtr[int64](nil) != nil {
			t.Fatal("ClonePtr(nil) should return nil")
		}
	})

	t.Run("copy value", func(t *testing.T) {
		src := util.Ptr(int64(42))
		cloned := util.ClonePtr(src)
		if cloned == nil {
			t.Fatal("ClonePtr should return a non-nil pointer")
		}
		if cloned == src {
			t.Fatal("ClonePtr should create a new pointer")
		}
		if *cloned != *src {
			t.Fatalf("ClonePtr value mismatch: got %d want %d", *cloned, *src)
		}

		*cloned = 7
		if *src != 42 {
			t.Fatalf("ClonePtr should not mutate source: got %d want %d", *src, 42)
		}
	})
}

func TestPtrIfNonZero(t *testing.T) {
	t.Parallel()

	t.Run("zero returns nil", func(t *testing.T) {
		if util.PtrIfNonZero(uint64(0)) != nil {
			t.Fatal("PtrIfNonZero(zero) should return nil")
		}
	})

	t.Run("non zero returns pointer", func(t *testing.T) {
		got := util.PtrIfNonZero(uint64(42))
		if got == nil {
			t.Fatal("PtrIfNonZero(non-zero) should return a non-nil pointer")
		}
		if *got != 42 {
			t.Fatalf("PtrIfNonZero(non-zero) = %d, want %d", *got, 42)
		}
	})
}

func TestNilIfZero(t *testing.T) {
	t.Parallel()

	t.Run("nil", func(t *testing.T) {
		if util.NilIfZero[uint64](nil) != nil {
			t.Fatal("NilIfZero(nil) should return nil")
		}
	})

	t.Run("zero returns nil", func(t *testing.T) {
		value := util.Ptr(uint64(0))
		if util.NilIfZero(value) != nil {
			t.Fatal("NilIfZero(zero) should return nil")
		}
	})

	t.Run("non zero returns cloned pointer", func(t *testing.T) {
		value := util.Ptr(uint64(42))
		got := util.NilIfZero(value)
		if got == nil {
			t.Fatal("NilIfZero(non-zero) should return a non-nil pointer")
		}
		if got == value {
			t.Fatal("NilIfZero(non-zero) should return a cloned pointer")
		}
		if *got != 42 {
			t.Fatalf("NilIfZero(non-zero) = %d, want %d", *got, 42)
		}
	})
}

func TestDerefZero(t *testing.T) {
	t.Parallel()

	t.Run("nil returns zero", func(t *testing.T) {
		if got := util.DerefZero[int64](nil); got != 0 {
			t.Fatalf("DerefZero(nil) = %d, want 0", got)
		}
	})

	t.Run("non nil returns value", func(t *testing.T) {
		value := util.Ptr(int64(42))
		if got := util.DerefZero(value); got != 42 {
			t.Fatalf("DerefZero(value) = %d, want 42", got)
		}
	})
}

func TestDerefOr(t *testing.T) {
	t.Parallel()

	t.Run("nil returns default", func(t *testing.T) {
		if got := util.DerefOr[string](nil, "fallback"); got != "fallback" {
			t.Fatalf("DerefOr(nil) = %q, want %q", got, "fallback")
		}
	})

	t.Run("non nil returns value", func(t *testing.T) {
		value := util.Ptr("value")
		if got := util.DerefOr(value, "fallback"); got != "value" {
			t.Fatalf("DerefOr(value) = %q, want %q", got, "value")
		}
	})
}

func TestMustDeref(t *testing.T) {
	t.Parallel()

	value := util.Ptr(int64(42))
	if got := util.MustDeref(value); got != 42 {
		t.Fatalf("MustDeref(value) = %d, want 42", got)
	}
}
