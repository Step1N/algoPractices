package graph

func bfsSearch(start int, nodes map[int][]int, visitFunc func(int)) {
	frontier := []int{start}
	visited := map[int]bool{}
	next := []int{}

	for 0 < len(frontier) {
		next = []int{}
		for _, node := range frontier {
			visited[node] = true
			visitFunc(node)
			for _, n := range bfsHelper(node, nodes, visited) {
				next = append(next, n)
			}
		}
		frontier = next
	}
}

func bfsHelper(node int, nodes map[int][]int, visited map[int]bool) []int {
	next := []int{}
	isVisited := func(n int) bool { _, ok := visited[n]; return !ok }
	for _, n := range nodes[node] {
		if isVisited(n) {
			next = append(next, n)
		}
	}
	return next
}

func dfsSearch(node int, nodes map[int][]int, visitFunc func(int)) {
	dfsVisit(node, nodes, map[int]bool{}, visitFunc)
}

func dfsVisit(node int, nodes map[int][]int, v map[int]bool, visitFunc func(int)) {
	v[node] = true
	visitFunc(node)
	for _, n := range nodes[node] {
		if _, ok := v[n]; !ok {
			dfsVisit(n, nodes, v, visitFunc)
		}
	}
}
