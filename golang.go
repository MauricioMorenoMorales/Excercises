package main

import (
	"fmt"
)

func main() {
	fmt.Println("-------------------------------------------------")
	string := "abcabcabcabc"
	fmt.Println(string[0:2])
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

type ListNode struct {
	Val  int
	Next *ListNode
}

//! Non descending = ascending

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

// *****************************************************************************************************


