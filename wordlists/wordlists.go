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
package wordlists

type Wordlist struct {
	Dice    int
	WordMap map[int]string
}

func ReadWordList(chosenList ...string) (wordList Wordlist) {
	if len(chosenList) == 0 {
		return effLargeWordlist()
	} else {
		switch chosenList[0] {
		case "large":
			return effLargeWordlist()
		case "medium", "short2":
			return effMediumWordlist()
		case "short", "short1":
			return effShortWordlist()
		default:
			return effLargeWordlist()
		}
	}
}
