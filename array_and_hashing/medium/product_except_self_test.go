package medium

import (
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	productExceptSelf([]int{1, 2, 3, 4})
}

func productExceptSelf(nums []int) []int {
	var finalResult []int
	for i := 0; i < len(nums); i++ {
		result := 1
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			result = result * nums[j]
		}
		finalResult = append(finalResult, result)
	}
	return finalResult
}
