# Hacerle flattening a una lista doblemente enlazada

#  Ejercicios a practicar
https://leetcode.com/problems/copy-list-with-random-pointer/


```go
func flatten(head *Node) *Node {
	if head == nil {
		return nil
	}
	pseudoHead := &Node{Val: -1, Prev: nil, Next: head, Child: nil}
	prev := pseudoHead

	var stack []*Node
	stack = append(stack, head)

	for len(stack) > 0 {
		// Pop
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		prev.Next = current
		current.Prev = prev

		if current.Next != nil {
			stack = append(stack, current.Next)
		}

		if current.Child != nil {
			stack = append(stack, current.Child)
			current.Child = nil
		}
		prev = current
	}

	pseudoHead.Next.Prev = nil
	return pseudoHead.Next
}
```