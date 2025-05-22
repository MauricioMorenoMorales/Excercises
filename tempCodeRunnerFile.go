package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("This is the first function")
	fmt.Println("This is the second test")
	fmt.Println(findNumbers([]int{555, 901, 482, 1771}))
}

func sortedSquares(nums []int) []int {
	response := []int{}

	for _, value := range nums {
		response = append(response, value*value)
	}

	sort.Ints(response)
	return response
}
