package disjoint

import (
	"fmt"
)

type Node struct {
	data   int
	parent *Node
	rank   int
}

func disjointSet(data int) *Node {
	node := &Node{}
	node.data = data
	node.parent = node
	node.rank = 0

	return node
}

func newDisjointSet(n int) map[int]*Node {
	jmap := make(map[int]*Node, 0)
	for i := 0; i < n; i++ {
		n := disjointSet(i)
		n.rank = 1
		jmap[i] = n
	}
	return jmap
}

func unionSet(value, value1 int, jmap map[int]*Node) {
	p1 := findParent(jmap[value])
	p2 := findParent(jmap[value1])
	if p1.data == p2.data {
		return
	}
	p1.parent = p2
	p2.rank += p1.rank
	fmt.Println(p1.parent.data, "Become parent of ", p1.data, " after union rank of ", p2.data, "become ", p2.rank)
}

func findParent(node *Node) *Node {
	parent := node.parent
	if parent == node {
		return parent
	}
	node.parent = findParent(node.parent)
	return node.parent
}

func find(in int, jmap map[int]*Node) *Node {

	return findParent(jmap[in])
}

func FindAstronautsForMoonMission(n, s int) int {
	jtmap := newDisjointSet(n)
	for i := 0; i < s; i++ {
		var p, q int
		fmt.Scan(&p, &q)
		unionSet(p, q, jtmap)
	}
	var r int
	for i := 0; i < n; i++ {
		p := find(i, jtmap)
		if i == p.data {
			r += (p.rank * (n - p.rank))
		}
	}

	return r / 2
}
