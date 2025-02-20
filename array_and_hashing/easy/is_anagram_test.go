package easy

import (
	"log"
	"testing"
)

func TestAnagram(t *testing.T) {
	log.Println(validationAnagram("aacc", "ccac"))
	//log.Println(validationAnagram("anagram", "nagaram"))
}

// anagram => angrm => 5
// nagaram => nagrm => 5

// aacc => ac => 2
// ccac => ca => 2

func validationAnagram(s, t string) bool {
	runesAnagram1 := []rune(s)
	runesAnagram2 := []rune(t)

	if len(runesAnagram1) == len(runesAnagram2) {
		var isValid []bool
		anagramMaps := map[rune]rune{}
		duplicateCharRunesAnagram1 := map[rune]int{}
		duplicateCharRunesAnagram2 := map[rune]int{}

		// save all text to map
		for i := 0; i < len(runesAnagram1); i++ {
			if anagramMaps[runesAnagram1[i]] == runesAnagram1[i] {
				duplicateCharRunesAnagram1[runesAnagram1[i]] = duplicateCharRunesAnagram1[runesAnagram1[i]] + 1
			}

			anagramMaps[runesAnagram1[i]] = runesAnagram1[i]
		}

		// validate if runes from string 2 having runes in map from string 1
		for i := 0; i < len(runesAnagram2); i++ {
			if anagramMaps[runesAnagram2[i]] == runesAnagram2[i] {
				duplicateCharRunesAnagram2[runesAnagram2[i]] = duplicateCharRunesAnagram2[runesAnagram2[i]] + 1
				isValid = append(isValid, true)
			}
		}

		log.Println("test: ", duplicateCharRunesAnagram1)

		if len(isValid) == len(t) {
			return true
		}

		return false
	}

	return false
}
