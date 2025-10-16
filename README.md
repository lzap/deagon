# Deagon — human readable random name generator

> **Note**  
> This is a fork of [lzap/deagon](https://github.com/lzap/deagon).  
> The maintained Go module is now published as: **`github.com/oakroots/deagon`**

Out of ideas for host names in your cluster? This little Go library and a CLI
can help. It generates unique names based on frequently occurring given names and
surnames from the 1990 US Census (public domain data):

- 256 (8 bits) unique male given names
- 256 (8 bits) unique female given names
- 65,536 (16 bits) unique surnames
- with over 120 gender-neutral given names

Given names were filtered to be 3–5 characters long, surnames 5–8 characters,
therefore generated names are never longer than 14 characters (5+1+8).

This gives **33,554,432 (25 bits)** total male and female name combinations.
The built-in generator can either generate:
- a pseudo-random **unique succession** (full-cycle linear feedback shift register), or
- pseudo-random names via a time/random source.

Both a command-line utility and a Go library are available.

A similar project exists for Ruby: https://github.com/lzap/deacon

---

## Changes in this fork

- Module path changed to **`github.com/oakroots/deagon`**.
- Switched random API to **`math/rand/v2`** (no global seeding, no deprecated `rand.Seed`).
- Fixed and clarified LFSR implementation (25-bit Fibonacci LFSR with proper feedback parity).
- Modernized tests to property-based checks (fast, stable; optional long test with `-short` skip).
- Code comments standardized in English; minor formatting utilities added.

---

## Installation

```bash
go install github.com/oakroots/deagon@latest
```

---

## Command-line tool

Generate a random name:

```bash
# generate
Elisabeth Sobeck
```

Lowercase with a dash:

```bash
# generate -d
ted-faron
```

Generate many:

```bash
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

Generate a **unique sequence**: provide a starting seed between **1** and **2^25−2 (33,554,430)**.  
Names are guaranteed unique across the sequence. The last output line prints the **next seed** to continue the sequence.

```bash
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

The algorithm eliminates pairs with the same given name or the same surname between consecutive states.
There are exactly **66,046** such states, so the total number of unique consecutive names is **33,488,385**.

---

## Go library

```go
package main

import (
	"fmt"

	"github.com/oakroots/deagon"
)

func main() {
	// lowercase with a dash
	fmt.Println(deagon.RandomName(deagon.NewLowercaseDashFormatter()))

	// capitalized with space
	fmt.Println(deagon.RandomName(deagon.NewCapitalizedSpaceFormatter()))

	// generate from an integer (only low 25 bits are used)
	fmt.Println(deagon.Name(deagon.NewCapitalizedSpaceFormatter(), 130513))

	// generate pseudorandom unique sequence
	seed := 543235432
	nextSeed, name1 := deagon.PseudoRandomName(seed, true, deagon.NewCapitalizedSpaceFormatter())
	_, name2 := deagon.PseudoRandomName(nextSeed, true, deagon.NewCapitalizedSpaceFormatter())
	fmt.Println(name1, name2)
}
```

---

## Pseudo-random generator

Generating names randomly does not guarantee uniqueness. There is, however, a
technique called full cycle feedback register that ensures that two outputs
never repeat for a given sequence of pseudorandom numbers until all numbers
are exhausted.

How it works: you pick a random integer between 1 and 2^25-2 (33,554,430) and
pass it to a function which returns a random name and the next number in the
full cycle sequence. You need to store the number somewhere (e.g. a database)
and the next time the function is called, use the stored number and repeat.

This guarantees that two same names are only returned after 33,554,432 calls,
so there is plenty of names for everyone. This is guaranteed (and tested) to
never return the same name so there is no need to do uniqueness checks.

It is based on a [Fibonacci linear feedback shift
register](https://en.wikipedia.org/wiki/Linear_feedback_shift_register) with
polynomial (tap 0x10002A3):

```
x^25 + x^10 + x^8 + x^6 + x^2 + x + 1
```

There are exactly 66,046 states when given name or surname is the same as the
previous state. Fun fact – due to nature of the pseudorandom generator, these
names are only: Aaron, Wilma, Aaberg and Zywiec (the last and the first of
firstnames and surnames). There is a boolean flag that will cause these names
with the same firstname or surname to be skipped.

---

## Contributing

Fork and send a Pull Request. Thanks!

---

## Copyright

Original work: © 2016 Lukas Zapletal  
Forked and maintained under: © 2025 Oakroots

PUBLIC DOMAIN
