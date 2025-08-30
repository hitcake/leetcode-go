package fourSum

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	size := len(nums)
	for i := 0; i < size; i++ {
		if i+3 > size {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		result = threeSum(nums, i+1, size-1, target-nums[i], nums[i], result)
	}
	return result
}

func threeSum(nums []int, start, end, target, preValue int, result [][]int) [][]int {
	for start < end {
		result = twoSum(nums, start+1, end, target-nums[start], nums[start], preValue, result)
		for nums[start] == nums[start+1] && start < end-1 {
			start++
		}
		start++
	}
	return result
}

func twoSum(nums []int, start, end, target, preValue, prePreValue int, result [][]int) [][]int {
	for start < end {
		sum := nums[start] + nums[end]
		if sum == target {
			result = append(result, []int{prePreValue, preValue, nums[start], nums[end]})
			for start < end && nums[start] == nums[start+1] {
				start++
			}
			start++
			for start < end && nums[end] == nums[end-1] {
				end--
			}
			end--
		} else if sum < target {
			start++
		} else {
			end--
		}
	}
	return result
}
