package d_test

import (
	"testing"

	"github.com/jhonghee/d"
)

func TestVersion(t *testing.T) {
	expected := "D v1.2->E v1.1"
	if d.Version() != expected {
		t.Error("Expected", expected, "but got", d.Version())
	}
}
