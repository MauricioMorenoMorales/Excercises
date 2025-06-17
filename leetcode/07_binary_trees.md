# Traversals

## Preorder Traversals
```go
func preorderTraversal(root *TreeNode) []int {
	var response []int
	var traverseTree func(node *TreeNode)

	traverseTree = func(head *TreeNode) {
		if head == nil {
			return
		}

		response = append(response, head.Val)
		traverseTree(head.Left)
		traverseTree(head.Right)
	}
	traverseTree(root)

	return response
}

func preorderTraversalIteration(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{root}

	output := []int{}

	for len(stack) > 0 {
		length := len(stack)
		root = stack[length-1] //! POP
		stack = stack[:length-1]

		if root != nil {
			output = append(output, root.Val)
		}

		if root.Right != nil {
			stack = append(stack, root.Right)
		}

		if root.Left != nil {
			stack = append(stack, root.Left)
		}
	}

	return output
}

func morrisPreorderTraversal(root *TreeNode) []int {
	var output []int
	node := root
	for node != nil {
		if node.Left == nil {
			output = append(output, node.Val)
			node = node.Right
		} else {
			predecessor := node.Left
			for predecessor.Right != nil && predecessor.Right != node {
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				output = append(output, node.Val)
				predecessor.Right = node
				node = node.Left
			} else {
				predecessor.Right = nil
				node = node.Right
			}
		}
	}
	return output
}
```
## In OrderTraversal
```go
func inorderTraversal(root *TreeNode) []int {
	var response []int
	var traverseTree func(node *TreeNode)

	traverseTree = func(head *TreeNode) {
		if head == nil {
			return
		}

		traverseTree(head.Left)
		response = append(response, head.Val)
		traverseTree(head.Right)
	}
	traverseTree(root)

	return response
}

func inorderTraversalStack(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	curr := root
	for curr != nil || len(stack) != 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1] //! POP
		res = append(res, curr.Val)
		curr = curr.Right
	}
	return res
}

func inorderTraversalMorris(root *TreeNode) []int {
	res := []int{}
	curr := root
	var pre *TreeNode
	for curr != nil {
		if curr.Left == nil {
			res = append(res, curr.Val)
			curr = curr.Right // move to next right node
		} else { // has a left subtree
			pre = curr.Left
			for pre.Right != nil && pre.Right != curr { // find rightmost
				pre = pre.Right
			}
			if pre.Right == nil { // if the rightmost node's right child is nil
				pre.Right = curr // put current after the pre node
				curr = curr.Left // move curr to the top of the new tree
			} else { // the rightmost node's right child is not null
				pre.Right = nil             // restore the tree
				res = append(res, curr.Val) // add current value to the result list
				curr = curr.Right           // move current node to the right
			}
		}
	}
	return res
}

```

## PostOrder Traversal
```go
func inorderTraversal(root *TreeNode) []int {
	var response []int
	var traverseTree func(node *TreeNode)

	traverseTree = func(head *TreeNode) {
		if head == nil {
			return
		}

		traverseTree(head.Left)
		traverseTree(head.Right)
		response = append(response, head.Val)
	}
	traverseTree(root)

	return response
}

func postOrderTraversalStack(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	var previousNode *TreeNode

	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			top := stack[len(stack)-1]
			if top.Right == nil || top.Right == previousNode {
				result = append(result, top.Val)
				stack = stack[:len(stack)-1]
				previousNode = top
				root = nil
			} else {
				root = top.Right
			}
		}
	}

	return result
}

//! Morris postorderTraversal =======================

func postorderTraversalMorris(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	dummy := &TreeNode{Val: -1, Left: root}
	curr := dummy

	for curr != nil {
		if curr.Left != nil {
			pre := curr.Left

			// Buscar el predecesor
			for pre.Right != nil && pre.Right != curr {
				pre = pre.Right
			}

			if pre.Right == nil {
				// Crear el "thread"
				pre.Right = curr
				curr = curr.Left
			} else {
				// Reversar camino, agregar valores en reversa
				reverseSubtreeLinks(curr.Left, pre)
				node := pre
				for node != curr.Left {
					result = append(result, node.Val)
					node = node.Right
				}
				result = append(result, node.Val)

				// Restaurar árbol
				reverseSubtreeLinks(pre, curr.Left)
				pre.Right = nil
				curr = curr.Right
			}
		} else {
			curr = curr.Right
		}
	}

	return result
}

// Función auxiliar para invertir los enlaces del árbol temporalmente
func reverseSubtreeLinks(start, end *TreeNode) {
	if start == end {
		return
	}

	var prev *TreeNode
	curr := start

	for curr != end {
		next := curr.Right
		curr.Right = prev
		prev = curr
		curr = next
	}
	curr.Right = prev // último nodo
}

```
## Path sum to the leafs
```go
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	sum -= root.Val
	if root.Left == nil && root.Right == nil { // if reach a leaf
		return sum == 0
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
```