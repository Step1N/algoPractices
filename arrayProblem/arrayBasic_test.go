package arrayProblem

import (
	"fmt"
	"testing"
)

func TestFindMinimumSellerPrice(t *testing.T) {
	//1 5 20
	//3 8 15
	//7 10 8
	//7 8 9
	/*var testCase int
	_, _ = fmt.Scan(&testCase)*/
	//takeInput(3)
}

func TestComputPrefixSum(t *testing.T) {
	s := []int{3, 3, 9, 9, 5}
	fmt.Print(ComputePrefixSum(s, 7))
}

func TestFindLongestIncreasingSubsequence(*testing.T) {
	in := []int{3, 10, 2, 1, 20}
	fmt.Println(findLISSequenceDP(in, len(in)))
}
func TestFindLCSInSequence(*testing.T) {
	s, s1 := "AGGTAB", "GXTXAYB"
	if len(s) > len(s1) {
		fmt.Println(findLCSInSequence(s, s1))
	} else {
		fmt.Println(findLCSInSequence(s1, s))

	}
}

func TestFindLCSInSequenceRec(*testing.T) {
	s, s1 := "AGGTAB", "GXTXAYB"
	fmt.Println(findLCSInSequenceRec(s1, s, len(s)-1, len(s1)-1))
}
func TestFindLCSInSequenceDP(*testing.T) {
	s, s1 := "AGGTAB", "GXTXAYB"
	fmt.Println(findLCSInSequenceDP(s1, s, len(s)-1, len(s1)-1))
}
