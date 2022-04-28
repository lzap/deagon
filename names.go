package deagon

import (
	"github.com/lzap/deagon/corpus"
	"strings"
)

const (
	// Full 25 bit mask (33_554_432)
	maskFull int32 = 0x1FFFFFF
	// Masks for male, female and surname
	maskGender    int32 = 0x0000001
	maskGivenName int32 = 0x00001FE
	maskSurname   int32 = 0x1FFFE00
	// Total number of entries
	totalEntriesFull      int32 = 33554432 // 2 ^ 25
	totalEntriesGender    int32 = 1
	totalEntriesGivenName int32 = 256   // 2 ^ 8
	totalEntriesSurname   int32 = 65536 // 2 ^ 16
)

func findName(ix int32, data []byte, length int32) string {
	pos := ix * length
	if int(pos+length) > len(data) {
		return ""
	}
	return strings.TrimSpace(string(data[pos : pos+length]))
}

func getName(index int32, formatter Formatter) string {
	var firstname, surname string
	givenIx := (index & maskGivenName) >> 1
	surIx := index & maskSurname >> 9
	if (index & maskGender) == 0 {
		firstname = findName(givenIx, corpus.MaleNamesBlob, corpus.NameLength)
	} else {
		firstname = findName(givenIx, corpus.FemaleNamesBlob, corpus.NameLength)
	}
	surname = findName(surIx, corpus.SurnamesBlob, corpus.SurnameLength)
	return formatter.Format(firstname, surname)
}
