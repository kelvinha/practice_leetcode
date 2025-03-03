package medium

import (
	"log"
	"testing"
)

func TestGroupAnagram(t *testing.T) {
	log.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	//log.Println(groupAnagrams([]string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}))
	//log.Println(groupAnagrams([]string{"duh", "ill"}))
	//log.Println(groupAnagrams([]string{"ac", "c"}))
	//log.Println(groupAnagrams([]string{"hhhhu", "tttti", "tttit", "hhhuh", "hhuhh", "tittt"}))
	//log.Println(groupAnagrams([]string{"tho", "tin", "erg", "end", "pug", "ton", "alb", "mes", "job", "ads", "soy", "toe", "tap", "sen", "ape", "led", "rig", "rig", "con", "wac", "gog", "zen", "hay", "lie", "pay", "kid", "oaf", "arc", "hay", "vet", "sat", "gap", "hop", "ben", "gem", "dem", "pie", "eco", "cub", "coy", "pep", "wot", "wee"}))
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
			//log.Printf("word: %s rune:%d", string(runes[i]), runes[i])
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
