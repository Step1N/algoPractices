package graph

import (
	"fmt"
	"testing"
)

func TestGraphCreation(t *testing.T) {
	s := `
		S|A,15|B,14|C,9
		A|S,15|B,5|D,20|T,44
		B|S,14|A,5|D,30|E,18
		C|S,9|E,24
		D|A,20|B,30|E,2|F,11|T,16
		E|B,18|C,24|D,2|F,6|T,19
		F|D,11|E,6|T,6
		T|A,44|D,16|F,6|E,19
	`
	G2 := plotGraph(s)
	DepthFirstSearch(G2.GetVertexByID("S"))
	fmt.Println()
}

func TestPathExisBetweenTwoVertex(t *testing.T) {
	s := `
		0|1,10|2,10
		1|2,10
		2|0,10
		2|3,10
		3|3,10
	`
	G2 := plotGraph(s)
	BreadthFirstSearch(G2.GetVertexByID("0"))
	fmt.Println()
}

func TestConnectedGraph(t *testing.T) {
	s := `
		0|1,10|3,10
		1|2,10
		2|3,10|4,10
		3|0,10
		4|4,10
	`
	G2 := plotGraph(s)
	BreadthFirstSearch(G2.GetVertexByID("0"))
	fmt.Println()
}

func TestDisConnectedGraph(t *testing.T) {
	s := `
		0|1,10
		1|2,10
		2|3,10
		3
	`
	G2 := plotGraph(s)
	DepthFirstSearch(G2.GetVertexByID("0"))
	fmt.Println()
}

func TestSynchronizeShopping(t *testing.T) {
	s := `
		1|2,10|3,10
		2|1,10|4,10
		3|1,10|5,10
		4|2,10|5,10
		5|3,10|4,10
	`
	G2 := plotGraph(s)
	tmap := make(map[int][]int, 5)
	i := 1
	total := NoWayToReach(G2.GetVertexByID("1"), G2.GetVertexByID("5"), tmap, i)
	fmt.Println("total ", total)
	fmt.Println(tmap)
}

func TestCountingOnTree(t *testing.T) {
	s := `
		1|2,1|3,1
		2|1,1|10,1
		3|1,1|4,1|5,1|6,1
		4|3,1|7,1
		5|3,1|8,1
		6|3,1
		7|4,1|9,1
		8|5,1
		9|7,1
		10|2,1
	`
	G2 := plotGraph(s)
	CountingOnTreeProblem(G2.GetVertexByID("1"))
	fmt.Println()
}
