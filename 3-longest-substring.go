package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("asjrgapa"))
	fmt.Println(lengthOfLongestSubstring("bbbb"))
}

func lengthOfLongestSubstring(s string) int {
	longest := 0
	bytes := []byte(s)
	tmp := make(map[byte]int)
	var shortBytes, deleteBytes []byte

	for _, v1 := range bytes {
		_, ok := tmp[v1]

		if ok {
			for sbIndex, sbValue := range shortBytes {

				if sbValue == v1 {
					shortBytes = append(shortBytes, v1)
					deleteBytes = shortBytes[:sbIndex]
					shortBytes = shortBytes[sbIndex+1:]
					for _, db := range deleteBytes {
						if db != v1 {
							delete(tmp, db)
						}

					}
					break
				}
			}
		} else {
			tmp[v1] = 1
			shortBytes = append(shortBytes, v1)
			if len(shortBytes) > longest {
				longest = len(shortBytes)
			}
		}
	}

	return longest
}
