package leetcode0207

func canFinish(numCourses int, prerequisites [][]int) bool {
	state := make(map[int]int) // 0 未访问  1正在访问  2 访问完成
	dependency := make(map[int][]int)
	for i := 0; i < len(prerequisites); i++ {
		src := prerequisites[i][0]
		dest := prerequisites[i][1]
		dependency[src] = append(dependency[src], dest)
	}
	var hasRing = false
	var dfs func(course int)
	dfs = func(course int) {
		courseState := state[course]
		if courseState == 0 {
			state[course] = 1
			if destList, ok := dependency[course]; ok {
				for _, v := range destList {
					if hasRing {
						return
					}
					dfs(v)
				}
			}
			state[course] = 2
		} else if courseState == 1 {
			hasRing = true
			return
		}
	}
	for i := 0; i < numCourses; i++ {
		if hasRing {
			return false
		}
		dfs(i)
	}
	return true

}
