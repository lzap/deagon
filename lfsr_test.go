package deagon

import (
	"bytes"
	"testing"
)

func TestLfsr25Fixtures(t *testing.T) {
	fixtures := []struct {
		in  int
		out int
	}{
		{1, 16777216},
		{2, 16777217},
		{3, 1},
		{4, 16777218},
		{5, 2},
		{6, 3},
		{7, 16777219},
		{8, 4},
		{9, 16777220},
		{10, 16777221},
		{42, 16777237},
		{100, 50},
	}

	for _, f := range fixtures {
		if got := lfsr25(f.in); got != f.out {
			t.Errorf("lfsr25(%d) = %d, want %d", f.in, got, f.out)
		}
	}
}

func testWithInitialValue(t *testing.T, start int) {
	value := start
	for i := 0; i < int(totalEntriesFull)-1; i++ {
		value = lfsr25(value)
	}
	if value != start {
		t.Error("period is not 2^25!")
	}
}

func TestPseudoRandom25Starting1(t *testing.T) {
	testWithInitialValue(t, 1)
}

func TestPseudoRandom25Starting42(t *testing.T) {
	testWithInitialValue(t, 42)
}

func TestCloseNames(t *testing.T) {
	value := 1
	seen := make(map[string]struct{})
	counter := 0
	var last1, last2 string

	for range int(totalEntriesFull) - 1 {
		value = lfsr25(value)
		n1, n2 := getNames(value)
		if last1 == n1 {
			counter++
			seen[n1] = struct{}{}
		}
		if last2 == n2 {
			counter++
			seen[n2] = struct{}{}
		}
		last1 = n1
		last2 = n2
	}

	if counter > 66046 {
		t.Errorf("expected %d number of close names, got %d", 66046, counter)
	}

	if len(seen) != 4 {
		t.Errorf("expected see exactly AARON, WILMA, AABERG and ZYWIEC in the repetitions but got something else: %v", seen)
	}
}

func TestLFSR25Cycle(t *testing.T) {
	seed := 1

	for range int(totalEntriesFull) - 1 {
		seed = lfsr25(seed)
	}

	if seed != 1 {
		t.Errorf("LFSR25 did not cycle back to initial seed, got %d", seed)
	}
}

func TestPseudoRandomName(t *testing.T) {
	var b bytes.Buffer
	var name string
	expected := `Aaron Lebario
Aaron Essner
Aaron Carda
Aaron Bertels
Aaron Aversa
Aaron Amores
Aaron Albany
Aaron Adjei
Aaron Abundiz
Aaron Abele
Aaron Abaya
Aaron Aavang
Aaron Aanenson
Aaron Aaland
Aaron Aagaard
Aaron Lebaron
Jimmy Essner
Doug Myhre
Caleb Gibbard
Bert Orick
`
	seed := 1
	for range 20 {
		seed, name = PseudoRandomName(seed, false, NewCapitalizedSpaceFormatter())
		b.WriteString(name)
		b.WriteString("\n")
	}

	if b.String() != expected {
		t.Fatalf("result not expected: %s", b.String())
	}
}

func TestPseudoRandomNameEliminations(t *testing.T) {
	var b bytes.Buffer
	var name string
	expected := `Aaron Lebario
Jimmy Essner
Doug Myhre
Caleb Gibbard
Bert Orick
Alvin Goslin
Alex Coullard
Adam Bonnema
Abel Magliolo
Ada Sabol
Aaron Tapscott
Jimmy Femia
Doug Catalani
Lewis Bickell
Eric Lunford
Luke Fertig
Ethan Cavasos
Clay Bielby
Bobby Lupkes
Leah Fessler
`
	seed := 1
	for range 20 {
		seed, name = PseudoRandomName(seed, true, NewCapitalizedSpaceFormatter())
		b.WriteString(name)
		b.WriteString("\n")
	}

	if b.String() != expected {
		t.Fatalf("result not expected: %s", b.String())
	}
}
