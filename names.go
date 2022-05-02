package deagon

import (
	"github.com/lzap/deagon/corpus"
	"strings"
)

const (
	// Masks for male, female and surname
	maskGender    int = 0x0000001
	maskGivenName int = 0x00001FE
	maskSurname   int = 0x1FFFE00
	// Total number of entries
	totalEntriesFull int = 33554432 // 2 ^ 25
)

func findName(ix int, data []byte, length int) string {
	pos := ix * length
	if int(pos+length) > len(data) {
		return ""
	}
	return strings.TrimSpace(string(data[pos : pos+length]))
}

func getNames(index int) (string, string) {
	var firstname, surname string
	givenIx := (index & maskGivenName) >> 1
	surIx := index & maskSurname >> 9
	if (index & maskGender) == 0 {
		firstname = findName(givenIx, corpus.MaleNamesBlob, corpus.NameLength)
	} else {
		firstname = findName(givenIx, corpus.FemaleNamesBlob, corpus.NameLength)
	}
	surname = findName(surIx, corpus.SurnamesBlob, corpus.SurnameLength)
	return firstname, surname
}

func getName(index int, formatter Formatter) string {
	firstname, surname := getNames(index)
	return formatter.Format(firstname, surname)
}
