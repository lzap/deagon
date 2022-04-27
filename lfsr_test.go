package deagon

import (
	"testing"
)

func TestLfsr32(t *testing.T) {
	var seed uint8 = 0
	var v uint8 = lfsr8(seed)
	c := 1

	for v != seed {
		v = lfsr8(v)
		c++
		println(c, v)
		if c > 0xffffffff {
			t.Error("lfsr32 did not complete the full period and restarted at ", c)
		}
	}

}
