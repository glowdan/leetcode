package main

func twoSum(nums []int, target int) []int {
	var data []int
	for x, xv := range nums {
		for y, yv := range nums {
			if x == y {
				continue
			}
			if xv+yv == target {
				data = append(data, x, y)
				return data
			}
		}
	}
	return data
}
