package deagon

const MASK int = 0x1ffffff

// Taps from http://users.ece.cmu.edu/~koopman/lfsr/25.txt
// 1000004 1000007 1000016 ... 1FFFFE9 1FFFFEA 1FFFFF7 1FFFFF8

func lfsr25(seed int) int {
	s := seed
	if s == 0 {
		s = 1
	}
	// 0x10002A3 = 0b0001000000000000001010100011
	b := ((s >> 0) ^ (s >> 1) ^ (s >> 2) ^ (s >> 6) ^ (s >> 8) ^ (s >> 10) ^ (s >> 25)) & MASK
	return ((s >> 1) | (b << 24)) & MASK
}

// PseudoRandomName returns a pseudorandom name for a given seed value and formats it
// via the provided formatter. Returns the next state which must be provided for the
// next call, and the generated name.
//
// The sequence of names is guaranteed to be unique until it cycles after 2^25-1 calls.
//
// When eliminateCloseNames is set, the generated name will not share a firstname or
// surname with the name associated with the input seed. This eliminates 66046
// possible names from the sequence, making the cycle shorter.
func PseudoRandomName(seed int, eliminateCloseNames bool, formatter Formatter) (int, string) {
	var next int = seed
	seedN1, seedN2 := getNames(seed)
	if eliminateCloseNames {
		for {
			next = lfsr25(next)
			n1, n2 := getNames(next)
			if n1 != seedN1 && n2 != seedN2 {
				break
			}
		}
	} else {
		next = lfsr25(seed)
	}
	return next, getName(next, formatter)
}
