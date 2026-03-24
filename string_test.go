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
