package deagon

import (
	"testing"
)

func TestIndexName(t *testing.T) {
	want := "Susie Adhami"
	fmt := NewCapitalizedSpaceFormatter()
	if got := Name(fmt, 130513); got != want {
		t.Errorf("Name() = %v, want %v", got, want)
	}
}
