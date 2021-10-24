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
	"os"
	"strconv"
	"strings"
)

// ReadWordList returns a map of dicerolls to words based on user choice of EFF wordlist
// chosenList takes a single string relating to the chosen wordlist. It is implemented as an interface to allow for the user to provide nothing as a default.
func ReadWordList(chosenList ...string) (wordList map[int]string, err error) {
	var fileName string
	switch chosenList[0] {
	case "large":
		fileName = "./eff_large_wordlist.txt"
	case "medium", "short2":
		fileName = "./eff_short_wordlist_2.txt"
	case "short", "short1":
		fileName = "./eff_short_wordlist_1.txt"
	default:
		fileName = "./eff_large_wordlist.txt"
	}

	file, err := os.Open(fileName)
	if err != nil {
		return wordList, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) == 2 {
			dice, err := strconv.Atoi(line[0])
			if err != nil {
				return wordList, err
			}
			wordList[dice] = line[1]
		}
	}
	return wordList, scanner.Err()
}
