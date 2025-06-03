package main

import (
	"fmt"
)

func main() {
	fmt.Println("-------------------------------------------------")
	fmt.Println([]int{1, 2, 3, 4}[:3+1])
}

/*
	Un pivot index es donde la suma de todos los numeros
	de la izquierda son iguales a los de la derecha
*/

// ###

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func createListNode(numbers []int) *ListNode {
	head := &ListNode{0, nil}
	for _, value := range numbers {
		head = &ListNode{value, nil}
	}
	return head
}

func printLL(head *ListNode) {
	var response []int
	for current := head; current != nil; current = current.Next {
		response = append(response, current.Val)
	}
	fmt.Println("LinkedList: ------------------------------------------------")
	fmt.Println(response)
	fmt.Println("------------------------------------------------------------")
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//! Non descending = ascending

// *****************************************************************************************************

func containsNearbyDuplicate(nums []int, k int) bool {
	hashMap := make(map[int]int)
	for currentIndex, currentNumber := range nums {
		if foundIndex, ok := hashMap[currentNumber]; ok && abs(currentIndex, foundIndex) <= k {
			return true
		} else {
			hashMap[currentNumber] = currentIndex
		}
	}
	return false
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
