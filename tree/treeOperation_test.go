package tree

import (
	"testing"
)

func TestTreeInsert(t *testing.T) {
	test := NewTree()
	test.Insert(50)
	test.Insert(30)
	test.Insert(45)
	test.Insert(42)
	test.Insert(46)
	test.Insert(25)
	test.Insert(70)
	test.Insert(65)
	test.Insert(75)
	inOrder(test.root)
}
