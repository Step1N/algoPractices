package graph

import (
	"fmt"
	"testing"
)

func TestBFSSearch(t *testing.T) {
	var nodes = map[int][]int{
		1:  []int{2, 3, 4},
		2:  []int{1, 5, 6},
		3:  []int{1},
		4:  []int{1, 7, 8},
		5:  []int{2, 9, 10},
		6:  []int{2},
		7:  []int{4, 11, 12},
		8:  []int{4},
		9:  []int{5},
		10: []int{5},
		11: []int{7},
		12: []int{7},
	}
	visited := []int{}
	bfsSearch(1, nodes, func(node int) {
		visited = append(visited, node)
	})
	fmt.Println(visited)
}

func TestDFSSearch(t *testing.T) {
	var nodes = map[int][]int{
		1:  []int{2, 7, 8},
		2:  []int{1, 3, 6},
		3:  []int{2, 4, 5},
		4:  []int{3},
		5:  []int{3},
		6:  []int{2},
		7:  []int{1},
		8:  []int{1, 9, 12},
		9:  []int{8, 10, 11},
		10: []int{9},
		11: []int{9},
		12: []int{8},
	}
	visited := []int{}
	dfsSearch(1, nodes, func(node int) {
		visited = append(visited, node)
	})
	fmt.Println(visited)
}

func TestSynchronizeShopping(t *testing.T) {

	visited := []int{}
	count := 0

	fmt.Println(count)
	fmt.Println(visited)
}

func TestFindAstronautPairMoonMission(t *testing.T) {
	var nodes = map[int][]int{
		0: []int{1},
		2: []int{3},
	}

	visited := []int{}
	FindAstronautPairMoonMission(0, nodes, func(node int) {
		visited = append(visited, node)
	})

	fmt.Println(visited)

	visited = []int{}
	FindAstronautPairMoonMission(2, nodes, func(node int) {
		visited = append(visited, node)
	})
	fmt.Println(visited)

	//Second test case
	var nodes1 = map[int][]int{
		0: []int{2},
		1: []int{4, 8},
		2: []int{0, 6, 8},
		3: []int{5},
		4: []int{1},
		5: []int{3},
		6: []int{9},
		7: []int{7},
		8: []int{1},
		9: []int{6},
	}

	visited = []int{}
	FindAstronautPairMoonMission(0, nodes1, func(node int) {
		visited = append(visited, node)
	})
	fmt.Println(visited)
}
