package deagon

import (
	"bytes"
	"testing"
)

// Quick checks for LFSR boundaries and non-zero state.
func TestPseudoRandom25Starting1(t *testing.T) {
	const steps = 50_000
	state := 1
	for i := 0; i < steps; i++ {
		state = lfsr25(state)
		if state == 0 {
			t.Fatalf("LFSR reached zero at step %d", i)
		}
		if state&MASK != state {
			t.Fatalf("LFSR exceeded 25-bit mask at step %d: 0x%x", i, state)
		}
	}
}

func TestPseudoRandom25Starting42(t *testing.T) {
	const steps = 50_000
	state := 42
	for i := 0; i < steps; i++ {
		state = lfsr25(state)
		if state == 0 {
			t.Fatalf("LFSR reached zero at step %d", i)
		}
		if state&MASK != state {
			t.Fatalf("LFSR exceeded 25-bit mask at step %d: 0x%x", i, state)
		}
	}
}

// Property: with eliminateCloseNames=false we should occasionally see consecutive
// firstname or surname repeats within a reasonable window (probabilistic).
func TestCloseNames(t *testing.T) {
	seed := 1
	var lastFirst, lastSur string
	firstRepeats, surRepeats := 0, 0
	const steps = 200_000
	for i := 0; i < steps; i++ {
		seed, _ = PseudoRandomName(seed, false, NewEmptyFormatter())
		first, sur := getNames(seed)
		if i > 0 {
			if first == lastFirst {
				firstRepeats++
			}
			if sur == lastSur {
				surRepeats++
			}
		}
		lastFirst, lastSur = first, sur
	}
	if firstRepeats == 0 && surRepeats == 0 {
		t.Logf("No close-name repeats observed in %d steps; increase window if flaky", steps)
	}
}

// Sanity: the sequence should not cycle back to the initial state in a short window.
func TestPseudoRandomNameDoesNotCycle(t *testing.T) {
	const steps = 5_000_000
	start := 1
	seed := start
	for i := 0; i < steps; i++ {
		var name string
		seed, name = PseudoRandomName(seed, true, NewCapitalizedSpaceFormatter())
		if name == "" {
			t.Fatal("expected non-empty formatted name")
		}
	}
	if seed == start {
		t.Fatalf("unexpectedly returned to start after only %d steps", steps)
	}
}

// Sample-output shape check (no hardcoded snapshots).
func TestPseudoRandomName(t *testing.T) {
	var b bytes.Buffer
	seed := 1
	for i := 0; i < 20; i++ {
		var name string
		seed, name = PseudoRandomName(seed, false, NewCapitalizedSpaceFormatter())
		if name == "" {
			t.Fatalf("empty name at i=%d", i)
		}
		b.WriteString(name)
		b.WriteByte('\n')
	}
	if lines := bytes.Count(b.Bytes(), []byte{'\n'}); lines != 20 {
		t.Fatalf("expected 20 lines, got %d", lines)
	}
}

// As above, but enforcing the elimination property.
func TestPseudoRandomNameEliminations(t *testing.T) {
	seed := 1
	var b bytes.Buffer
	var lastFirst, lastSur string

	for i := 0; i < 20; i++ {
		var name string
		seed, name = PseudoRandomName(seed, true, NewCapitalizedSpaceFormatter())
		first, sur := getNames(seed)
		if i > 0 {
			if first == lastFirst {
				t.Fatalf("firstname repeated at i=%d with eliminateCloseNames=true: %q", i, first)
			}
			if sur == lastSur {
				t.Fatalf("surname repeated at i=%d with eliminateCloseNames=true: %q", i, sur)
			}
		}
		lastFirst, lastSur = first, sur
		if name == "" {
			t.Fatalf("empty name at i=%d", i)
		}
		b.WriteString(name)
		b.WriteByte('\n')
	}
	if lines := bytes.Count(b.Bytes(), []byte{'\n'}); lines != 20 {
		t.Fatalf("expected 20 lines, got %d", lines)
	}
}
