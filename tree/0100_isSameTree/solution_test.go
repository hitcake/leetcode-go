package leetcode0100

import (
	. "leetcode-go/tree"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	p := BuildTree([]interface{}{1, 2, 3, nil, 5, 6, 7})
	q := BuildTree([]interface{}{1, 2, 3, nil, 5, 6, 7})
	if !IsSameTree(p, q) {
		t.Errorf("IsSameTree failed")
	}
	p = BuildTree([]interface{}{1, 2, 3})
	q = BuildTree([]interface{}{1, 2})
	if IsSameTree(p, q) {
		t.Errorf("IsSameTree failed")
	}
}
