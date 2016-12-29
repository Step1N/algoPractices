package graph

/*
type Graph struct {
	nodes map[int]Node
}

type Node struct {
	payload  interface{}
	adjNodes map[int]struct{}
}

func (g *Graph) addNode(key int, payload interface{}) Node{
	if node, found := g.getNode(key); !found {
		node = Node{payload: payload, adjNodes: make(map[int]struct{})}
		g.nodes[key] = node
	}
	return node
}

func (g *Graph) getNode(key int) Node{
	if key < 0 {
		panic("Node key can't be negative")
	}
	if node, found := g.nodes[key]; !found {
		panic("Node key not found")
	}
	return node
}

func (g *Graph) addEdge(src, dst int) {
	if srcNode, found := g.getNode(src); !found {
		panic("No node " + string(src))
	}
	if dstNode, found := g.getNode(dst); !found {
		panic("No node " + string(dst))
	}
	// connect by adjacent list
	srcNode.adjNodes[dst] = struct{}{}
	dstNode.adjNodes[src] = struct{}{}
}*/
