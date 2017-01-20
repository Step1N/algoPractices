package dp

import (
	"fmt"
	"strconv"
	"strings"
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

func fibInitial(a, b, n int) string {
	lookup := make(map[int]string, n)
	lookup[0] = strconv.Itoa(a)
	lookup[1] = strconv.Itoa(b)
	for i := 2; i < 5; i++ {
		x, _ := strconv.Atoi(lookup[i-2])
		y, _ := strconv.Atoi(lookup[i-1])
		lookup[i] = strconv.Itoa(x + y*y)
	}

	for j := 5; j < n; j++ {
		fmt.Println("Lookup  ", j, lookup[j-1], lookup[j-2])
		sqr := multiplyTwoStringNumber(lookup[j-1], lookup[j-1])
		addition := addTwoStringNumber(lookup[j-2], sqr)
		lookup[j] = addition
	}
	return lookup[n-1]
}

func multiplyTwoStringNumber(a, b string) string {
	num1 := getIntArray(strings.Split(a, ""))
	num2 := getIntArray(strings.Split(b, ""))
	s1 := len(num1)
	s2 := len(num2)
	rsSize := (s1 + s2) + 1
	res := make([]int, rsSize)
	var carry int
	m := 0
	l := 0
	for i := s2 - 1; i >= 0; i-- {
		carry = 0
		l = rsSize - 1 - m
		for j := s1 - 1; j >= 0; j-- {
			t := num1[j]*num2[i] + carry
			carry = t / 10
			rm := t % 10
			if m > 0 {
				s := res[l] + rm
				carry += s / 10
				rm = s % 10
			}
			res[l] = rm
			l--
		}
		if carry > 0 {
			res[l] = carry
			l--
		}
		m++
	}
	res[l] = -1
	return getString(res)
}

func getString(s []int) string {
	res := ""
	if len(s) < 4 {
		res += strconv.Itoa(s[1])
		res += strconv.Itoa(s[2])
		return res
	}
	j := 1
	for i := 0; i < len(s); i++ {
		if s[i] != -1 {
			j++
			break
		}

	}
	for ; j < len(s); j++ {
		res += strconv.Itoa(s[j])
	}
	fmt.Println(res)
	return res
}

func addTwoStringNumber(a, b string) string {
	num1 := getIntArray(strings.Split(a, ""))
	num2 := getIntArray(strings.Split(b, ""))
	s1 := len(num1)
	s2 := len(num2)
	res := make([]int, s1+s2+2)
	var carry int
	if s1 <= s2 {
		j := 0
		k := s2 - 1
		for i := s1 - 1; i >= 0; i-- {
			t := num1[i] + num2[k]
			if carry > 0 {
				t += carry
			}
			res[j] = t % 10
			carry = t / 10
			j++
			k--
		}
		for m := k; m >= 0; m-- {
			if carry > 0 {
				t := num2[m] + carry
				res[j] = t % 10
				carry = t / 10
			} else {
				res[j] = num2[m]
			}
			j++
		}
		if carry > 0 {
			res[j] = carry
			res[j+1] = -1
		} else {
			res[j] = -1
		}

	} else if s1 > s2 {
		j := 0
		k := s1 - 1
		for i := s2 - 1; i >= 0; i-- {
			t := num1[k] + num2[i]
			if carry > 0 {
				t += carry
			}
			res[j] = t % 10
			carry = t / 10
			j++
			k--
		}
		for m := k; m >= 0; m-- {
			if carry > 0 {
				t := num1[m] + carry
				res[j] = t % 10
				carry = t / 10
			} else {
				res[j] = num1[m]
			}
			j++
		}
		if carry > 0 {
			res[j] = carry
			res[j+1] = -1
		} else {
			res[j] = -1
		}
	}

	return getStringFromIntArray(res)
}

func getStringFromIntArray(s []int) string {
	res := ""
	j := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == -1 {
			j = i - 1
		}
		if j >= 0 {
			res += strconv.Itoa(s[j])
			j--
		}
	}

	return res
}

func getIntArray(s []string) []int {
	size := len(s)
	res := make([]int, size)
	for i := size - 1; i >= 0; i-- {
		res[i], _ = strconv.Atoi(s[i])
	}

	return res
}
