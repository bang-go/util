package util

import "testing"

func TestMin(t *testing.T) {
	t.Parallel()

	if got := Min(3, 7); got != 3 {
		t.Fatalf("Min(3, 7) = %d, want 3", got)
	}
	if got := Min("b", "a"); got != "a" {
		t.Fatalf("Min(\"b\", \"a\") = %q, want %q", got, "a")
	}
}

func TestMax(t *testing.T) {
	t.Parallel()

	if got := Max(3, 7); got != 7 {
		t.Fatalf("Max(3, 7) = %d, want 7", got)
	}
	if got := Max("b", "a"); got != "b" {
		t.Fatalf("Max(\"b\", \"a\") = %q, want %q", got, "b")
	}
}
