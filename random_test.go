package deagon

import "testing"

func TestRandomName(t *testing.T) {
	f := NewCapitalizedSpaceFormatter()

	// Basic: non-empty names
	const n = 200
	seen := make(map[string]struct{}, n)
	for i := 0; i < n; i++ {
		got := RandomName(f)
		if got == "" {
			t.Fatalf("RandomName returned empty string at i=%d", i)
		}
		seen[got] = struct{}{}
	}

	// Heuristic: we should observe more than one unique value.
	// (This is intentionally weak to avoid flakiness.)
	if len(seen) < 2 {
		t.Fatalf("expected at least 2 distinct names in %d samples, got %d", n, len(seen))
	}
}
