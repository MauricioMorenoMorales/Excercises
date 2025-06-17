# Minimum Size Subarray Sum

Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.

Input: target = 4, nums = [1,4,4]
Output: 1

Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0

```go
func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0

	var minLength int
	var sum int

	for right < len(nums) {
		sum += nums[right]

		if sum >= target {
			if minLength == 0 {
				minLength = right - left + 1
			} else {
				minLength = min(minLength, right-left+1)
			}

			sum -= nums[left]

			sum -= nums[right]

			left++
		} else {
			right++
		}
	}

	return minLength
}

```

# TwoPointer replacing elements technique

Remueve los elementos que aparecen mas de dos veces, dejando solo dos elementos
```go
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 1
	j := 1
	count := 1

	for i < len(nums) {
		if nums[i] == nums[i-1] {
			count++
			if count > 2 {
				i++
				continue
			}
		} else {
			count = 1
		}
		nums[j] = nums[i]
		j++
		i++
	}

	return j
}

```

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
