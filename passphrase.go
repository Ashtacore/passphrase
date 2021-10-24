/*
Copyright © 2021 Joshua Runyan

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package passphrase

import (
	"math/rand"
	"time"

	"github.com/Ashtacore/passphrase/wordlists"
)

func GeneratePassphrase(length int, prefixSuffix ...string) (passphrase string) {
	wordList := wordlists.ReadWordList()

	return buildPassphrase(length, &wordList, prefixSuffix...)
}

func GenerateLargePassphrase(length int, prefixSuffix ...string) (passphrase string) {
	wordList := wordlists.ReadWordList()

	return buildPassphrase(length, &wordList, prefixSuffix...)
}

func GenerateMediumPassphrase(length int, prefixSuffix ...string) (passphrase string) {
	wordList := wordlists.ReadWordList("medium")

	return buildPassphrase(length, &wordList, prefixSuffix...)
}

func GenerateShortPassphrase(length int, prefixSuffix ...string) (passphrase string) {
	wordList := wordlists.ReadWordList("short")

	return buildPassphrase(length, &wordList, prefixSuffix...)
}

func buildPassphrase(length int, wordList *wordlists.Wordlist, prefixSuffix ...string) (passphrase string) {
	// Create passphrase by rolling dice
	for i := 0; i < length; i++ {
		var diceRolls []int
		rand.Seed(rand.Int63n(time.Now().UnixNano()))
		for j := 0; j < wordList.Dice; j++ {
			diceRolls = append(diceRolls, rand.Intn(6)+1)
		}
		passphrase += wordList.WordMap[concatIntArray(diceRolls)] + " "
	}
	//trim trailing space
	passphrase = passphrase[:len(passphrase)-1]

	// Add prefix and suffix if provided
	if len(prefixSuffix) > 0 && prefixSuffix[0] != "" {
		passphrase = prefixSuffix[0] + " " + passphrase
	}
	if len(prefixSuffix) > 1 && prefixSuffix[1] != "" {
		passphrase += " " + prefixSuffix[1]
	}

	return passphrase
}

// Concatenate an array of ints into a single int, used for combining dice rolls
// credit: https://stackoverflow.com/a/44730447/13443483
func concatIntArray(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}
