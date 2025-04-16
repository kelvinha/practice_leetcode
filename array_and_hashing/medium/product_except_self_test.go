package medium

import (
	"log"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	log.Println("V1", productExceptSelf([]int{1, 2, 3, 4}))
	log.Println("V2", productExceptSelfV2([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {
	var finalResult []int
	var result int
	total := len(nums)
	for i := 0; i < total; i++ {
		result = 1
		for j := 0; j < total; j++ {
			if i == j {
				continue
			}
			result *= nums[j]
		}
		finalResult = append(finalResult, result)
	}
	return finalResult
}

func productExceptSelfV2(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	prefix := 1
	for i := 0; i < n; i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}
