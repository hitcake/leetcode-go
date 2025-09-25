package classic

import (
	"reflect"
	"testing"
)

var tests = []struct {
	nums []int
	want []int
}{
	{[]int{2, 7, 1, 0, 3, 8, 5}, []int{0, 1, 2, 3, 5, 7, 8}},
	{[]int{9, 1, 13, 22, 5, 8}, []int{1, 5, 8, 9, 13, 22}},
}

func TestSort(t *testing.T) {
	sortAndCompare(t, BubbleSort)
	sortAndCompare(t, SelectionSort)
	sortAndCompare(t, InsertionSort)
	sortAndCompare(t, ShellSort)
	sortAndCompare(t, MergeSort)
}

func sortAndCompare(t *testing.T, f func([]int)) {
	for _, tt := range tests {
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)
		f(nums)
		if !reflect.DeepEqual(nums, tt.want) {
			//method := reflect.ValueOf(tt.want)
			// todo  如何打印方法的信息
			t.Errorf("got %v, want %v", nums, tt.want)
		}
	}
}
