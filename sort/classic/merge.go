package classic

func MergeSort(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, low, high int) {
	if low < high {
		mid := (low + high) / 2
		mergeSort(nums, low, mid)
		mergeSort(nums, mid+1, high)
		merge(nums, low, mid, high)
	}
}
func merge(nums []int, low, mid, high int) {
	temp := make([]int, high-low+1)
	k, i, j := 0, low, mid+1
	for i <= mid && j <= high {
		if nums[i] <= nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			temp[k] = nums[j]
			j++
		}
		k++
	}
	for i <= mid {
		temp[k] = nums[i]
		i++
		k++
	}
	for j <= high {
		temp[k] = nums[j]
		j++
		k++
	}
	for k = 0; k < len(temp); k++ {
		nums[low+k] = temp[k]
	}
}
