# Sliding window

Longest substring without repeating characters

```go
func lengthOfLongestSubstring(s string) int {
	chars := make(map[byte]int)

	left := 0
	right := 0

	res := 0
	for right < len(s) {
		rightCharacter := s[right]
		chars[rightCharacter]++

		for chars[rightCharacter] > 1 {
			l := s[left]
			chars[l]--
			left++
		}

		res = max(res, right-left+1)
		right++
	}

	return res
}
```
