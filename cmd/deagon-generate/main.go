package main

import (
	"flag"
	"fmt"

	"github.com/oakroots/deagon"
)

func main() {
	// Command-line flags
	num := flag.Int("n", 1, "number of names to generate")
	dashFormat := flag.Bool("d", false, "use lowercase with dash format")
	initialSeed := flag.Int("s", 0, "pseudo-random seeded sequence (1 to 2^25-1)")
	flag.Parse()

	// Choose the formatter based on the flag
	var formatter deagon.Formatter
	if *dashFormat {
		formatter = deagon.NewLowercaseDashFormatter()
	} else {
		formatter = deagon.NewCapitalizedSpaceFormatter()
	}

	seed := *initialSeed

	// Generate names
	for i := 0; i < *num; i++ {
		if *initialSeed > 0 {
			// Deterministic pseudo-random sequence (your own function)
			var name string
			seed, name = deagon.PseudoRandomName(seed, true, formatter)
			fmt.Println(name)
		} else {
			// Non-deterministic random name using math/rand/v2
			fmt.Println(deagon.RandomName(formatter))
		}
	}

	// Print the final seed so the sequence can be continued
	if *initialSeed > 0 {
		fmt.Println(seed)
	}
}
