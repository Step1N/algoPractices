package main

import (
	st "algoPractice/disjoint"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	var s int
	_, _ = fmt.Scan(&s)

	fmt.Println(st.FindAstronautsForMoonMission(n, s))

	/*var testCase, elements int
	_, _ = fmt.Scan(&testCase)
	input := make([][]int, 6)
	for i := 0; i < testCase; i++ {
		_, _ = fmt.Scan(&elements)
		input[i] = takeArrayInput(elements)
	}
	for i := 0; i < len(input); i++ {
		out := make([][]int, 6)
		_ = findRandomPermutation(input[i], out, 0, 6)
		perSize := len(out)
		prmArray := make([][]int, perSize)
		for i := 0; i < perSize; i++ {
			prmArray[i] = findAllVertex(out[i])
		}
		fmt.Println(prmArray)
		//fmt.Println(findAvgOfPermutation(prmArray))
	}*/

}

func computPermutation(in int) {

}
func findAvgOfPermutation(input [][]int) string {
	size := len(input)
	sum := 0
	for i := 0; i < size; i++ {
		for j := 0; j < len(input[i]); j++ {
			sum += input[i][j]
		}
	}
	return fmt.Sprintf("%.2f", float64(sum)/float64(size))
}
func findRandomPermutation(input []int, output [][]int, index, npr int) int {
	size := len(input)
	out := make([]int, size)
	if index >= size-1 {
		for i := 0; i < size; i++ {
			out[i] = input[i]
		}
		npr--
		output[npr] = out
		return npr
	}
	for i := index; i < size; i++ {
		swap(input, index, i)
		npr = findRandomPermutation(input, output, index+1, npr)
		swap(input, index, i)
	}

	return npr
}

func swap(input []int, index1, index2 int) {
	tmp := input[index1]
	input[index1] = input[index2]
	input[index2] = tmp
}

func findAllVertex(input []int) []int {
	size := len(input)
	vertex := make([]int, size)
	for i := 0; i < size; i++ {
		vertex[i] = findDistanceFromXAxis(i, input)
	}

	return vertex
}

func findDistanceFromXAxis(curIndex int, input []int) int {
	distanceFromX := 1
	if curIndex == 0 {
		return distanceFromX
	}
	for i := curIndex - 1; i >= 0; i-- {
		if input[i] > input[curIndex] {
			return distanceFromX
		} else if input[i] < input[curIndex] {
			distanceFromX++
		}
	}

	return distanceFromX
}

func chocolateDistribute(input []int, count int) int {
	if !isAllElementSame(input) {
		max := findMaxInArray(input)
		output := distributeChocolates(max, input)
		count++
		count = chocolateDistribute(output, count)
	}

	return count
}

func isAllElementSame(input []int) bool {
	for i := 1; i < len(input); i++ {
		if input[0] != input[i] {
			return false
		}
	}

	return true
}

func distributeChocolates(max int, input []int) []int {
	resultArray := make([]int, len(input))
	diff := findMaxDiffInArray(input)
	resultArray = increaseArrayElementByDiff(diff, max, input)
	return resultArray
}

func increaseArrayElementByDiff(diff, max int, input []int) []int {
	for i := 0; i < len(input); i++ {
		if max != input[i] {
			input[i] = input[i] + diff
		}
	}
	return input
}

func findMaxDiffInArray(input []int) int {
	maxDiff := -1
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			diff := input[i] - input[j]
			if maxDiff < diff {
				maxDiff = diff
			}
		}
	}
	return maxDiff
}

func findMaxInArray(input []int) int {
	max := 0
	for i := 0; i < len(input); i++ {
		if max < input[i] {
			max = input[i]
		}
	}
	return max
}

func isEvenArray(input []int) bool {
	isEven := true
	for i := 0; i < len(input); i++ {
		if input[i]%2 != 0 {
			isEven = false
			break
		}
	}
	return isEven
}
func printShiftedElement(noElm, shift, queries int) {
	outputArray := shiftArrayByGivenBit(shift, takeArrayInput(noElm))
	outputElement := make([]int, queries)
	for i := 0; i < queries; i++ {
		var in int
		if _, err := fmt.Scanf("%d", &in); err != nil {
			fmt.Println("  Scan for index failed, due to ", err)
			return
		}
		outputElement[i] = in
	}
	for i := 0; i < len(outputElement); i++ {
		fmt.Println(outputArray[outputElement[i]])
	}
}
func shiftArrayByGivenBit(shift int, input []int) []int {
	size := len(input)
	outputArray := make([]int, size)
	for i := 0; i < size; i++ {
		outputArray[(i+shift)%size] = input[i]
	}

	return outputArray
}

