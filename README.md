# [Deagon](https://github.com/lzap/deagon) - Human-Readable Random Name Generator

Out of ideas for hostnames in your cluster? This little Go library and CLI can
help. It generates unique names based on frequently occurring given names and
surnames from the 1990 US Census (public domain data):

* 256 (8 bits) unique male given names
* 256 (8 bits) unique female given names
* 65,536 (16 bits) unique surnames
* Includes over 120 gender-neutral given names

Given names are filtered to be 3–5 characters long and surnames 5–8 characters.
Therefore, generated names are never longer than 14 characters (5+1+8).

This provides a total of 33,554,432 (2^25) combinations of male and female names.
The built-in generator can produce either a pseudo-random unique succession of
names (using a full-cycle linear feedback shift register) or generate
pseudo-random names seeded by the current time.

A command-line utility and a Go library are both available, providing random or
unique pseudo-random sequence generation.

A similar project also exists for Ruby: https://github.com/lzap/deacon

### Installation

```
# go install github.com/lzap/deagon@latest
```

### Command line tool

To generate a random name, just use the CLI utility:

```
# generate
Elisabeth Sobeck
```

To generate output in lowercase with a dash instead of a space:

```
# generate -d
ted-faron
```

To generate an arbitrary number of random names:

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

To generate a sequence of unique names, you must provide a starting seed number
between 1 and 2^25 - 2 (33,554,430). The names are guaranteed to be unique
within a full cycle. The last line of the output contains the next seed value
that must be passed in to continue the unique sequence.

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

The algorithm can optionally eliminate pairs with the same given name or surname
as the previous entry. There are exactly 66,046 such pairs, which reduces the
total number of unique names in the sequence to 33,488,385. This feature
addresses the fact that the pseudo-random generator, by its nature, will
occasionally produce consecutive names that share a first name or surname. See
below for details.

### Go library

The library provides both random and sequential name generation:

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

	// Generate lower case with a dash
	fmt.Println(deagon.RandomName(deagon.NewLowercaseDashFormatter()))

	// Generate capitalized with a space
	fmt.Println(deagon.RandomName(deagon.NewCapitalizedSpaceFormatter()))

	// Generate from an integer (only the lower 25 bits are used)
	fmt.Println(deagon.Name(deagon.NewCapitalizedSpaceFormatter(), 130513))

	// Generate a pseudo-random unique sequence
	seed := 543235432
	nextSeed, name1 := deagon.PseudoRandomName(seed, true, deagon.NewCapitalizedSpaceFormatter())
	_, name2 := deagon.PseudoRandomName(nextSeed, true, deagon.NewCapitalizedSpaceFormatter())
	fmt.Println(name1, name2)
}
```

### Pseudo-Random Generator

Generating names randomly does not guarantee uniqueness. There is, however, a
technique called a Linear Feedback Shift Register (LFSR) with a full-cycle
polynomial that ensures no output is repeated until the entire sequence of
numbers is exhausted.

It works by picking an integer between 1 and 2^25 - 2 (33,554,430) and passing
it to a function. The function returns a random name and the next number in the
full-cycle sequence. You need to store this number, and the next time you call
the function, use the stored number to continue the sequence.

This guarantees that the same name is only returned after 33,554,432 calls, so
there are plenty of names for everyone. The sequence is guaranteed (and tested)
to never return the same name within a full cycle, so there is no need to
perform an external uniqueness check.

The implementation is based on a [Fibonacci linear feedback shift
register](https://en.wikipedia.org/wiki/Linear_feedback_shift_register) with
the polynomial `x^25 + x^10 + x^8 + x^6 + x^2 + x + 1` (tap `0x10002A3`).

Within the sequence, there are exactly 66,046 instances where a generated name
has the same given name or surname as the name from the previous state. This is
not a flaw but a mathematical property of the LFSR sequence. Fun fact: due to
the structure of the name lists, these repetitions only occur with the names
Aaron, Wilma, Aaberg, and Zywiec (the first and last entries of the first name
and surname lists). A boolean flag is available in the `PseudoRandomName`
function to skip these entries if a strictly unique first and last name is
required between consecutive calls.

## Contributing

Fork and send a Pull Request. Thanks!

## Copyright

Copyright (c) 2016 Lukas Zapletal

PUBLIC DOMAIN
