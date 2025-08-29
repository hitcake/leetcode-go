package twosum

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := numMap[target-nums[i]]; ok {
			return []int{numMap[target-nums[i]], i}
		}
		numMap[nums[i]] = i
	}
	return []int{}
}
