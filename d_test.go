package d_test

import (
	"testing"

	"github.com/jhonghee/d"
)

func TestVersion(t *testing.T) {
	expected := "D v1.4->E v1.2"
	if d.Version() != expected {
		t.Error("Expected", expected, "but got", d.Version())
	}
}
