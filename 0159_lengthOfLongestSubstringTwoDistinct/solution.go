package lengthOfLongestSubstringTwoDistinct

func lengthOfLongestSubstringTwoDistinct(s string) int {
	m := map[byte]int{}
	var longest, n, left, right = 0, len(s), 0, 0
	for left <= right && right < n {
		if left > 0 {
			delete(m, s[left-1])
		}
		for right < n {
			// 新字符不在map中
			if _, ok := m[s[right]]; !ok {
				// 当前最多一个字符
				if len(m) <= 1 {
					m[s[right]]++
					right++
				} else {
					break
				}
			} else {
				m[s[right]]++
				right++
			}
		}
		longest = max(longest, right-left)
		left++
	}
	return longest
}
