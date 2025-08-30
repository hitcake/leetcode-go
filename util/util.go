package util

import (
	"reflect"
	"sort"
)

type TwoDimArray [][]int

func NewTwoDimArray(nums [][]int) TwoDimArray {
	for i := 0; i < len(nums); i++ {
		sort.Ints(nums[i])
	}
	return nums

}

func TwoDimArrayEqual(a, b [][]int) bool {
	left := NewTwoDimArray(a)
	sort.Sort(left)
	right := NewTwoDimArray(b)
	sort.Sort(right)
	return reflect.DeepEqual(left, right)
}
func (a TwoDimArray) Len() int {
	return len(a)
}
func (a TwoDimArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a TwoDimArray) Less(i, j int) bool {
	m := 0
	n := 0
	for m < len(a[i]) || n < len(a[j]) {
		if m < len(a[i]) && n < len(a[j]) {
			if a[i][m] != a[j][n] {
				return a[i][m] <= a[j][n]
			} else {
				m++
				n++
			}
		} else {
			return len(a[i]) <= len(a[j])
		}

	}
	return false
}
