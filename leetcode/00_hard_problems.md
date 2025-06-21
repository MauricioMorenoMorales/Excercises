
# Combinatorics

```go
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

```
# Creating trees
```go
func buildTree(inorder []int, postorder []int) *TreeNode {
	idx_map := map[int]int{}
	post_idx := len(postorder) - 1
	var helper func(in_left int, in_right int) *TreeNode
	helper = func(in_left int, in_right int) *TreeNode {
		if in_left > in_right {
			return nil
		}
		root_val := postorder[post_idx]
		root := &TreeNode{Val: root_val}
		index := idx_map[root_val]
		post_idx--
		root.Right = helper(index+1, in_right)
		root.Left = helper(in_left, index-1)
		return root
	}
	for i := 0; i < len(inorder); i++ {
		idx_map[inorder[i]] = i
	}
	return helper(0, len(inorder)-1)
}
```