package graph

import (
	"fmt"
	"strconv"
	"strings"
)

type Edge struct {
	Source      *Vertex
	Destination *Vertex
	Weight      int
}

type Vertex struct {
	Id      string
	Visited bool
	AdjEdge []*Edge
}

type Graph struct {
	VertexArray []*Vertex
}

func NewGraph() *Graph {
	return &Graph{
		make([]*Vertex, 0),
	}
}

func NewVertex(s string) *Vertex {
	return &Vertex{
		Id:      s,
		Visited: false,
		AdjEdge: make([]*Edge, 0),
	}
}

func NewEdge(s, d *Vertex, w int) *Edge {
	return &Edge{
		s,
		d,
		w,
	}
}

func (G *Graph) GetVertexArray() []*Vertex {
	return G.VertexArray
}

func (G *Graph) AddVertices(vertices ...*Vertex) {
	for _, vertex := range vertices {
		G.VertexArray = append(G.VertexArray, vertex)
	}
}

func (A *Vertex) AddEdges(edges ...*Edge) {
	for _, edge := range edges {
		A.AdjEdge = append(A.AdjEdge, edge)
	}
}

func (G *Graph) GetVertexByID(id string) *Vertex {
	for _, vertex := range G.VertexArray {
		if vertex.Id == id {
			return vertex
		}
	}
	return nil
}

func (G *Graph) GetVertex(id string) *Vertex {
	vertex := G.GetVertexByID(id)
	if vertex == nil {
		vertex = NewVertex(id)
		G.AddVertices(vertex)
	}
	return vertex
}

func (A *Vertex) GetAdjacentEdge() chan *Edge {
	edgechan := make(chan *Edge)

	go func() {
		defer close(edgechan)
		for _, edge := range A.AdjEdge {
			edgechan <- edge
		}
	}()

	return edgechan
}

func StrToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		fmt.Errorf("failed to convert string, %v", err)
	}
	return result
}

func plotGraph(s string) *Graph {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	newGraph := NewGraph()

	for _, line := range lines {
		edgeValues := strings.Split(line, "|")
		sID := edgeValues[0]
		otherEdges := edgeValues[1:]
		newGraph.GetVertex(sID)
		for _, oEdge := range otherEdges {
			if len(strings.Split(oEdge, ",")) == 1 {
				continue
			}
			dID := strings.Split(oEdge, ",")[0]
			w := StrToInt(strings.Split(oEdge, ",")[1])

			sVtx := newGraph.GetVertex(sID)
			dVtx := newGraph.GetVertex(dID)
			e := NewEdge(sVtx, dVtx, w)
			sVtx.AddEdges(e)

			e = NewEdge(dVtx, sVtx, w)
			dVtx.AddEdges(e)
		}
	}

	return newGraph
}

func DepthFirstSearch(s *Vertex) {
	if s.Visited {
		return
	}

	s.Visited = true
	fmt.Printf("%v ", s.Id)
	for edge := range s.GetAdjacentEdge() {
		DepthFirstSearch(edge.Destination)
	}
}

func BreadthFirstSearch(s *Vertex) {
	front := []*Vertex{s}
	s.Visited = true
	fmt.Printf("%v \n", s.Id)
	for i := 0; i < len(front); i++ {
		u := front[i]
		for _, edge := range u.AdjEdge {
			d := edge.Destination
			if !d.Visited {
				fmt.Printf("%v ", d.Id)
				d.Visited = true
				front = append(front, d)
			}
		}
	}

}

const MAXWEIGHT = 1000000

type minDistanceFromVertex map[*Vertex]int

var Predecessor []string

func (G *Graph) MinSpanning(StartSource, TargetSource *Vertex) {
	D := make(minDistanceFromVertex)
	for _, vertex := range G.VertexArray {
		D[vertex] = MAXWEIGHT
	}
	D[StartSource] = 0

	for edge := range StartSource.GetAdjacentEdge() {
		D[edge.Destination] = edge.Weight
	}
	findDistance(StartSource, D)
	for k, v := range D {
		fmt.Println(k.Id, v)
	}
}

func findDistance(StartSource *Vertex, D minDistanceFromVertex) {
	for edge := range StartSource.GetAdjacentEdge() {
		if D[edge.Destination] > D[edge.Source]+edge.Weight {
			D[edge.Destination] = D[edge.Source] + edge.Weight
			Predecessor = append(Predecessor, edge.Source.Id+"->"+edge.Destination.Id)
		} else if D[edge.Destination] < D[edge.Source]+edge.Weight {
			continue
		}
		findDistance(edge.Destination, D)
	}
}

func NoWayToReach(startSource, destination *Vertex, routWeights map[int][]int, i int) int {
	if startSource.Id == destination.Id {
		return 1
	}
	total := 0
	startSource.Visited = true
	var rout []int
	if len(routWeights[i]) != 0 {
		rout = routWeights[i]
	}
	for edge := range startSource.GetAdjacentEdge() {
		if !edge.Destination.Visited {

			rout = append(rout, edge.Weight)
			total += NoWayToReach(edge.Destination, destination, routWeights, i+1)
		}
	}
	routWeights[i] = rout

	return total
}

func CountingOnTreeProblem(startSource, destination *Vertex) {
	front := []*Vertex{startSource}
	startSource.Visited = true
	fmt.Printf("%v \n", startSource.Id)
	isDestinationArrived := false
	for i := 0; i < len(front); i++ {
		u := front[i]
		for _, edge := range u.AdjEdge {
			d := edge.Destination
			if !d.Visited {
				fmt.Printf("%v ", d.Id)
				d.Visited = true
				front = append(front, d)
			}
			if startSource.Id == destination.Id {
				isDestinationArrived = true
				break
			}
		}

		if isDestinationArrived {
			break
		}
	}
}

func DijkstraMinimumDistanceFromSourceToDestination(startSource, destination *Vertex) {

}
