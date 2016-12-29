package disjoint

import (
	"fmt"
	"testing"
)

func TestDisjointSet(t *testing.T) {
	/*
			TestCaseOne
			1000 1
			1, 2
			output 499499
			TestCaseTwo
			4 2
			0, 1
			2, 3
			output 4
			TestCase Three

			10 7
			0 2
			1 8
			1 4
			2 8
			2 6
			3 5
			6 9
		TestCase Four
		5000 1
		1893 3230
		output 12497499
	*/
	n, s := 10, 7
	fmt.Println(s)
	jtmap := newDisjointSet(n)

	unionSet(0, 2, jtmap)
	unionSet(1, 8, jtmap)
	unionSet(1, 4, jtmap)
	unionSet(2, 8, jtmap)
	unionSet(2, 6, jtmap)
	unionSet(3, 5, jtmap)
	unionSet(6, 9, jtmap)
	var r int
	for i := 0; i < n; i++ {
		p := find(i, jtmap)
		fmt.Println("Found parent : ", p.data)
		if i == p.data {
			r += (p.rank * (n - p.rank))
			fmt.Println(i, "Value is : ", p.rank * (n - p.rank))
		}

	}
	fmt.Println(r / 2)
}
