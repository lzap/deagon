package deagon

import (
	"github.com/lzap/deagon/corpus"
	"strings"
)

const (
	mask25bits    int32 = 0x1FFFFFF
	maskGender    int32 = 0x0000001
	maskGivenName int32 = 0x00001FE
	maskSurname   int32 = 0x1FFFE00
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
