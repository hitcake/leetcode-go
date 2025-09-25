package classic

func InsertionSort(nums []int) {
	size := len(nums)
	for i := 1; i < size; i++ {
		temp := nums[i]
		j := i
		for j > 0 && temp < nums[j-1] {
			nums[j] = nums[j-1]
			j--
		}
		nums[j] = temp
	}
}
