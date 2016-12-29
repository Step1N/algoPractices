package stringmodule

import (
	"fmt"
	"testing"
)

func TestReduceString(t *testing.T) {
	fmt.Println(reduceStringToUniqe("aab", 0))
}
func TestCountWord(t *testing.T) {
	var testCase string
	_, _ = fmt.Scan(&testCase)
	fmt.Println(countWord(testCase))
}

func TestTwoCharacterProblem(t *testing.T) {
	fmt.Println(constructStringWithTwoChar("beabeefeab"))
}

func TestIsAnagram(t *testing.T) {
	fmt.Println(panagram("We promptly judged antique ivory buckles for the next prize"))
}

func TestFunnyString(t *testing.T) {

	fmt.Println(funnyString("acxz"))
	fmt.Println(funnyString("bcxz"))
}

func TestShiftString(t *testing.T) {
	fmt.Println(ShiftCharByGivenBit(byte(87), "W"))
}

func TestBeautifulString(t *testing.T) {
	fmt.Println(BeautifulString("0101010"))
}
