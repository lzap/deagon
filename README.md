# [Deagon](https://github.com/lzap/deagon) - human readable random name generator

Out of ideas for host names in your cluster? This little Go library and a CLI can help.
Generates unique names based on frequently occurring given names and surnames
from the 1990 US Census (public domain data):

* 256 (8 bits) unique male given names
* 256 (8 bits) unique female given names
* 65,536 (16 bits) unique surnames
* with over 120 gender-neutral given names

Given names were filtered to be 3-5 characters long, surnames 5-8 characters,
therefore generated names are never longer than 14 characters (5+1+8).

This gives 33,554,432 (25 bits) total of male and female name combinations.
Built-in generator can either generate randomized succession, or generate
combinations based on MAC addresses.

Both command line utility and Go library with random or unique pseudorandom sequence
generators are available.

Similar project exists for Ruby too: https://github.com/lzap/deacon

### Command line tool

To generate a random name, just use the CLI utility:

```
# generate
Elisabeth Sobeck
```

To generate output in lower case without space:

```
# generate -d
ted-faron
```

To generate arbitrary amount of random names:

```
# generate -n 10
Tyler Vilar
Ester Boniface
Melba Forkell
Irma Paolello
Sara Stika
Pedro Dockins
Molly Stogden
Bryan Mayhue
Logan Bushner
Shane Bondi
```

To generate sequence of unique names, you must provide a starting seed number between 1 and 2^25-2 (33,554,430).
The names are guaranteed to be unique. The last line contains the next seed value that must be passed in
to continue in the unique sequence.

```
# generate -n 10 -s 130513
Jamie Abundis
Neil Abelardo
Ruben Abaunza
Teri Lebert
Tony Rizer
Vicky Sutler
Wanda Keaser
Wendy Doubek
Jim Burda
Emma Markee
18612351
```

The algorithm eliminates results with same given names or surnames, there is exactly 66046 of them, therefore the
total number of unique names available is 33,488,385. This is an implementation detail that is elaborated below.

### Go library

The library provides random generation:

```go
package main

import (
	"fmt"
	"github.com/lzap/deagon"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	
	// generate lower case with dash
	fmt.Println(deagon.RandomName(deagon.NewLowercaseDashFormatter()))
	
	// generate capitalized with space
	fmt.Println(deagon.RandomName(deagon.NewCapitalizedSpaceFormatter()))
	
	// generate pseudorandom unique sequence
	seed := 543235432
	nextSeed, name1 := deagon.PseudoRandomName(seed, true, deagon.NewCapitalizedSpaceFormatter())
	_, name2 := deagon.PseudoRandomName(nextSeed, true, deagon.NewCapitalizedSpaceFormatter())
	fmt.Println(name1, name2)
}
```

### Pseudo-random generator

Generating names randomly does not guarantee uniqueness. There is, however, a technique
called full cycle feedback register that ensures that two outputs never repeat for a given
sequence of pseudorandom numbers.

How it works, you pick a random integer between 1 and 2^25-2 (33,554,430) and pass it to a
function which returns a random name and the next number in the full cycle sequence. You need
to store the number somewhere (database) and the next time the function is called, use the
stored number and repeat.

This guarantees that two same names are only returned after 33,554,432 calls, so there is plenty
of names for everyone.

If you want more details, it is based on [Fibonacci linear feedback shift register](https://en.wikipedia.org/wiki/Linear_feedback_shift_register)
with polynomial (tap 0x10002A3):

	x^25 + x^10 + x^8 + x^6 + x^2 + x + 1.

There is exactly 66046 states when given name or surname is the same as the previous
state. Fun fact - due to nature of the pseudorandom generator, these names are only:
Aaron, Wilma, Aaberg and Zywiec (the last and the first of firstnames and surnames).
There is a boolean flag that will cause these names with she same firstname or surname
to be skipped.

## Contributing

Fork and send a Pull Request. Thanks!

## Copyright

Copyright (c) 2016 Lukas Zapletal

PUBLIC DOMAIN
