package dp

import (
	"fmt"
	"testing"
)

func TestSherlockCost(t *testing.T) {
	//in := []int{10, 1, 10, 1, 10, 1} //Output 36
	in := []int{100, 2, 100, 2, 100} //Output 396
	findBestCombination(in)
	res := sherlockSum(in, 0, 0, len(in)-1, 1)
	fmt.Println(res)
}
