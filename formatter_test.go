package deagon

import (
	"testing"
)

func TestUppercaseSpaceFormatter(t *testing.T) {
	want := []string{
		"AARON AABERG",
		"ADA AABERG",
		"ABEL AABERG",
		"ADDIE AABERG",
		"ADAM AABERG",
		"ADELE AABERG",
		"ALAN AABERG",
		"AGNES AABERG",
	}
	for i := 0; i < 8; i++ {
		got := getName(int32(i), NewUppercaseSpaceFormatter())
		if want[i] != got {
			t.Fatalf(`Expected to get %s but got %#q`, want[i], got)
		}
	}
}

func TestCapitalizedSpaceFormatter(t *testing.T) {
	want := []string{
		"Aaron Aaberg",
		"Ada Aaberg",
		"Abel Aaberg",
		"Addie Aaberg",
		"Adam Aaberg",
		"Adele Aaberg",
		"Alan Aaberg",
		"Agnes Aaberg",
	}
	for i := 0; i < 8; i++ {
		got := getName(int32(i), NewCapitalizedSpaceFormatter())
		if want[i] != got {
			t.Fatalf(`Expected to get %s but got %#q`, want[i], got)
		}
	}
}
