package deagon

import "math/rand"

func RandomName(formatter Formatter) string {
	n := rand.Int31n(totalEntriesFull)
	return getName(n, formatter)
}
