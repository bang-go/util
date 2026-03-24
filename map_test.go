package util_test

import (
	"testing"

	"github.com/bang-go/util"
)

func TestMapHelpers(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	key, ok := util.MapKeyForValue(m, 2)
	if !ok || key != "b" {
		t.Fatalf("MapKeyForValue() = (%q, %v)", key, ok)
	}

	if util.MapContainsValue(m, 3) {
		t.Fatal("MapContainsValue() should be false")
	}
	if !util.MapContainsValue(m, 1) {
		t.Fatal("MapContainsValue() should be true")
	}
}
