package util_test

import (
	"testing"

	"github.com/bang-go/util"
	"github.com/google/uuid"
)

func TestNewUUID(t *testing.T) {
	got := util.NewUUID()
	if _, err := uuid.Parse(got); err != nil {
		t.Fatalf("NewUUID() = %q, parse error = %v", got, err)
	}
}
