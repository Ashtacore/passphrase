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
	"bufio"
	"math/rand"
	"os"
	"strings"
)

func GeneratePassphrase(wordCount int, prefixSuffix ...string) (passphrase string, err error) {
	wordList, err := readStandardWordList()
	if err != nil {
		return "", err
	}

	// Create passphrase
	for i := 0; i < wordCount; i++ {
		passphrase += wordList[rand.Intn(len(wordList))]
		passphrase += " "
	}
	passphrase = strings.TrimSuffix(passphrase, " ")

	// Add prefix and suffix if provided
	if len(prefixSuffix) > 0 {
		passphrase = prefixSuffix[0] + " " + passphrase
	}
	if len(prefixSuffix) > 1 {
		passphrase += " " + prefixSuffix[1]
	}

	return passphrase, nil

}

func GeneratePassphraseFromWordList(wordCount int, wordList []string, prefixSuffix ...string) (passphrase string, err error) {
	// Create passphrase
	for i := 0; i < wordCount; i++ {
		passphrase += wordList[rand.Intn(len(wordList))]
		passphrase += " "
	}
	passphrase = strings.TrimSuffix(passphrase, " ")

	// Add prefix and suffix if provided
	if len(prefixSuffix) > 0 {
		passphrase = prefixSuffix[0] + " " + passphrase
	}
	if len(prefixSuffix) > 1 {
		passphrase += " " + prefixSuffix[1]
	}

	return passphrase, nil

}

func readStandardWordList() (wordList []string, err error) {
	file, err := os.Open("words_alpha.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}
	return wordList, scanner.Err()
}
