package interview_challenge

import (
	"fmt"
	"strings"
	"testing"
)

func TestReverseString(t *testing.T) {
	ReverseString("kelvin")
}

func ReverseString(data string) {
	sliceData := strings.Split(data, "")
	totalSliceData := len(sliceData) - 1
	reverseSliceData := []string{}

	for i := totalSliceData; i >= 0; i-- {
		reverseSliceData = append(reverseSliceData, sliceData[i])
	}

	reverseString := strings.Join(reverseSliceData, "")

	fmt.Println("before: ", data)
	fmt.Println("after: ", reverseString)

	return
}
