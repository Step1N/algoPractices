package tree

import (
	"fmt"
)

// A Node is a node with left and right link and value
type Node struct {
	left  *Node
	value int
	right *Node
}

// A Tree is a binary tree with integer values.
type Tree struct {
	root *Node
	size int
}

// Create new tree
func NewTree() *Tree {
	tree := new(Tree)
	tree.size = 0

	return tree
}

func (tree *Tree) Insert(v int) {
	if tree.root == nil {
		tree.root = &Node{nil, v, nil}
	}
	tree.size++
	tree.root.insert(&Node{nil, v, nil})
}

func (node *Node) insert(t *Node) {
	if t.value > node.value {
		if node.right == nil {
			node.right = t
		} else {
			node.right.insert(t)
		}
	} else if t.value < node.value {
		if node.left == nil {
			node.left = t
		} else {
			node.left.insert(t)
		}
	}
}

func inOrder(t *Node) {
	if t == nil {
		return
	}
	inOrder(t.left)
	fmt.Printf("\t %v", t.value)
	inOrder(t.right)
}

func postOrder(t *Node) {
	if t == nil {
		return
	}
	postOrder(t.left)
	postOrder(t.right)
	fmt.Printf("\t %v", t.value)
}

func preOrder(t *Node) {
	if t == nil {
		return
	}
	fmt.Printf("\t %v", t.value)
	preOrder(t.left)
	preOrder(t.right)
}
