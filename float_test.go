package util_test

import (
	"errors"
	"testing"

	"github.com/bang-go/util"
)

func TestFloatHelpers(t *testing.T) {
	if got := util.FloatToString(1.2345, 2); got != "1.23" {
		t.Fatalf("FloatToString() = %q", got)
	}
	if got := util.FloatAdd(0.1, 0.2); got != 0.3 {
		t.Fatalf("FloatAdd() = %v", got)
	}
	if got := util.FloatSub(1.5, 0.25, 0.25); got != 1.0 {
		t.Fatalf("FloatSub() = %v", got)
	}
	if got := util.FloatMul(1.2, 3.0); got != 3.6 {
		t.Fatalf("FloatMul() = %v", got)
	}
	if got := util.FloatCeil(1.231, 2); got != 1.24 {
		t.Fatalf("FloatCeil() = %v", got)
	}
	if got := util.FloatFloor(1.239, 2); got != 1.23 {
		t.Fatalf("FloatFloor() = %v", got)
	}
	if got := util.FloatTruncate(1.239, 2); got != 1.23 {
		t.Fatalf("FloatTruncate() = %v", got)
	}
	if got := util.FloatCompare(1.23, 1.23); got != 0 {
		t.Fatalf("FloatCompare(equal) = %d", got)
	}
	if got := util.FloatCompare(1.24, 1.23); got != 1 {
		t.Fatalf("FloatCompare(greater) = %d", got)
	}
	if got := util.FloatCompare(1.22, 1.23); got != -1 {
		t.Fatalf("FloatCompare(less) = %d", got)
	}
	if got := util.FloatAbs(-1.5); got != 1.5 {
		t.Fatalf("FloatAbs() = %v", got)
	}

	var sum32 float32 = util.FloatAdd(float32(0.1), float32(0.2))
	if sum32 != 0.3 {
		t.Fatalf("FloatAdd(float32) = %v", sum32)
	}
}

func TestFloatDiv(t *testing.T) {
	got, err := util.FloatDiv(8.0, 2.0, 2.0)
	if err != nil {
		t.Fatalf("FloatDiv() error = %v", err)
	}
	if got != 2.0 {
		t.Fatalf("FloatDiv() = %v", got)
	}

	if _, err := util.FloatDiv(1.0, 0.0); !errors.Is(err, util.ErrDivisionByZero) {
		t.Fatalf("FloatDiv(div by zero) error = %v", err)
	}

	defer func() {
		recovered := recover()
		err, ok := recovered.(error)
		if !ok || !errors.Is(err, util.ErrDivisionByZero) {
			t.Fatalf("MustFloatDiv() panic = %v", recovered)
		}
	}()
	util.MustFloatDiv(1.0, 0.0)
}
