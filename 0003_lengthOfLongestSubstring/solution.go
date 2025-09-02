package lengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	var longest = 0
	var left, right, n = 0, 0, len(s)
	for left < n && right < n {
		if left > 0 {
			delete(m, s[left-1])
		}
		for right < n && m[s[right]] == 0 {
			m[s[right]]++
			right++
		}
		longest = max(longest, right-left)
		left++
	}
	return longest
}
