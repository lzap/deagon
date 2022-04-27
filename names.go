package deagon

import (
	"strings"
)

func findName(ix int, data []byte, length int) string {
	pos := ix * length
	if pos+length > len(data) {
		return ""
	}
	return strings.TrimSpace(string(data[pos : pos+length]))
}
