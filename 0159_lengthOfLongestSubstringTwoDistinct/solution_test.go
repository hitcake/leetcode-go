package lengthOfLongestSubstringTwoDistinct

import "testing"

func TestLengthOfLongestSubstringTwoDistinct(t *testing.T) {
	var tests = []struct {
		s    string
		want int
	}{
		{"abcabcbb", 4},
		{"bbbbb", 5},
		{"pwwkew", 3},
		{"a", 1},
		{"", 0},
	}
	for _, tt := range tests {
		got := lengthOfLongestSubstringTwoDistinct(tt.s)
		if got != tt.want {
			t.Errorf("lengthOfLongestSubstringTwoDistinct(%s) got: %d, want: %d", tt.s, got, tt.want)
		}
	}
}
