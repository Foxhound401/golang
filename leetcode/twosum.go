package main

import (
	"fmt"
)

func main() {
	nums, target := []int{2, 7, 11, 15}, 9
	fmt.Println(twosum(nums, target))
}

func twosum(nums []int, target int) []int {
	mapNums := make(map[int]int)

	var ans []int
	for i, num := range nums {
		var complement int = target - num
		if c, ok := mapNums[complement]; ok {
			ans = []int{c, i}
			break
		}
		mapNums[num] = i
	}

	return ans
}
