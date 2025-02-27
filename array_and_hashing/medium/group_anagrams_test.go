package medium

import (
	"log"
	"testing"
)

func TestGroupAnagram(t *testing.T) {
	//log.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	//log.Println(groupAnagrams([]string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}))
	//log.Println(groupAnagrams([]string{"ill", "duh"}))
	log.Println(groupAnagrams([]string{"duh", "ban"}))

	// [["max"],["buy"],["doc"],["may"],["ill"],["duh"],["tin"],["bar"],["pew"],["cab"]]
}

func groupAnagrams(strs []string) [][]string {
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
