package easy

import (
	"log"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	log.Println(duplicate([]int{0}))
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
