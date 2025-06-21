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

func getScores(scores []int) map[byte]int {
	letters := "abcdefghijklmnopqrstuvwxyz"
	hashMap := make(map[byte]int)
	for i, v := range scores {
		hashMap[letters[i]] = v
	}

	for key, value := range hashMap {
		if value == 0 {
			delete(hashMap, key)
		}
	}
	return hashMap
}

func getLettersHashmap(letters []byte) map[byte]int {
	hashMap := make(map[byte]int)
	for _, letter := range letters {
		hashMap[letter]++
	}
	return hashMap
}

func isValidWord(word string, letterCount map[byte]int) bool {
	letterCountCopy := make(map[byte]int)
	for k, v := range letterCount {
		letterCountCopy[k] = v
	}
	for _, letter := range word {
		if count, ok := letterCountCopy[byte(letter)]; !ok || count <= 0 {
			return false
		}
		letterCountCopy[byte(letter)]--
	}
	return true
}

func getWordValue(word string, scoresMap map[byte]int) int {
	var response int
	for _, v := range word {
		response += scoresMap[byte(v)]
	}
	return response
}

func sumElements(numbers []int) int {
	var response int
	for _, number := range numbers {
		response += number
	}
	return response
}

func isValidSubset(words []string, letterCount map[byte]int) bool {
	copyCount := make(map[byte]int)
	for k, v := range letterCount {
		copyCount[k] = v
	}

	for _, word := range words {
		for i := 0; i < len(word); i++ {
			c := word[i]
			if copyCount[c] == 0 {
				return false
			}
			copyCount[c]--
		}
	}
	return true
}

func maxScoreWords(words []string, letters []byte, score []int) int {
	scoresMap := getScores(score)
	letterCount := getLettersHashmap(letters)
	maxScore := 0

	for mask := 0; mask < (1 << len(words)); mask++ {
		subset := []string{}
		for i := 0; i < len(words); i++ {
			if mask&(1<<i) != 0 {
				subset = append(subset, words[i])
			}
		}

		if isValidSubset(subset, letterCount) {
			total := 0
			for _, word := range subset {
				total += getWordValue(word, scoresMap)
			}
			if total > maxScore {
				maxScore = total
			}
		}
	}

	return maxScore
}
