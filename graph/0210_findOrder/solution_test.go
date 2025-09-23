package leetcode0210

import (
	"reflect"
	"testing"
)

func TestFindOrder(t *testing.T) {
	var tests = []struct {
		numCourses    int
		prerequisites [][]int
		expected      [][]int
	}{
		{2, [][]int{{1, 0}, {0, 1}}, [][]int{{}}},
		{2, [][]int{{1, 0}}, [][]int{{0, 1}}},
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, [][]int{{0, 1, 2, 3}, {0, 2, 1, 3}}},
		{1, [][]int{}, [][]int{[]int{0}}},
		{5, [][]int{{1, 2}, {2, 3}, {3, 4}}, [][]int{{4, 3, 2, 1, 0}, {4, 3, 2, 0, 1}, {4, 3, 0, 2, 1}, {4, 0, 3, 2, 1}, {0, 4, 3, 2, 1}}},
		{3, [][]int{}, [][]int{[]int{0, 1, 2}, {0, 2, 1}, {1, 2, 0}, {1, 0, 2}, {2, 1, 0}, {2, 0, 1}}},
	}
	for _, testCase := range tests {
		match := false
		order := findOrder(testCase.numCourses, testCase.prerequisites)
		for _, o := range testCase.expected {
			if reflect.DeepEqual(order, o) {
				match = true
				break
			}
		}
		if !match {
			t.Error("expected", testCase.expected, "got", order)
		}
	}

}
