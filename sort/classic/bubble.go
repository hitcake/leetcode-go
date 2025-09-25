package classic

func BubbleSort(nums []int) {
	size := len(nums)
	for i := 0; i < size-1; i++ {
		var alreadyHasOrder = true
		for j := 0; j < size-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				alreadyHasOrder = false
			}
		}
		if alreadyHasOrder {
			break
		}
	}
}
