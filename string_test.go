package util_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/bang-go/util"
)

func TestIsBlank(t *testing.T) {
	if !util.IsBlank(" \n\t ") {
		t.Fatal("IsBlank() should be true")
	}
	if util.IsBlank(" x ") {
		t.Fatal("IsBlank() should be false")
	}
}

func TestNilIfBlank(t *testing.T) {
	t.Parallel()

	t.Run("nil returns nil", func(t *testing.T) {
		if got := util.NilIfBlank(nil); got != nil {
			t.Fatalf("NilIfBlank(nil) = %v, want nil", *got)
		}
	})

	t.Run("blank returns nil", func(t *testing.T) {
		value := util.Ptr(" \n\t ")
		if got := util.NilIfBlank(value); got != nil {
			t.Fatalf("NilIfBlank(blank) = %v, want nil", *got)
		}
	})

	t.Run("trimmed value returns cloned pointer", func(t *testing.T) {
		value := util.Ptr("  bang  ")
		got := util.NilIfBlank(value)
		if got == nil {
			t.Fatal("NilIfBlank(non-blank) should return a non-nil pointer")
		}
		if got == value {
			t.Fatal("NilIfBlank should return a new pointer")
		}
		if *got != "bang" {
			t.Fatalf("NilIfBlank(non-blank) = %q, want %q", *got, "bang")
		}
	})
}

func TestPtrIfNonBlank(t *testing.T) {
	t.Parallel()

	t.Run("blank returns nil", func(t *testing.T) {
		if got := util.PtrIfNonBlank(" \n\t "); got != nil {
			t.Fatalf("PtrIfNonBlank(blank) = %v, want nil", *got)
		}
	})

	t.Run("trimmed value returns pointer", func(t *testing.T) {
		got := util.PtrIfNonBlank("  bang  ")
		if got == nil {
			t.Fatal("PtrIfNonBlank(non-blank) should return a non-nil pointer")
		}
		if *got != "bang" {
			t.Fatalf("PtrIfNonBlank(non-blank) = %q, want %q", *got, "bang")
		}
	})
}

func TestDerefTrimmed(t *testing.T) {
	t.Parallel()

	t.Run("nil returns empty string", func(t *testing.T) {
		if got := util.DerefTrimmed(nil); got != "" {
			t.Fatalf("DerefTrimmed(nil) = %q, want empty string", got)
		}
	})

	t.Run("blank returns empty string", func(t *testing.T) {
		value := util.Ptr(" \n\t ")
		if got := util.DerefTrimmed(value); got != "" {
			t.Fatalf("DerefTrimmed(blank) = %q, want empty string", got)
		}
	})

	t.Run("trimmed value returns string", func(t *testing.T) {
		value := util.Ptr("  bang  ")
		if got := util.DerefTrimmed(value); got != "bang" {
			t.Fatalf("DerefTrimmed(non-blank) = %q, want %q", got, "bang")
		}
	})
}

func TestStringRand(t *testing.T) {
	tests := []struct {
		name    string
		fn      func(int) (string, error)
		allowed string
	}{
		{name: "all", fn: util.StringRand, allowed: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"},
		{name: "number", fn: util.StringRandNumber, allowed: "0123456789"},
		{name: "letter", fn: util.StringRandLetter, allowed: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"},
		{name: "lower", fn: util.StringRandLowerLetter, allowed: "abcdefghijklmnopqrstuvwxyz"},
		{name: "upper", fn: util.StringRandUpperLetter, allowed: "ABCDEFGHIJKLMNOPQRSTUVWXYZ"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fn(64)
			if err != nil {
				t.Fatalf("rand error = %v", err)
			}
			if len(got) != 64 {
				t.Fatalf("len = %d, want 64", len(got))
			}
			if !containsOnly(got, tt.allowed) {
				t.Fatalf("generated string contains invalid chars: %q", got)
			}
		})
	}

	got, err := util.StringRand(0)
	if err != nil || got != "" {
		t.Fatalf("StringRand(0) = (%q, %v)", got, err)
	}

	if _, err := util.StringRand(-1); !errors.Is(err, util.ErrNegativeLength) {
		t.Fatalf("StringRand(-1) error = %v", err)
	}
}

func containsOnly(s, allowed string) bool {
	for _, r := range s {
		if !strings.ContainsRune(allowed, r) {
			return false
		}
	}
	return true
}
