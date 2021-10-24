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
	"strings"
	"testing"
)

func TestGeneratePassphrase(t *testing.T) {
	passphrase := GeneratePassphrase(5)
	passphraseArray := strings.Split(passphrase, " ")
	passphrase2 := GeneratePassphrase(5)
	passphrase2Array := strings.Split(passphrase2, " ")

	//Check length
	if len(strings.Split(passphrase, " ")) != 5 || len(strings.Split(passphrase2, " ")) != 5 {
		t.Error("Passphrase length should be 5, but is not")
	}

	//Check Randomness
	for _, word := range passphraseArray {
		for _, word2 := range passphrase2Array {
			if word == word2 {
				t.Error("Passphrases should be random")
			}
		}
	}

	//Check prefix and suffix functionality
	passphrase = GeneratePassphrase(5, "prefix", "suffix")
	passphraseArray = strings.Split(passphrase, " ")
	passphrase2 = GeneratePassphrase(5, "", "suffix")
	passphrase2Array = strings.Split(passphrase, " ")
	if len(strings.Split(passphrase, " ")) != 7 || len(strings.Split(passphrase2, " ")) != 6 {
		t.Error("Passphrase did not add prefix/suffix properly")
	}
}

func TestGenerateLargePassphrase(t *testing.T) {
	passphrase := GenerateLargePassphrase(5)
	passphraseArray := strings.Split(passphrase, " ")
	passphrase2 := GenerateLargePassphrase(5)
	passphrase2Array := strings.Split(passphrase2, " ")

	//Check length
	if len(strings.Split(passphrase, " ")) != 5 || len(strings.Split(passphrase2, " ")) != 5 {
		t.Error("Passphrase length should be 5, but is not")
	}

	//Check Randomness
	for _, word := range passphraseArray {
		for _, word2 := range passphrase2Array {
			if word == word2 {
				t.Error("Passphrases should be random")
			}
		}
	}

	//Check prefix and suffix functionality
	passphrase = GenerateLargePassphrase(5, "prefix", "suffix")
	passphraseArray = strings.Split(passphrase, " ")
	passphrase2 = GenerateLargePassphrase(5, "", "suffix")
	passphrase2Array = strings.Split(passphrase, " ")
	if len(strings.Split(passphrase, " ")) != 7 || len(strings.Split(passphrase2, " ")) != 6 {
		t.Error("Passphrase did not add prefix/suffix properly")
	}
}

func TestGenerateMediumPassphrase(t *testing.T) {
	passphrase := GenerateMediumPassphrase(5)
	passphraseArray := strings.Split(passphrase, " ")
	passphrase2 := GenerateMediumPassphrase(5)
	passphrase2Array := strings.Split(passphrase2, " ")

	//Check length
	if len(strings.Split(passphrase, " ")) != 5 || len(strings.Split(passphrase2, " ")) != 5 {
		t.Error("Passphrase length should be 5, but is not")
	}

	//Check Randomness
	for _, word := range passphraseArray {
		for _, word2 := range passphrase2Array {
			if word == word2 {
				t.Error("Passphrases should be random")
			}
		}
	}

	//Check prefix and suffix functionality
	passphrase = GenerateMediumPassphrase(5, "prefix", "suffix")
	passphraseArray = strings.Split(passphrase, " ")
	passphrase2 = GenerateMediumPassphrase(5, "", "suffix")
	passphrase2Array = strings.Split(passphrase, " ")
	if len(strings.Split(passphrase, " ")) != 7 || len(strings.Split(passphrase2, " ")) != 6 {
		t.Error("Passphrase did not add prefix/suffix properly")
	}
}

func TestGenerateShortPassphrase(t *testing.T) {
	passphrase := GenerateShortPassphrase(5)
	passphraseArray := strings.Split(passphrase, " ")
	passphrase2 := GenerateShortPassphrase(5)
	passphrase2Array := strings.Split(passphrase2, " ")

	//Check length
	if len(strings.Split(passphrase, " ")) != 5 || len(strings.Split(passphrase2, " ")) != 5 {
		t.Error("Passphrase length should be 5, but is not")
	}

	//Check Randomness
	for _, word := range passphraseArray {
		for _, word2 := range passphrase2Array {
			if word == word2 {
				t.Error("Passphrases should be random")
			}
		}
	}

	//Check prefix and suffix functionality
	passphrase = GenerateShortPassphrase(5, "prefix", "suffix")
	passphraseArray = strings.Split(passphrase, " ")
	passphrase2 = GenerateShortPassphrase(5, "", "suffix")
	passphrase2Array = strings.Split(passphrase, " ")
	if len(strings.Split(passphrase, " ")) != 7 || len(strings.Split(passphrase2, " ")) != 6 {
		t.Error("Passphrase did not add prefix/suffix properly")
	}
}
