package leetcode0207

import (
	"testing"
)

func TestFindOrder(t *testing.T) {
	var tests = []struct {
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{2, [][]int{{1, 0}, {0, 1}}, false},
		{2, [][]int{{1, 0}}, true},
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, true},
		{1, [][]int{}, true},
		{5, [][]int{{1, 2}, {2, 3}, {3, 4}}, true},
		{3, [][]int{}, true},
	}
	for _, testCase := range tests {
		finish := canFinish(testCase.numCourses, testCase.prerequisites)
		if finish != testCase.want {
			t.Fatalf("numCourses %d prerequisites %v expected: %v, got: %v", testCase.numCourses, testCase.prerequisites, testCase.want, finish)
		}
	}

}
