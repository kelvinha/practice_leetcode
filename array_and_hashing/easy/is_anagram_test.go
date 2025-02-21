package easy

import (
	"log"
	"testing"
)

func TestAnagram(t *testing.T) {
	log.Println(validationAnagram("aacc", "ccac"))
	//log.Println(validationAnagram("anagram", "nagaram"))
}

func validationAnagram(s, t string) bool {
	runesAnagram1 := []rune(s)
	runesAnagram2 := []rune(t)

	if len(runesAnagram1) == len(runesAnagram2) {
		anagramMaps := map[rune]rune{}
		duplicateCharRunesAnagram1 := map[rune]int{}
		duplicateCharRunesAnagram2 := map[rune]int{}

		// save all text to map
		for i := 0; i < len(runesAnagram1); i++ {
			anagramMaps[runesAnagram1[i]] = runesAnagram1[i]
			if anagramMaps[runesAnagram1[i]] == runesAnagram1[i] {
				duplicateCharRunesAnagram1[runesAnagram1[i]] = duplicateCharRunesAnagram1[runesAnagram1[i]] + 1
			}
		}

		// validate if runes from string 2 having runes in map from string 1
		for i := 0; i < len(runesAnagram2); i++ {
			if anagramMaps[runesAnagram2[i]] == runesAnagram2[i] {
				duplicateCharRunesAnagram2[runesAnagram2[i]] = duplicateCharRunesAnagram2[runesAnagram2[i]] + 1
			}
		}

		// validate if each character from anagram2 is available in anagram1 or not
		for _, v := range anagramMaps {
			if duplicateCharRunesAnagram1[v] != duplicateCharRunesAnagram2[v] {
				return false
			}
		}

		return true
	}

	return false
}
