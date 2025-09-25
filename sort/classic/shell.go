package classic

func ShellSort(nums []int) {
	size := len(nums)
	gap := size / 2
	for gap > 0 {
		for i := gap; i < size; i++ {
			j := i
			temp := nums[j]
			for j >= gap && nums[j-gap] > temp {
				nums[j] = nums[j-gap]
				j = j - gap
			}
			nums[j] = temp
		}
		gap = gap / 2
	}
}
