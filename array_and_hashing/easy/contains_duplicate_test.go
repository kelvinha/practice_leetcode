package easy

import (
	"log"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	log.Println(duplicateV2([]int{0, 1, 2, 3, 3}))
}

func duplicate(nums []int) bool {
	numsNotDuplicate := map[int]interface{}{}
	for _, num := range nums {
		if numsNotDuplicate[num] == num {
			return true
		}
		numsNotDuplicate[num] = num
	}
	return false
}

func duplicateV2(nums []int) bool {
	numbers := map[int]int{}
	for _, num := range nums {
		if _, ok := numbers[num]; ok {
			if numbers[num] > 0 {
				return true
			}
		}
		numbers[num] = 1
	}

	return false
}
