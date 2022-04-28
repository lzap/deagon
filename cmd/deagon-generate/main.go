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
	num := flag.Int("n", 1, "number of names generated")
	flag.Parse()

	for i := 0; i < *num; i++ {
		fmt.Println(deagon.RandomName(deagon.NewUppercaseSpaceFormatter()))
	}
}
