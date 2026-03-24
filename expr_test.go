package util_test

import (
	"testing"

	"github.com/bang-go/util"
)

func TestIf(t *testing.T) {
	if got := util.If(true, 1, 2); got != 1 {
		t.Fatalf("If(true) = %d, want 1", got)
	}
	if got := util.If(false, 1, 2); got != 2 {
		t.Fatalf("If(false) = %d, want 2", got)
	}
}
