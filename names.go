package deagon

import (
	"strings"

	"github.com/oakroots/deagon/corpus"
)

const (
	// Bitmask for gender (0 = male, 1 = female)
	maskGender int = 0x0000001

	// Bitmask for given name index
	maskGivenName int = 0x00001FE

	// Bitmask for surname index
	maskSurname int = 0x1FFFE00

	// Total number of entries (2^25)
	totalEntriesFull int = 33554432
)

// findName extracts a name from the given byte slice by index and fixed length.
func findName(ix int, data []byte, length int) string {
	pos := ix * length
	if pos+length > len(data) {
		return ""
	}
	return strings.TrimSpace(string(data[pos : pos+length]))
}

// getNames returns the first name and surname for a given index.
func getNames(index int) (string, string) {
	var firstname, surname string

	// Extract given name index (shift right by 1 bit)
	givenIx := (index & maskGivenName) >> 1

	// Extract surname index (shift right by 9 bits)
	surIx := (index & maskSurname) >> 9

	// Select male or female name blob depending on the gender bit
	if (index & maskGender) == 0 {
		firstname = findName(givenIx, corpus.MaleNamesBlob, corpus.NameLength)
	} else {
		firstname = findName(givenIx, corpus.FemaleNamesBlob, corpus.NameLength)
	}

	// Always select surname from surname blob
	surname = findName(surIx, corpus.SurnamesBlob, corpus.SurnameLength)

	return firstname, surname
}

// getName returns the formatted full name (first name + surname)
// using the provided Formatter implementation.
func getName(index int, formatter Formatter) string {
	firstname, surname := getNames(index)
	return formatter.Format(firstname, surname)
}
