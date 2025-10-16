package deagon

import "math/rand/v2"

// RandomName returns a random formatted name using the provided Formatter.
// It selects a random index from the full set of available entries.
func RandomName(formatter Formatter) string {
	n := rand.IntN(totalEntriesFull)
	return getName(n, formatter)
}
