package lengthOfLongestSubstring

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var tests = []struct {
		s    string
		want int
	}{
		{"pwwkew", 3},
		{"bbbbb", 1},
		{"abcabcbb", 3},
	}
	for _, c := range tests {
		got := lengthOfLongestSubstring(c.s)
		if got != c.want {
			t.Errorf("func lengthOfLongestSubstring(%v) = %v, want %v", c.s, got, c.want)
		}
	}
}
