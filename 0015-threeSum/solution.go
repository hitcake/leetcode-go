package threeSum

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if k, ok := numMap[0-(nums[j]+nums[i])]; ok {
				if k != i && k != j {
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return result
}
