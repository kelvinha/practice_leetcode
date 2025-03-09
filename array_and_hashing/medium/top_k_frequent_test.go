package medium

import (
	"log"
	"sort"
	"testing"
)

func TestTopKFrequentElements(t *testing.T) {
	//nums := []int{1, 2}
	nums := []int{5, 2, 5, 3, 5, 3, 1, 1, 3}
	//nums := []int{1, 1, 1, 2, 2, 3}
	//nums := []int{3, 1, 0, 9, 2, 3, 1, 4, 1}
	k := 2
	log.Println(topKFrequent(nums, k))
}

func topKFrequent(nums []int, k int) []int {
	var result []int
	saveTheDuplicateNums := map[int]int{}
	type Frequent struct {
		key   int
		value int
	}

	var frequents []Frequent
	for _, num := range nums {
		if saveTheDuplicateNums[num] == num {
			saveTheDuplicateNums[num] += 1
		} else {
			saveTheDuplicateNums[num] += 1
		}
	}

	for key, v := range saveTheDuplicateNums {
		if v > 1 {
			frequents = append(frequents, Frequent{key: key, value: v})
		}
	}

	for key, v := range saveTheDuplicateNums {
		if v == 1 {
			frequents = append(frequents, Frequent{key: key, value: v})
		}
	}

	sort.Slice(frequents, func(i, j int) bool {
		return frequents[i].value > frequents[j].value
	})

	for i := 0; i < k; i++ {
		result = append(result, frequents[i].key)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	log.Println("given:", nums, "result:", result, "frequent:", k, "struct:", frequents)

	return result
}
