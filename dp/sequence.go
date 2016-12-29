package dp

import (
	"math/big"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fibTopBottom(n int, lookup map[int]int) int {
	if lookup[n] == 0 {

		if n <= 1 {
			lookup[n] = n
		} else {
			lookup[n] = fibTopBottom(n-1, lookup) + fibTopBottom(n-2, lookup)
		}
	}
	return lookup[n]
}

func fibBottomUp(n int) int {
	lookup := make(map[int]int, n+1)
	lookup[0], lookup[1] = 0, 1
	for i := 2; i <= n; i++ {
		lookup[i] = lookup[i-1] + lookup[i-2]
	}

	return lookup[n]
}

func longestCommonString(s1, s2 []string, m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if s1[m-1] == s2[n-1] {
		return 1 + longestCommonString(s1, s2, m-1, n-1)
	} else {
		return max(longestCommonString(s1, s2, m, n-1), longestCommonString(s1, s2, m-1, n))
	}
}
func longestCommonStringdp(s1, s2 []string, m, n int) int {
	L := make([][]int, m+1)
	for i := 0; i < m; i++ {
		L[i] = make([]int, n+1)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				L[i][j] = 0
			} else if s1[m-1] == s2[n-1] {
				L[i][j] = L[i-1][j-1] + 1
			} else {
				L[i][j] = max(L[i-1][j], L[i][j-1])
			}
		}
	}
	return L[m-1][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func coinChange(s []int, m, n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	if m <= 0 && n >= 1 {
		return 0
	}
	return coinChange(s, m-1, n) + coinChange(s, m, n-s[m-1])
}


func countdp(s []int, m, n int) int {
	table := make([]int, n+1)
	table[0] = 1
	for i := 0; i < m; i++ {
		for j := s[i]; j <= n; j++ {
			table[j] += table[j-s[i]]
		}
	}

	return table[n]
}

func fibModified(a, b *big.Int, n int) *big.Int {
	lookup := make(map[int]*big.Int)
	lookup[0] = a
	lookup[1] = b
	for i := 2; i < n; i++ {
		lookup[i] = lookup[i-2].Add(lookup[i-2], (lookup[i-1].Mul(lookup[i-1], lookup[i-1])))
	}
	return lookup[n-1]
}
