package dp

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"log"
	"time"
	"fmt"
	"math/big"
)

func TestFibnacciSeries(t *testing.T) {
	//Normal approach
	start := time.Now()
	res := fib(9)
	elapsed := time.Since(start)
	log.Printf("Time took normal %s \n", elapsed)
	assert.Equal(t, 34, res)

	//Topdown approach
	start = time.Now()
	lookup := make(map[int]int, 0)
	res = fibTopBottom(9, lookup)
	elapsed = time.Since(start)
	log.Printf("Time took for top down %s \n", elapsed)
	assert.Equal(t, 34, res)

	//Bottom up approach
	start = time.Now()
	res = fibBottomUp(9)
	elapsed = time.Since(start)
	log.Printf("Time took for bottom up %s \n", elapsed)
	assert.Equal(t, 34, res)
}

func TestLongestCommonString(t *testing.T) {
	s1 := strings.Split("AGGTAB", "")
	s2 := strings.Split("GXTXAYB", "")
	res := longestCommonStringdp(s1, s2, len(s1), len(s2))
	assert.Equal(t, 4, res)
}

func TestCountMinNumberOfCoin(t *testing.T) {
	arr := []int{1, 2, 3}
	res := countdp(arr, len(arr), 4)
	assert.Equal(t, 4, res)
}

func TestFibModified(t *testing.T) {
	a, b, n := 0, 1 ,10
	fmt.Print(fibModified(big.NewInt(int64(a)), big.NewInt(int64(b)), n))
}
