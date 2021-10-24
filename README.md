# passphrase

A super basic go package for generating random passphrases. Uses the Diceware algorithm and wordlists from EFF.

`go get "github.com/Ashtacore/passphrase"`

## Example Usage
```go
package main

import (
	"github.com/Ashtacore/passphrase"
)

func main() {
	// Standard method, GeneratePassphrase(int, ...string) uses eff_large_wordlist.txt and alows you to specify the amount of words generated
	// eff_large_wordlist.txt is the recomended wordlist from the EFF. It has 7776 words.
	println(passphrase.GeneratePassphrase(5))
    // Sample output: "ploy appendix squeezing haste scarily"

	// You may also specify a prefix and suffix for the passphrase
	println(passphrase.GeneratePassphrase(5, "prefix", "suffix"))
    // Sample output: "prefix poppy crisply backless reword region suffix"

	// The first string provided will always be a prefix, so if you'd like only a suffix you can provide an empty string first
	println(passphrase.GeneratePassphrase(5, "", "suffix"))
    // Sample output: "sliceable collage ardently bullwhip mummy suffix"

	// GenerateLargePassphrase(int, ...string) is an alias for GeneratePassphrase(int, ...string)
	println(passphrase.GenerateLargePassphrase(5, "prefix", "suffix"))
    // Sample output: "prefix aside justly angelic cathedral acquaint suffix"

	// GenerateMediumPassphrase(int, ...string) uses eff_short_wordlist_2.txt and functions the same as GeneratePassphrase(int, ...string)
	// eff_short_wordlist_2.txt has 1296 words which are a selection of some longer/easier to memorize words from the large word list
	println(passphrase.GenerateMediumPassphrase(5, "prefix", "suffix"))
    // Sample output: "prefix enactment rugby yearbook quilt osmosis suffix"

	// GenerateSmallPassphrase(int, ...string) uses eff_short_wordlist_1.txt and functions the same as GeneratePassphrase(int, ...string)
	// eff_short_wordlist_2.txt has 1296 words which are a selection of some shorter words from the large word list
	println(passphrase.GenerateShortPassphrase(5, "prefix", "suffix"))
    // Sample output: "prefix mace uncut charm props boned suffix"
}
```

## How Random Is This?
The Diceware algorithm relies on rolling dice. This is implemented in Go with `rand.Intn(6)+1`. Of course math/rand will produce predictable results unless the seed is randomized, so each time this package is called a new seed is generated with `rand.Seed(rand.Int63n(time.Now().UnixNano()))`. 

I make no guarentees about the security of passphrases generated with this package, but I've tried to make my code as simple to read as possible so you can make an educated decision whether you want to use this package.

## What If I Think Your Implementation of Diceware Sucks?
Then you could use the wordlists package on it's own. Get it with `go get "github.com/Ashtacore/passphrase/wordlists"`

This package exports a single function, `ReadWordList(...string)`. This function takes a single string which will pick one of the three wordlists provided by the EFF. The following table explains which wordlist is returned based on input.

|Input|Returned Wordlist|
|--|--|
|short|eff_short_wordlist_1.txt|
|short1|eff_short_wordlist_1.txt|
|medium|eff_short_wordlist_2.txt|
|short2|eff_short_wordlist_2.txt|
|large|eff_large_wordlist.txt|
||eff_large_wordlist.txt|
|{anything else}|eff_large_wordlist.txt|

This will return a struct like follows, where `Dice` is the amount of dice rolls to make and `WordMap` is a map of words related to dice rolls.
```go
type Wordlist struct {
	Dice    int
	WordMap map[int]string
}
```

## References

World list and algorithm from EFF at https://www.eff.org/dice
Deeper explaination of Diceware algorithm and EFF wordlists: https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases