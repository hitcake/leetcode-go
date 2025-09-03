package minSubArrayLen

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	n := len(nums)
	total := 0
	shortest := n + 1
	for left < n {
		if left > 0 {
			total -= nums[left-1]
		}
		for right < n {
			if total+nums[right] < target {
				total += nums[right]
				right++
			} else {
				shortest = min(shortest, right-left+1)
				break
			}
		}
		left++
	}
	if shortest > n {
		shortest = 0
	}
	return shortest
}
