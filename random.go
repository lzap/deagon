package deagon

import "math/rand"

// RandomName returns a random name and formats it with the provided formatter.
func RandomName(formatter Formatter) string {
	n := rand.Int31n(int32(totalEntriesFull))
	return getName(int(n), formatter)
}
