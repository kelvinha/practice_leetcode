package interview_challenge

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestReverseInteger(t *testing.T) {
	ReverseInteger(12345678)
}

func ReverseInteger(number int) {
	fmt.Println("before: ", number)
	numberStr := strconv.Itoa(number)
	numberStrSlice := strings.Split(numberStr, "")
	totalIndex := len(numberStrSlice) - 1
	reverseNumberStrSlice := []string{}

	for i := totalIndex; i >= 0; i-- {
		reverseNumberStrSlice = append(reverseNumberStrSlice, numberStrSlice[i])
	}

	reverseNumberStr := strings.Join(reverseNumberStrSlice, "")
	reverseNumber, err := strconv.Atoi(reverseNumberStr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("after: ", reverseNumber)
}
