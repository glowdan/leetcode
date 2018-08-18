package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(longestPalindrome("abcbcdba"))
}

func longestPalindrome(s string) string {
	bytes := []byte(s)
	longest := 1
	var tmpLong []byte
	for xIndex := range bytes {
		part := bytes[xIndex:]
		for yIndex := range part {
			littlePart := part[:yIndex+1]
			if isPalin(littlePart, longest) {
				longest = len(littlePart)
				tmpLong = littlePart
			}
		}
	}
	return string(tmpLong)
}

func isPalin(str []byte, long int) bool {
	length := len(str)
	if length <= long {
		return false
	}
	times := int(math.Floor(float64(length) / 2.0))
	for index := 0; index < times; index++ {
		if str[index] != str[length-1-index] {
			return false
		}
	}
	return true
}
