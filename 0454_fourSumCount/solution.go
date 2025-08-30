package fourSumCount

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0
	aMap := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			aMap[nums1[i]+nums2[j]]++
		}
	}
	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			c := aMap[0-(nums3[i]+nums4[j])]
			if c > 0 {
				count += c
			}
		}
	}
	return count
}
