package deagon

const MASK int = 0x1ffffff

// Taps from http://users.ece.cmu.edu/~koopman/lfsr/25.txt
// 1000004 1000007 1000016 1000040 1000049 1000062 100007F ...
func lfsr25(seed int) int {
	s := seed
	if s == 0 {
		s = 1
	}
	// 1000004 (hex) = 0001000000000000000000000100 (bin)
	b := ((s >> 0) ^ (s >> 3) ^ (s >> 25)) & MASK
	return (s >> 1) | (b<<24)&MASK
}
