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
