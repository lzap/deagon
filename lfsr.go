package deagon

func lfsr8(seed uint8) uint8 {
	s := seed
	b := (s >> 0) ^ (s >> 2) ^ (s >> 3) ^ (s >> 4)
	return (s >> 1) | (b << 7)
}

func lfsr32(seed uint32) uint32 {
	s := seed
	b := (s >> 0) ^ (s >> 2) ^ (s >> 6) ^ (s >> 7)
	return (s >> 1) | (b << 31)
}
