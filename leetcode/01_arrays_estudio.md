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