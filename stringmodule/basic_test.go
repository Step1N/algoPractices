package stringmodule

import (
	"fmt"
	"testing"
)

func TestReduceString(t *testing.T) {
	fmt.Println(ReduceStringToUnique("aaabccddd"))
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

func TestFindValidString(t *testing.T) {
	//s := "ibfdgaeadiaefgbhbdghhhbgdfgeiccbiehhfcggchgghadhdhagfbahhddgghbdehidbibaeaagaeeigffcebfbaieggabcfbiiedcabfihchdfabifahcbhagccbdfifhghcadfiadeeaheeddddiecaicbgigccageicehfdhdgafaddhffadigfhhcaedcedecafeacbdacgfgfeeibgaiffdehigebhhehiaahfidibccdcdagifgaihacihadecgifihbebffebdfbchbgigeccahgihbcbcaggebaaafgfedbfgagfediddghdgbgehhhifhgcedechahidcbchebheihaadbbbiaiccededchdagfhccfdefigfibifabeiaccghcegfbcghaefifbachebaacbhbfgfddeceababbacgffbagidebeadfihaefefegbghgddbbgddeehgfbhafbccidebgehifafgbghafacgfdccgifdcbbbidfifhdaibgigebigaedeaaiadegfefbhacgddhchgcbgcaeaieiegiffchbgbebgbehbbfcebciiagacaiechdigbgbghefcahgbhfibhedaeeiffebdiabcifgccdefabccdghehfibfiifdaicfedagahhdcbhbicdgibgcedieihcichadgchgbdcdagaihebbabhibcihicadgadfcihdheefbhffiageddhgahaidfdhhdbgciiaciegchiiebfbcbhaeagccfhbfhaddagnfieihghfbaggiffbbfbecgaiiidccdceadbbdfgigibgcgchafccdchgifdeieicbaididhfcfdedbhaadedfageigfdehgcdaecaebebebfcieaecfagfdieaefdiedbcadchabhebgehiidfcgahcdhcdhgchhiiheffiifeegcfdgbdeffhgeghdfhbfbifgidcafbfcd"
	//s := "aabbcd"
	//s:="aaaabbbb"
	s := "aaaaabbbb"
	fmt.Println(FindValidString(s))
}

func TestAllCombination(t *testing.T) {
	input := []int{3, 3, 9, 9, 5}
	prefixSum(input)
}
