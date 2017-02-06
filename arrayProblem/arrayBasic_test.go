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

func TestCountSubArrayHavingSumDivisibalByK(*testing.T) {
	in := []int{4, 5, 0, -2, -3, 1}
	fmt.Println(CountSubArrayHavingSumDivisibleByK(in, 5))
}

func TestCountMaxSubArray(*testing.T) {
	in := []int{4, 5, 0, -2, -3}
	fmt.Println(CountMaxSumOfSubSubArray(in))
}

func TestFindNumberOfSubArrayWhichMaxElementGreatorThanGiven(*testing.T) {
	in := []int{1, 2, 3}
	fmt.Println(FindNumberOfSubArrayWhichMaxElementGreaterThanGiven(in, 2))
}

func TestFindMaxIntegerOccurredInRange(*testing.T) {
	st := []int{1, 4, 9, 13, 21}
	en := []int{15, 8, 12, 20, 30}
	fmt.Println(FindMaxIntegerOccurredInRange(st, en))
}
