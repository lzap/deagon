package deagon

import (
	"testing"

	"github.com/oakroots/deagon/corpus"
)

func TestFirstFemaleNameFind(t *testing.T) {
	want := "ADA"
	got := findName(0, corpus.FemaleNamesBlob, corpus.NameLength)
	if want != got {
		t.Fatalf(`Expected to get %s but got %#q`, want, got)
	}
}

func TestLastFemaleNameFind(t *testing.T) {
	want := "WILMA"
	got := findName(255, corpus.FemaleNamesBlob, corpus.NameLength)
	if want != got {
		t.Fatalf(`Expected to get %s but got %#q`, want, got)
	}
}

func TestFirstMaleNameFind(t *testing.T) {
	want := "AARON"
	got := findName(0, corpus.MaleNamesBlob, corpus.NameLength)
	if want != got {
		t.Fatalf(`Expected to get %s but got %#q`, want, got)
	}
}

func TestLastMaleNameFind(t *testing.T) {
	want := "WILL"
	got := findName(255, corpus.MaleNamesBlob, corpus.NameLength)
	if want != got {
		t.Fatalf(`Expected to get %s but got %#q`, want, got)
	}
}

func TestFirstSurnameFind(t *testing.T) {
	want := "AABERG"
	got := findName(0, corpus.SurnamesBlob, corpus.SurnameLength)
	if want != got {
		t.Fatalf(`Expected to get %s but got %#q`, want, got)
	}
}

func TestSecondSurnameFind(t *testing.T) {
	want := "AADLAND"
	got := findName(1, corpus.SurnamesBlob, corpus.SurnameLength)
	if want != got {
		t.Fatalf(`Expected to get %s but got %#q`, want, got)
	}
}

func Test13thSurnameFind(t *testing.T) {
	want := "AARONSON"
	got := findName(13, corpus.SurnamesBlob, corpus.SurnameLength)
	if want != got {
		t.Fatalf(`Expected to get %s but got %#q`, want, got)
	}
}

func TestLastSurnameFind(t *testing.T) {
	want := "ZYWIEC"
	got := findName(65535, corpus.SurnamesBlob, corpus.SurnameLength)
	if want != got {
		t.Fatalf(`Expected to get %s but got %#q`, want, got)
	}
}

func TestOutOfBoundsSurnameFind(t *testing.T) {
	want := ""
	got := findName(65536, corpus.SurnamesBlob, corpus.SurnameLength)
	if want != got {
		t.Fatalf(`Expected to get %s but got %#q`, want, got)
	}
}
