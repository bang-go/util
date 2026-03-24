package util_test

import (
	"testing"

	"github.com/bang-go/util"
)

func TestIntRandRange(t *testing.T) {
	got, err := util.IntRandRange(5, 5)
	if err != nil {
		t.Fatalf("IntRandRange(equal) error = %v", err)
	}
	if got != 5 {
		t.Fatalf("IntRandRange(equal) = %d", got)
	}

	seen := map[int]bool{}
	for i := 0; i < 200; i++ {
		got, err := util.IntRandRange(1, 2)
		if err != nil {
			t.Fatalf("IntRandRange() error = %v", err)
		}
		if got < 1 || got > 2 {
			t.Fatalf("IntRandRange() = %d, want [1,2]", got)
		}
		seen[got] = true
	}
	if !seen[1] || !seen[2] {
		t.Fatalf("IntRandRange() should hit both bounds, seen = %#v", seen)
	}

	got, err = util.IntRandRange(3, -3)
	if err != nil {
		t.Fatalf("IntRandRange(swapped) error = %v", err)
	}
	if got < -3 || got > 3 {
		t.Fatalf("IntRandRange(swapped) = %d, want [-3,3]", got)
	}
}

func TestIntRandRangeUint64(t *testing.T) {
	got, err := util.IntRandRange[uint64](1<<63, 1<<63)
	if err != nil {
		t.Fatalf("IntRandRange(uint64) error = %v", err)
	}
	if got != 1<<63 {
		t.Fatalf("IntRandRange(uint64) = %d, want %d", got, uint64(1<<63))
	}
}
