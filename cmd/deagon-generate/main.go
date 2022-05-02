package main

import (
	"flag"
	"fmt"
	"github.com/lzap/deagon"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var formatter deagon.Formatter
	num := flag.Int("n", 1, "number of names generated")
	dashFormat := flag.Bool("d", false, "lower case with dash format")
	initialSeed := flag.Int("s", 0, "pseudo-random seeded sequence (1 to 2^25-1)")
	flag.Parse()
	var seed int = *initialSeed

	if *dashFormat {
		formatter = deagon.NewLowercaseDashFormatter()
	} else {
		formatter = deagon.NewCapitalizedSpaceFormatter()
	}

	for i := 0; i < *num; i++ {
		if *initialSeed > 0 {
			var name string
			seed, name = deagon.PseudoRandomName(seed, true, formatter)
			fmt.Println(name)
		} else {
			fmt.Println(deagon.RandomName(formatter))
		}
	}

	if *initialSeed > 0 {
		fmt.Println(seed)
	}
}
