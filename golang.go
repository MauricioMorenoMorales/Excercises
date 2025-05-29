package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("-------------------------------------------------")
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
	Val  int
	Next *Node
}

//! Non descending = ascending

// *****************************************************************************************************

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		response := &Node{Val: x}
		response.Next = response
		return response
	}
	if aNode.Next == aNode {
		node := &Node{Val: x, Next: aNode}
		aNode.Next = node
		return aNode
	}
	current := aNode
	iterations := 0
	maxValue := math.MinInt
	for {
		if current.Val <= x && x <= current.Next.Val {
			break
		}

		// Same number cases
		if current.Val == current.Next.Val {
			break
		}

		if x == 0 {
			if current.Next == aNode {
				iterations++
			}
			if current.Val == maxValue && iterations == 2 {
				break
			}

			//? Checks for max value
			if current.Val > maxValue {
				maxValue = current.Val
			}
		}
	}
	newNode := &Node{Val: x, Next: current.Next}
	current.Next = newNode
	return aNode
}
