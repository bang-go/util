package util_test

import (
	"reflect"
	"testing"

	"github.com/bang-go/util"
)

func TestSliceMap(t *testing.T) {
	t.Parallel()

	got := util.SliceMap([]int{1, 2, 3}, func(value int) string {
		return string(rune('0' + value))
	})
	expect := []string{"1", "2", "3"}
	if !reflect.DeepEqual(got, expect) {
		t.Fatalf("SliceMap() = %v, want %v", got, expect)
	}
}

func TestSliceToAny(t *testing.T) {
	t.Parallel()

	got := util.SliceToAny([]uint64{1, 2, 3})
	expect := []any{uint64(1), uint64(2), uint64(3)}
	if !reflect.DeepEqual(got, expect) {
		t.Fatalf("SliceToAny() = %v, want %v", got, expect)
	}
}

func TestSliceCompact(t *testing.T) {
	t.Parallel()

	got := util.SliceCompact([]uint64{0, 1, 0, 2, 3, 0})
	expect := []uint64{1, 2, 3}
	if !reflect.DeepEqual(got, expect) {
		t.Fatalf("SliceCompact() = %v, want %v", got, expect)
	}
}

func TestSliceCompactUnique(t *testing.T) {
	t.Parallel()

	got := util.SliceCompactUnique([]uint64{0, 3, 1, 3, 0, 2, 1, 0})
	expect := []uint64{3, 1, 2}
	if !reflect.DeepEqual(got, expect) {
		t.Fatalf("SliceCompactUnique() = %v, want %v", got, expect)
	}
}

func TestSliceUnique(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []uint64
		expect []uint64
	}{
		{
			name:   "nil",
			input:  nil,
			expect: nil,
		},
		{
			name:   "empty",
			input:  []uint64{},
			expect: nil,
		},
		{
			name:   "preserve order",
			input:  []uint64{3, 1, 3, 2, 1, 4},
			expect: []uint64{3, 1, 2, 4},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := util.SliceUnique(tc.input)
			if !reflect.DeepEqual(got, tc.expect) {
				t.Fatalf("SliceUnique() = %v, want %v", got, tc.expect)
			}
		})
	}
}

func TestSliceUniqueBy(t *testing.T) {
	t.Parallel()

	type item struct {
		ID    uint64
		Title string
	}

	input := []item{
		{ID: 1, Title: "a"},
		{ID: 2, Title: "b"},
		{ID: 1, Title: "c"},
	}
	expect := []item{
		{ID: 1, Title: "a"},
		{ID: 2, Title: "b"},
	}

	got := util.SliceUniqueBy(input, func(value item) uint64 {
		return value.ID
	})
	if !reflect.DeepEqual(got, expect) {
		t.Fatalf("SliceUniqueBy() = %v, want %v", got, expect)
	}
}
