package util_test

import (
	"testing"
	"time"

	"github.com/bang-go/util"
)

func TestLoadLocationOrFixed(t *testing.T) {
	loc := util.LoadLocationOrFixed("UTC", 8*3600)
	_, offset := time.Unix(0, 0).In(loc).Zone()
	if offset != 0 {
		t.Fatalf("LoadLocationOrFixed(UTC) offset = %d, want 0", offset)
	}

	fallback := util.LoadLocationOrFixed("Invalid/Timezone", 8*3600)
	name, offset := time.Unix(0, 0).In(fallback).Zone()
	if name != "Invalid/Timezone" || offset != 8*3600 {
		t.Fatalf("LoadLocationOrFixed(fallback) = (%q, %d)", name, offset)
	}
}
