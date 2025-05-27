package main

import (
	"fmt"
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

// *****************************************************************************************************

// Encuentra el final de la primera mitad usando técnica tortuga y liebre
func endOfFirstHalf(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

// Revierte una lista enlazada
func reverseList(head *ListNode) *ListNode {
	var previous *ListNode

	for current := head; current != nil; current = current.Next {
		previous = &ListNode{Val: current.Val, Next: previous}
	}

	return previous
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	// Encontrar el final de la primera mitad y revertir la segunda mitad
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)

	// Verificar si hay un palíndromo
	result := true
	firstPosition := head
	secondPosition := secondHalfStart

	for secondPosition != nil {
		if firstPosition.Val != secondPosition.Val {
			result = false
			break
		}
		firstPosition = firstPosition.Next
		secondPosition = secondPosition.Next
	}

	// Restaurar la lista original y retornar el resultado
	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}
