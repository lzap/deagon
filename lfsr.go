package deagon

import "math/bits"

const MASK int = 0x1ffffff

// lfsr25Taps is a bitmask for the tap positions 0, 1, 2, 6, 8, 10, and 25 of
// 0x10002A3 which is in binary 0b0001000000000000001010100011.
//
// A tap from http://users.ece.cmu.edu/~koopman/lfsr/25.txt
const lfsr25Taps = (1 << 0) | (1 << 1) | (1 << 2) | (1 << 6) | (1 << 8) | (1 << 10) | (1 << 25)

func lfsr25(seed int) int {
	s := seed
	if s == 0 {
		s = 1
	}
	// Calculate the parity of the tapped bits.
	// If the number of set bits is odd, the parity is 1, otherwise 0.
	b := bits.OnesCount(uint(s&lfsr25Taps)) & 1
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
	next := lfsr25(seed)
	if eliminateCloseNames {
		seedN1, seedN2 := getNames(seed)
		n1, n2 := getNames(next)
		for n1 == seedN1 || n2 == seedN2 {
			next = lfsr25(next)
			n1, n2 = getNames(next)
		}
	}
	return next, getName(next, formatter)
}
