package arrayProblem

import (
	"algoPractices/sorting"
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

func TestFind3MForGivenInput(*testing.T) {
	in := []int{64630, 11735, 14216, 99233, 14470, 4978, 73429, 38120, 51135, 67060}
	fmt.Println(Find3MForGivenInput(in))
}

func TestFindQuartilesGivenInput(*testing.T) {
	//in := []int{3, 7, 8, 5, 12, 14, 21, 13, 18}
	in := []int{10, 3, 7, 8, 5, 12, 14, 21, 15, 18, 14} //Output : 7 13 15
	sorting.QuickSort(in)
	q2, index1 := FindQuartilesGivenInput(in)
	q1, _ := FindQuartilesGivenInput(in[:index1])
	q3, _ := FindQuartilesGivenInput(in[index1+1:])
	fmt.Println(q1)
	fmt.Println(q2)
	fmt.Println(q3)
}

func TestFindSDForGivenArray(*testing.T) {
	in := []int{10, 40, 30, 50, 20}
	FindSDForGivenArray(in)
}
