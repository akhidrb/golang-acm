package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	for _, pre := range prerequisites {
		from, to := pre[1], pre[0]
		graph[from] = append(graph[from], to)
		indegree[to]++
	}

	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	num := 0
	for len(queue) > 0 {
		from := queue[0]
		queue = queue[1:]
		num++
		for _, to := range graph[from] {
			indegree[to]--
			if indegree[to] == 0 {
				queue = append(queue, to)
			}
		}
	}
	return num == numCourses
}
