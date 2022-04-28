package deagon

import (
	"testing"
)

func testWithInitialValue(t *testing.T, start int) {
	value := start
	for i := 0; i < int(totalEntriesFull)-1; i++ {
		value = lfsr25(value)
	}
	if value != start {
		t.Error("period is not 2^25!")
	}
}

func TestPseudoRandom25Starting1(t *testing.T) {
	testWithInitialValue(t, 1)
}

func TestPseudoRandom25Starting42(t *testing.T) {
	testWithInitialValue(t, 42)
}
