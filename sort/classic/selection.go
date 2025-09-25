package classic

func SelectionSort(nums []int) {
	size := len(nums)
	for i := 0; i < size; i++ {
		minValueIndex := i
		for j := i + 1; j < size; j++ {
			if nums[j] < nums[minValueIndex] {
				minValueIndex = j
			}
		}
		if minValueIndex != i {
			nums[i], nums[minValueIndex] = nums[minValueIndex], nums[i]
		}
	}
}
