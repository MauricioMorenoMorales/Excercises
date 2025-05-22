package main

import (
	"fmt"
)

func main() {
	fmt.Println("-------------------------------------------------")
}

/*
	remove all values "val" in nums
	remaining elements are not important and also the size of nums
	return the count of deleted elements
*/

func removeDuplicates(nums []int) int {
	currentValue := nums[0]
	k := 1

	for i, element := range nums {
		if currentValue != element {
			nums[k] = nums[i]
			currentValue = nums[i]
			k++
		}
	}
	return k
}
