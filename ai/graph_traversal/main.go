package main

type Graph struct {
	edges map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		edges: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(from, to int) {
	g.edges[from] = append(g.edges[from], to)
}

func (g *Graph) ReachesEnd(start, last int) bool {
	visited := make(map[int]bool)
	queue := []int{start}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range g.edges[node] {
			if neighbor == last {
				return true
			}
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return false
}

func (g *Graph) HasCycle() bool {
	visited := make(map[int]bool)
	visiting := make(map[int]bool)

	for node := range g.edges {
		if g.hasCycleUtil(node, visited, visiting) {
			return true
		}
	}
	return false
}

func (g *Graph) hasCycleUtil(node int, visited, visiting map[int]bool) bool {
	if visiting[node] {
		return true
	}

	if visited[node] {
		return false
	}

	visited[node] = true
	visiting[node] = true

	for _, neighbor := range g.edges[node] {
		if g.hasCycleUtil(neighbor, visited, visiting) {
			return true
		}
	}

	visiting[node] = false
	return false
}
