package stringmodule

import (
	"fmt"
	"strings"
)

func reduceStringToUniqe(input string, index int) string {

	if len(strings.TrimSpace(input)) == 0 {
		return "Empty String"
	}
	if len(input) == 1 {
		return input
	}

	tmp := strings.Split(input, "")
	if index < len(input) && tmp[index] == tmp[index+1] {
		out := fmt.Sprintf("%s%s ", input[:index], input[index+2:])
		input = reduceStringToUniqe(out, index+1)
	} else {
		fmt.Println(input)
		input = reduceStringToUniqe(input, index+1)
	}

	return input
}

func countWord(s string) (count int) {
	sArray := strings.Split(s, "")
	for i := 0; i < len(sArray); i++ {
		if sArray[i] == strings.ToUpper(sArray[i]) {
			count++
		}
	}

	return count + 1
}

func constructStringWithTwoChar(s string) int {
	sArray := strings.Split(s, "")
	size := len(sArray)
	if size == 1 {
		return size
	}
	if size == 2 && sArray[0] != sArray[1] {
		return size
	}
	if size == 3 && sArray[0] == sArray[2] {
		return size
	}
	_ = findOccuranceOfChar(s)

	return 0

}

func findOccuranceOfChar(s string) map[string]int {
	occuranceMap := make(map[string]int, 0)
	sArray := strings.Split(s, "")
	for i := 0; i < len(sArray); i++ {
		if occuranceMap[sArray[i]] == 0 {
			occuranceMap[sArray[i]] = 1
		} else {
			occuranceMap[sArray[i]] = occuranceMap[sArray[i]] + 1
		}
	}
	return occuranceMap
}

func panagram(s string) string {
	ans := "not pangram"
	sArry := strings.Split(s, "")
	occurMap := make(map[string]int, 0)
	for i := 0; i < len(sArry); i++ {
		if strings.ToLower(sArry[i]) != " " {
			occurMap[strings.ToLower(sArry[i])] = occurMap[strings.ToLower(sArry[i])] + 1
		}
	}
	if len(occurMap) == 26 {
		ans = "pangram"
	}
	return ans
}

func funnyString(s string) string {
	ans := "Funny"
	rS := reverseStr(s)
	for i := 1; i < len(s); i++ {
		if findASCCIDiff(s[i], s[i-1]) != findASCCIDiff(rS[i], rS[i-1]) {
			return "Not Funny"
		}
	}

	return ans
}
func findASCCIDiff(a, b byte) byte {
	if a > b {
		return a - b
	} else {
		return b - a
	}

}

func reverseStr(s string) (rev string) {
	for _, r := range s {
		rev = string(r) + rev
	}
	return
}

func ShiftCharByGivenBit(shift byte, str string) string {
	shift = shift % 26
	result := ""
	var ind byte
	for i := 0; i < len(str); i++ {
		if byte(97) <= str[i] && str[i] <= byte(122) {
			ind = str[i] + shift
			if ind > byte(122) {
				diff := ind - byte(122)
				ind = byte(96) + diff
			}
			fmt.Println(ind)
			fmt.Println(string(ind))
			result += string(ind)

		} else if byte(65) <= str[i] && str[i] <= byte(90) {
			ind = str[i] + shift
			if ind > byte(90) {
				diff := ind - byte(90)
				ind = byte(64) + diff
			}
			result += string(ind)

		} else {
			result += string(str[i])
		}

	}

	return result
}

func BeautifulString(s string) string {
	sa := strings.SplitN(s, "010", 1)
	return strings.Join(sa, "")
}

func FindValidString(s string) string {
	res := "YES"
	smap := findMap(s)
	countMap := findMapOfOccurence(smap)
	fmt.Println(countMap)
	vals := getMapValues(countMap)
	//1. If all are same
	if isAllElementSame(vals) {
		return res
	} else if dif(vals[0], vals[1]) <= 1 {
		return res
	} else {
		vr := isUniquerElementInArray(vals)
		fmt.Println(vr)
	}

	return res
}

func isUniquerElementInArray(input []int) bool {
	m := make(map[int]int, len(input))
	for i := 1; i < len(input); i++ {
		if m[i] != 0 {
			t := m[i]
			t = t + 1
			m[i] = t
		} else {
			m[i] = 1
		}
	}
	if len(m) == 2 {
		fmt.Println(m)
	}

	return true
}

func isAllElementSame(input []int) bool {
	for i := 1; i < len(input); i++ {
		if input[0] != input[i] {
			return false
		}
	}

	return true
}

func getMapValues(s map[int]int) []int {
	values := make([]int, len(s)/2-1)
	for value, _ := range s {
		values = append(values, value)
	}

	return values
}

func dif(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func findMapOfMap(s map[int]int) map[int]int {
	m := make(map[int]int, len(s))
	for _, key := range s {
		if m[key] != 0 {
			t := m[key]
			t = t + 1
			m[key] = t
		} else {
			m[key] = 1
		}
	}
	return m
}

func findMapOfOccurence(s map[string]int) map[int]int {
	m := make(map[int]int, len(s))
	for _, key := range s {
		if m[key] != 0 {
			t := m[key]
			t = t + 1
			m[key] = t
		} else {
			m[key] = 1
		}
	}
	return m
}

func findMap(s string) map[string]int {
	m := make(map[string]int, len(s))
	for i := 0; i < len(s); i++ {
		if m[string(s[i])] != 0 {
			t := m[string(s[i])]
			t = t + 1
			m[string(s[i])] = t
		} else {
			m[string(s[i])] = 1
		}
	}
	return m
}

func maxSubArray(s string, index int) {
	if index > len(s) {
		return
	}
	fmt.Println(index, "  index ", s[index:])
	nt := index + 1
	maxSubArray(s, nt)
}

func powerSetOfArray(input []int, lookup map[uint]string, power uint) {
	var i, j uint
	for i = 0; i < power; i++ {
		tmp := ""
		for j = 0; j < uint(len(input)); j++ {
			if (i & (1 << j)) > 0 {
				tmp += fmt.Sprintf("%d", input[j])
			}
		}
		if len(tmp) > 0 {
			lookup[i] = tmp
		}
	}
	fmt.Println(lookup)
}

//total n*n=1/2

func prefixSum(input []int) {
	pSum := make([]int, len(input)+1)
	pSum[0] = 0
	for i := 0; i < len(input); i++ {
		pSum[i+1] = pSum[i] + input[i]
	}
	fmt.Print(pSum)
}

func laxicographicOrder(s, s1 string) string {
	res := ""
	for i := range s {
		for j := range s1 {
		}
	}
	return ""
}
