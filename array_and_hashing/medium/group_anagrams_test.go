package medium

import (
	"log"
	"testing"
)

func TestGroupAnagram(t *testing.T) {
	//log.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	log.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	var test [][]string
	collectChar := map[int]map[rune]rune{}

	for _, str := range strs {
		var total rune
		runes := []rune(str)
		for i := 0; i < len(runes); i++ {
			total += runes[i]
		}

		collectChar[len(runes)] = map[rune]string{}
	}

	return test
}
