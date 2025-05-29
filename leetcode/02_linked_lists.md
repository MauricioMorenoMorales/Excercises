Las listas enlazadas tienen mal performance a la hora de accesar a informacion, pero mejor performance a la hora de realizar operaciones como agregar o eliminar

Para agregar un valor se siguen estos pasos
1. Inicializas `cur` con el valor deseado
2. `cur.next->prev.next`
3. `prev.next->cur.`

# Agregar
Para agregar un nodo al inicio hacemos esto
1. Inicializamos nuevo nodo `cur`
2. `cur.next->head`
3. `head=cur`

# Eliminar
Para eliminar un nodo se sigue esto, la complejidad de esto es O(N), ya que se tiene que buscar data o posiciones dentro de la lista enlazada
1. Buscamos el nodo previo `prev` y el nodo siguiente `next`
2. Linkeamos `prev.next->cur.next`

Para eliminar el primer nodo solamente hacemos que el head sea el siguiente valor
1. `headh=head.next`

# Implementacion listas enlazadas en golang
```go
type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Head *Node
	Size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{&Node{}, 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1
	}
	curr := this.Head
	for i := 0; i <= index; i++ {
		curr = curr.Next
	}
	return curr.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{val, this.Head.Next}
	this.Head.Next = node
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{val, nil}
	curr := this.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = node
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Size {
		return
	}
	node := &Node{val, nil}
	curr := this.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	node.Next = curr.Next
	curr.Next = node
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	curr := this.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	this.Size--
}
```

# Double linked list
Es una lista en la que los nodos tienen dos punteros uno para el next y otro para el prev

## Agregar elemento a lista doblemente enlazada
1 Linkea `cur->prev` y cur`cur->next`

# Estrategias comunes

## Reverse linked list

```go
func reverseList(head *ListNode) *ListNode {
	var response *ListNode
	for current := head; current != nil; current = current.Next {
		response = &ListNode{Val: current.Val, Next: response}
	}
	return response
}
```
```go
//! Recursion
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	response := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return response
}

```

## Dummy pointer
Se usa para ser retornado y usar distintos pointers que van mutando distintos puntos de la linked list

Para eliminar un elemento dependiendo de que tan cerca esté del final

Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
Example 2:

Input: head = [1], n = 1
Output: []
Example 3:

Input: head = [1,2], n = 1
Output: [1]

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first := dummy
	second := dummy

	for i := 1; i <= n+1; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

```
```go
// Otra forma de llamar al dummy pointer es sentinel, tambien puedes usarlo para hacer que un pointer vaya hacia atras
func removeElement(head *ListNode, val int) *ListNode {
	sentinel := &ListNode{0, head}
	left := sentinel
	right := head

	for ; right != nil; right = right.Next {
		if right.Val == val {
			left.Next = left.Next.Next
		} else {
			left = left.Next
		}
	}
	return sentinel.Next
}
```

```go
// El dummy pointer tambien sirve para crear listas enlazadas sin hacer el reverse
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	response := &ListNode{-1, nil}
	currentNode := response
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			currentNode.Next = list1
			list1 = list1.Next
		} else {
			currentNode.Next = list2
			list2 = list2.Next
		}
		currentNode = currentNode.Next
	}

	if list1 != nil {
		currentNode.Next = list1
	} else if list2 != nil {
		currentNode.Next = list2
	}
	return response.Next
}

```

# Array technique
Si vas a recorrer distintas direcciones y no mutaras la data o cambiaras direcciones es buena esta tecnica

# ! Advanced themes to review
Advanced recursive method to palindrome array
https://leetcode.com/problems/palindrome-linked-list/editorial/


