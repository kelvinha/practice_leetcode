package easy

import (
	"log"
	"testing"
)

func TestTwoSum(t *testing.T) {
	//log.Println(twoSum([]int{2, 7, 11, 15}, 9))
	//log.Println(twoSum([]int{3, 2, 4}, 6))
	log.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	listNum := []int{}
	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			result := nums[i] + nums[j]
			if result == target {
				listNum = append(listNum, i)
				listNum = append(listNum, j)
			}
		}
	}

	return listNum
}
