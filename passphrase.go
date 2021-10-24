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
)

func GeneratePassphrase(length int, prefixSuffix ...string) (passphrase string, err error) {
	wordList, err := ReadWordList()
	if err != nil {
		return "", err
	}

	// Create passphrase by rolling dice
	for i := 0; i < length; i++ {
		var diceRolls []int
		for j := 0; j < 6; j++ {
			diceRolls[j] = rand.Intn(6) + 1
		}
		passphrase += wordList[concatIntArray(diceRolls)]
	}

	// Add prefix and suffix if provided
	if len(prefixSuffix) > 0 {
		passphrase = prefixSuffix[0] + " " + passphrase
	}
	if len(prefixSuffix) > 1 {
		passphrase += " " + prefixSuffix[1]
	}

	return passphrase, nil
}

func GenerateMediumPassphrase(length int, prefixSuffix ...string) (passphrase string, err error) {
	wordList, err := ReadWordList("medium")
	if err != nil {
		return "", err
	}

	// Create passphrase by rolling dice
	for i := 0; i < length; i++ {
		var diceRolls []int
		for j := 0; j < 4; j++ {
			diceRolls[j] = rand.Intn(6) + 1
		}
		passphrase += wordList[concatIntArray(diceRolls)]
	}

	// Add prefix and suffix if provided
	if len(prefixSuffix) > 0 {
		passphrase = prefixSuffix[0] + " " + passphrase
	}
	if len(prefixSuffix) > 1 {
		passphrase += " " + prefixSuffix[1]
	}

	return passphrase, nil
}

func GenerateShortPassphrase(length int, prefixSuffix ...string) (passphrase string, err error) {
	wordList, err := ReadWordList("short")
	if err != nil {
		return "", err
	}

	// Create passphrase by rolling dice
	for i := 0; i < length; i++ {
		var diceRolls []int
		for j := 0; j < 4; j++ {
			diceRolls[j] = rand.Intn(6) + 1
		}
		passphrase += wordList[concatIntArray(diceRolls)]
	}

	// Add prefix and suffix if provided
	if len(prefixSuffix) > 0 {
		passphrase = prefixSuffix[0] + " " + passphrase
	}
	if len(prefixSuffix) > 1 {
		passphrase += " " + prefixSuffix[1]
	}

	return passphrase, nil
}

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
