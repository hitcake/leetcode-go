package threeSum

import "sort"

const FinalTarget = 0

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	size := len(nums)
	for i := 0; i < size; i++ {
		if nums[i] > FinalTarget {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		result = twoSum(nums, i+1, size-1, FinalTarget-nums[i], nums[i], result)
	}
	return result
}

func twoSum(nums []int, start, end, target, preValue int, result [][]int) [][]int {
	for start < end {
		sum := nums[start] + nums[end]
		if sum == target {
			result = append(result, []int{preValue, nums[start], nums[end]})
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