func convertDigitalToMilitaryTime(digitalTime string) {
	pmTime := strings.Index(digitalTime, "PM")
	if pmTime >= 0 {
		tm, _ := strconv.Atoi(digitalTime[:strings.Index(digitalTime, ":")])
		if tm < 12 {
			digitalTime = fmt.Sprintf("%v%v", tm+12, digitalTime[strings.Index(digitalTime, ":"):])
		}
	}
	amTime := strings.Index(digitalTime, "AM")
	if amTime >= 0 {
		tm, _ := strconv.Atoi(digitalTime[:strings.Index(digitalTime, ":")])
		if tm == 12 {
			digitalTime = fmt.Sprintf("00%v", digitalTime[strings.Index(digitalTime, ":"):])
		}
	}

	fmt.Println(digitalTime[:len(digitalTime)-2])
}

func printStairCase(n int) {
	for i := 0; i < n; i++ {
		for j := (n - 2) - i; j >= 0; j-- {
			fmt.Print("", i)
		}
		for j := n - i; j <= n; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func countNumberTypeInArray(input []int) {
	size := len(input)
	positive, negative, zero := 0, 0, 0
	for i := 0; i < size; i++ {
		if input[i] > 0 {
			positive++
		} else if input[i] < 0 {
			negative++
		} else {
			zero++
		}
	}
	fmt.Printf("%.6f\n", float64(positive)/float64(size))
	fmt.Printf("%.6f\n", float64(float64(negative)/float64(size)))
	fmt.Printf("%.6f\n", float64(float64(zero)/float64(size)))
}

func takeInputForMatrix(size int) [][]int {
	input := make([][]int, size)
	for i := 0; i < size; i++ {
		input[i] = make([]int, size)
		for j := 0; j < size; j++ {
			var k int
			_, err := fmt.Scan(&k)
			if err != nil {
				fmt.Printf("  Scan for k failed, due to %v \n", err)
				continue
			}
			input[i][j] = k
		}
	}
	return input

}

func findDigonalDiff(input [][]int) int {
	left, right := findDigonalsOfMatrix(input)
	sLeft, sRight := 0, 0
	for i := 0; i < len(left); i++ {
		sLeft += left[i]
	}
	for i := 0; i < len(right); i++ {
		sRight += right[i]
	}
	diff := sLeft - sRight
	if sRight > sLeft {
		diff = sRight - sLeft
	}

	return diff
}

func findDigonalsOfMatrix(input [][]int) ([]int, []int) {
	size := len(input)
	leftDigonal, rightDigonal := make([]int, size), make([]int, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j {
				leftDigonal[i] = input[i][j]
			}
		}
	}

	for i := 0; i < size; i++ {
		rightDigonal[i] = input[i][(size-1)-i]
	}

	return leftDigonal, rightDigonal
}

func takeInt64ArrayInput(length int) []int64 {
	input := make([]int64, 0)
	for i := 0; i < length; i++ {
		var k int64
		_, err := fmt.Scan(&k)
		if err != nil {
			fmt.Errorf("  Scan for k failed, due to ", err)
			continue
		}
		input = append(input, k)
	}
	return input
}

func sumOfInt64Array(input []int64) (sum int64) {
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}

	return
}

func takeArrayInput(length int) []int {
	input := make([]int, 0)
	for i := 0; i < length; i++ {
		var k int
		_, err := fmt.Scan(&k)
		if err != nil {
			fmt.Errorf("  Scan for k failed, due to ", err)
			continue
		}
		input = append(input, k)
	}
	return input
}

func sumOfArray(input []int) (sum int) {
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}

	return
}

func tripletProblem(aTriplet, bTribple []int) {
	aScore, bScore := 0, 0
	for i := 0; i < 3; i++ {
		if aTriplet[i] > bTribple[i] {
			aScore++
		} else if aTriplet[i] < bTribple[i] {
			bScore++
		}
	}
	fmt.Printf("%d %d \n", aScore, bScore)
}
