package deagon

const (
	// MASK is a 25-bit mask (2^25 - 1). Kept exported to avoid breaking callers.
	MASK = (1 << 25) - 1

	// lfsrWidth is the bit-width of the LFSR.
	lfsrWidth = 25

	// msbPos is the position of the most significant bit within the 25-bit register.
	msbPos = lfsrWidth - 1
)

// lfsr25 advances a 25-bit Fibonacci LFSR by one step.
//
// Notes:
//   - The state is kept within 25 bits (MASK).
//   - A zero seed is mapped to 1 to avoid the all-zero lock-up state.
//   - Feedback taps are taken according to the 0x10002A3 mask (see comment below).
//
// The feedback bit is the parity of selected tap bits. It must be a single bit (0/1).
func lfsr25(seed int) int {
	// Normalize seed into 25 bits and avoid zero state.
	s := seed & MASK
	if s == 0 {
		s = 1
	}

	// Taps from polynomial mask 0x10002A3:
	// Bits used (0-indexed): 0, 1, 2, 6, 8, 10, 24 (MSB)
	// Compute parity (XOR) of the tap bits and keep only LSB as feedback bit.
	x := (s >> 0) ^ (s >> 1) ^ (s >> 2) ^ (s >> 6) ^ (s >> 8) ^ (s >> 10) ^ (s >> 24)
	newBit := x & 1

	// Shift right and insert feedback bit at MSB position.
	next := ((s >> 1) | (newBit << msbPos)) & MASK
	return next
}

// PseudoRandomName returns a pseudorandom name for a given seed value and formats it
// via the provided formatter. It returns the next LFSR state (to be used for the next call)
// and the generated, formatted name.
//
// The sequence is unique until it cycles after 2^25 - 2 states, excluding the all-zero state.
//
// When eliminateCloseNames is true, successive calls never return the same firstname or surname.
// This skips 66,046 possible states and the loop becomes slightly shorter.
func PseudoRandomName(seed int, eliminateCloseNames bool, formatter Formatter) (int, string) {
	// Normalize incoming seed similarly to lfsr25; if 0, bump to 1.
	next := seed & MASK
	if next == 0 {
		next = 1
	}

	seedN1, seedN2 := getNames(next)

	if eliminateCloseNames {
		// Safety cap prevents an accidental infinite loop if corpus or taps ever change.
		const safetyCap = totalEntriesFull // conservative upper bound
		for i := 0; i < safetyCap; i++ {
			next = lfsr25(next)
			n1, n2 := getNames(next)
			if n1 != seedN1 && n2 != seedN2 {
				break
			}
		}
	} else {
		next = lfsr25(next)
	}

	return next, getName(next, formatter)
}
