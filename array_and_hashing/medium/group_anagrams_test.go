package medium

import (
	"log"
	"sort"
	"testing"
)

func TestGroupAnagram(t *testing.T) {
	log.Println(groupAnagramsV2([]string{"eat", "tea", "tan", "ate", "nat", "bat", "hay", "rig", "hay"}))
	//log.Println(groupAnagramsV1([]string{"eat", "tea", "tan", "ate", "nat", "bat", "hay", "rig", "hay"}))
	//log.Println(groupAnagramsV1([]string{"ray", "cod", "abe", "ned", "arc", "jar", "owl", "pop", "paw", "sky", "yup", "fed", "jul", "woo", "ado", "why", "ben", "mys", "den", "dem", "fat", "you", "eon", "sui", "oct", "asp", "ago", "lea", "sow", "hus", "fee", "yup", "eve", "red", "flo", "ids", "tic", "pup", "hag", "ito", "zoo"}))
}

func groupAnagramsV1(strs []string) [][]string {
	var result [][]string
	collectChar := map[rune]map[string]rune{}
	groupingAnagram := map[rune][]string{}

	for _, word := range strs {
		var keyRune rune
		valueChar := map[string]rune{}
		runes := []rune(word) // convert to list of rune (code character)

		// loop each char for set up the key and set the value like 97 -> a
		for i := 0; i < len(runes); i++ {
			keyRune += runes[i]
			valueChar[string(runes[i])] = runes[i]
		}

		// TODO: keyRune diubah jadi sort asc
		// TODO: handling duplicate key

		// validate the key
		if _, ok := collectChar[keyRune]; ok {
			var validateCharacter []bool
			for i := 0; i < len(runes); i++ {
				if collectChar[keyRune][string(runes[i])] == runes[i] {
					validateCharacter = append(validateCharacter, true)
				}
			}

			if len(validateCharacter) == len(runes) {
				groupingAnagram[keyRune] = append(groupingAnagram[keyRune], word)
			} else {
				result = append(result, []string{string(runes)})
			}
		} else {
			collectChar[keyRune] = valueChar
			groupingAnagram[keyRune] = []string{word}
		}
	}

	for _, anagram := range groupingAnagram {
		result = append(result, anagram)
	}

	return result
}

func groupAnagramsV2(strs []string) [][]string {
	var result [][]string
	groupingMap := map[string][]string{}

	for _, word := range strs {
		runeWord := []rune(word)
		sort.Slice(runeWord, func(i, j int) bool {
			return runeWord[i] < runeWord[j]
		})

		groupingMap[string(runeWord)] = append(groupingMap[string(runeWord)], word)

	}
	for _, anagram := range groupingMap {
		result = append(result, anagram)
	}

	log.Println(result)

	return result
}
