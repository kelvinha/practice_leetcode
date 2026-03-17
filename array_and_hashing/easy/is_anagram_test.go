package easy

import (
	"log"
	"sort"
	"strings"
	"testing"
)

func TestAnagram(t *testing.T) {
	//log.Println(validationAnagram("aacc", "ccac"))
	log.Println(validationAnagramV2("bbcc", "ccbc"))
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

		log.Println("test passed")

		return true
	}

	return false
}

func validationAnagramV2(s string, t string) bool {
	anagramSlice := strings.Split(s, "")
	secondAnagramSlice := strings.Split(t, "")

	if len(anagramSlice) != len(secondAnagramSlice) {
		return false
	}

	// sorting
	sort.Slice(anagramSlice, func(i, j int) bool {
		return anagramSlice[i] < anagramSlice[j]
	})

	// sorting
	sort.Slice(secondAnagramSlice, func(i, j int) bool {
		return secondAnagramSlice[i] < secondAnagramSlice[j]
	})

	s = strings.Join(anagramSlice, "")
	t = strings.Join(secondAnagramSlice, "")

	if strings.Compare(s, t) != 0 {
		return false
	}

	return true
}
