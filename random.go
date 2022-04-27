package deagon

import "math/rand"

func RandomName(formatter Formatter) string {
	n := rand.Int31n(mask25bits)
	return getName(n, formatter)
}
