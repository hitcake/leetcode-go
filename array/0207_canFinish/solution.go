package leetcode0207

func canFinish(numCourses int, prerequisites [][]int) bool {
	dependency := make(map[int][]int)
	for i := 0; i < len(prerequisites); i++ {
		if list, ok := dependency[prerequisites[i][0]]; !ok {
			list = make([]int, 0)
			list = append(list, prerequisites[i][1])
			dependency[prerequisites[i][0]] = list
		} else {
			list = append(list, prerequisites[i][1])
			dependency[prerequisites[i][0]] = list
		}
	}

}
