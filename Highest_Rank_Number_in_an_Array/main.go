package main

import (
	"fmt"
	"sort"
)

func HighestRank(nums []int) int {
	sort.Ints(nums)
	var maxRepeat int
	minRepeat := 1
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			minRepeat++
		} else {
			minRepeat = 1
		}
		if maxRepeat <= minRepeat {
			maxRepeat = minRepeat
			result = nums[i]
		}
	}
	return result
}

func main() {
	fmt.Println(HighestRank([]int{2, 1, 2, 2}))
}
