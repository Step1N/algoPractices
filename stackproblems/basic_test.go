package stack

import (
	"fmt"
	"github.com/Step1N/goUtils/stack"
	"testing"
)

func TestFindMaxRectangle(t *testing.T) {
	//10 6320 6020 6098 1332 7263 672 9472 28338 3401 9494
	//18060
	s := []int{10, 6320, 6020, 6098, 1332, 7263, 672, 9472, 8338, 3401, 9494}
	ts := findSolidRectangle(s)
	fmt.Println(ts)
}

func TestFindNumberOfDayToDiePlant(t *testing.T) {
	in := []int{6, 5, 8, 4, 7, 10, 9}
	fmt.Println(findNumberOfDayToDiePlants(in))
}

func TestAndXorOr(t *testing.T) {
	//Input - 76969694 71698884 32888447 31877010 65564584 87864180 7850891 1505323 17879621 15722446
	//Output - 114613468
	//in := []int{76969694, 71698884, 32888447, 31877010, 65564584, 87864180, 7850891, 1505323, 17879621, 15722446}
	in := []int{15, 13, 1}
	fmt.Println(computeBitwiseEquation(in))

}

func TestBracketBalanced(t *testing.T) {
	//Input - 76969694 71698884 32888447 31877010 65564584 87864180 7850891 1505323 17879621 15722446
	//Output - 114613468
	//in := []int{76969694, 71698884, 32888447, 31877010, 65564584, 87864180, 7850891, 1505323, 17879621, 15722446}
	//in := "}([]]][[){}}[[)}[(}(}]{(}[{}][{}](}]}))]{][[}(({(]}[]{[{){{(}}[){[][{[]{[}}[)]}}]{}}(}"
	//in := "{}{}()}[]]{((([](){}([[]]((){[(({}())){}]({}{{{}[({{}})]}[]}[]{}[])})()))))}}}]}]})}}]}])}}]])}}})}]}]})]))}}})}}]]))"
	in := "()({}[[[[()]]]])()[]{}{()}({()}[(){}[[[{{[()[[({[]([(()())])})[[]({}([()]))]]{[()]{}[]}({()([{}{({{[()][]}{[{[(({})){}]{[[]([])]}}[{}]{[()[][]]}][]}[](({()()}(){()}()))})}])}[])]{}"
	s1 := stack.NewStack()
	fmt.Println(BracketBalanced(in, s1))

}

func TestSampleEditor(t *testing.T) {
	s1 := stack.NewStack()
	SampleEditorWithStack("giihangdtzjavcaxnbtudpcnmzggbdabocpeplzercnctkcvhoujbrpbztjlpuymabumbppgrydkdzbadcsjdtpdftlusapza", 1, 0, s1)
	SampleEditorWithStack("", 3, 6, s1)
	SampleEditorWithStack("", 4, 0, s1)
	SampleEditorWithStack("fcsivbvhtzydbhhopdnlaisgqhlkizqhezglxocqgymjedw", 1, 0, s1)
	SampleEditorWithStack("", 2, 24, s1)
	SampleEditorWithStack("", 3, 10, s1)
	SampleEditorWithStack("", 3, 11, s1)

	SampleEditorWithStack("bgqrvdnlmkjdbmishfpaazvujrkzctgfbizhhrrdgdtzxzysm", 1, 0, s1)
	SampleEditorWithStack("", 4, 0, s1)
	SampleEditorWithStack("tixqvwpiufqrqhgldlpaizamxxuivbhlmubnqpalwdowlf", 1, 0, s1)
	SampleEditorWithStack("", 2, 60, s1)
	SampleEditorWithStack("zsfbxajobmdztghgddnsmxfceaudttvjqqvehvpodkexaj", 1, 0, s1)
	SampleEditorWithStack("", 2, 16, s1)
	SampleEditorWithStack("", 2, 20, s1)
	SampleEditorWithStack("", 4, 0, s1)
	SampleEditorWithStack("yqcpbiuggbcpohztsrdx", 1, 0, s1)
	SampleEditorWithStack("", 4, 0, s1)
	SampleEditorWithStack("", 2, 11, s1)
	SampleEditorWithStack("", 3, 27, s1)
	SampleEditorWithStack("", 3, 15, s1)

}

func TestLargestRactangl(t *testing.T) {
	//Input : 6320 6020 6098 1332 7263 672 9472 2838 3401 9494
	//Output : 18060
	in := []int{6320, 6020, 6098, 1332, 7263, 672, 9472, 2838, 3401, 9494}
	fmt.Println(LargestRactangl(in))
}

func TestEqualStackSize(t *testing.T) {
	in1 := []int{3, 2, 1, 1, 1}
	in2 := []int{4, 3, 2}
	in3 := []int{1, 1, 4, 1}
	fmt.Println(EqualStackSize(in1, in2, in3))
}

func TestWaiterConfusion(t *testing.T) {
	//Input: 47 21
	//Input : 80 37 86 79 8 39 43 41 15 33 30 15 45 55 61 74 49 49 20 66 77 19 85 44 81 82 27 5 36 83 91 45 39 44 19 44 71 49 8 66 81 40 29 60 35 31 44
	//output : [{80 86 8 30 74 20 66 44 82 36 44 44 8 66 40 60 44},
	// 	    {81 39 45 27 81 45 15 33 15 39}
	// 	    {55 85 5 35 }
	// 	    {49 91 77 49 49}
	// 	    {19 19 29 31 37 41 43 61 71 79 83}
	in := []int{80, 37, 86, 79, 8, 39, 43, 41, 15, 33, 30, 15, 45, 55, 61, 74, 49, 49, 20, 66, 77, 19, 85, 44, 81, 82, 27, 5, 36, 83, 91, 45, 39, 44, 19, 44, 71, 49, 8, 66, 81, 40, 29, 60, 35, 31, 44}
	//in := []int{3, 4, 7, 6, 5}
	//in := []int{4, 4, 3, 9, 9}
	WaiterConfusion(in, 21)
}

func TestFindNPrimeNumber(t *testing.T) {
	fmt.Println(FindNPrimeNumber(3))
}
