package arrayProblem

import (
	"algoPractices/sorting"
	"fmt"
	"math"
)

type Vendor struct {
	start int
	end   int
	price int
}

func TakeInput(in int) (vendors []Vendor) {
	vendors = make([]Vendor, in)
	for i := 0; i < in; i++ {
		var st, en, pr int
		_, _ = fmt.Scan(&st, &en, &pr)
		v := Vendor{start: st, end: en, price: pr}
		vendors = append(vendors, v)
	}

	return vendors
}
func FindMinimumSellerPrice(vendors []Vendor, st, en int) {
	filterVendors := make([]Vendor, len(vendors))
	for i := 0; i < len(vendors); i++ {
		for j := 1; j < len(vendors)-1; j++ {

		}
		fmt.Println(filterVendors[i].start, filterVendors[i].end)
	}
}

//Circular shift of element
func shiftCircularByGivenBit(shift int, input []int) []int {
	size := len(input)
	outputArray := make([]int, size)
	for i := 0; i < size; i++ {
		outputArray[(i+shift)%size] = input[i]
	}

	return outputArray
}

//Count number positive and negative in array
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

//Find difference of digonal of a matrix
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

func printShiftedElement(noElm, shift, queries int) {
	outputArray := shiftCircularByGivenBit(shift, takeArrayInput(noElm))
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

func sumOfArray(input []int) (sum int) {
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}

	return
}

func ComputePrefixSum(in []int, m int) int {
	size := len(in)
	prefixSum := make([]int, 5)
	k := 0
	for i := 0; i < size; i++ {
		sum := 0
		for j := 0; j <= i; j++ {
			sum += in[j]
		}
		prefixSum[k] = sum % m
		k++
	}
	maxSum := 0
	for i := 0; i < size; i++ {
		for j := i - 1; j >= 0; j-- {
			maxSum = maxInt(maxSum, (prefixSum[i]-prefixSum[j]+m)%m)

		}
		maxSum = maxInt(maxSum, prefixSum[i])
	}

	return maxSum
}

func findMinDiff(in []int) int {
	size := len(in)
	sortedOutput := sorting.QuickSort(in)
	fmt.Println(in)
	fmt.Println(sortedOutput)
	diff := 10000
	for i := 0; i < size-1; i++ {
		fmt.Println(sortedOutput[i+1], sortedOutput[i])
		t := sortedOutput[i+1] - sortedOutput[i]
		if t < diff {
			diff = t
		}

	}
	return diff
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/*	Find Longest Increasing Subsequence (LIS) in given array
	In: {10, 22, 9, 33, 21, 50, 41, 60, 80} is 6 and LIS is
	Out: {10, 22, 33, 50, 60, 80}.
*/
func findLongestIncreasingSubSequenceBF(input []int) {
	lookup := make(map[int][]int, 0)
	size := len(input)
	for i := 0; i < size; i++ {
		tmp := make([]int, 0)
		f := input[i]
		tmp = append(tmp, f)
		for j := i + 1; j < size; j++ {
			if f < input[j] {
				tmp = append(tmp, input[j])
				f = input[j]
			}
		}
		lookup[len(tmp)] = tmp
	}

	fmt.Println(findMaxArrayFromMap(lookup))
}

func findMaxArrayFromMap(in map[int][]int) []int {
	min := 0
	for in := range in {
		if min < in {
			min = in
		}
	}
	return in[min]
}

//LIS longest increasing sub sequence find max in current aaray findLongestIncreasingSubSequenceDP(in, el)
func findLISSequenceRC(in []int, n, mLength int) int {
	if n == 1 {
		return 1
	}
	cLength := 1
	for i := 0; i < n; i++ {
		subLen := findLISSequenceRC(in, i, mLength)
		if in[i] < in[n-1] && cLength < (1+subLen) {
			cLength = 1 + subLen
		}
	}
	if mLength < cLength {
		mLength = cLength
	}

	return cLength
}

func findLISSequenceDP(in []int, n int) int {
	lookup := make(map[int]int, 0)
	for i := 0; i < n; i++ {
		lookup[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if in[i] > in[j] && lookup[i] < lookup[j]+1 {
				lookup[i] = lookup[j] + 1
			}
		}
	}

	min := 0
	for in := range in {
		if min < in {
			min = in
		}
	}

	return min
}

/*
For two sequence find larges common sequence
in: AGGTAB, GXTXAYB
out: GTAB
*/
func findLCSInSequence(s, s1 string) string {
	lastMatchIndex := 0
	firstMatched := false
	matched := ""
	j := 0
	for i := 0; i < len(s); i++ {
		if firstMatched {
			j = lastMatchIndex + 1
		}
		for ; j < len(s1); j++ {
			if s[i] == s1[j] {
				matched += string(s[i])
				lastMatchIndex = j
				firstMatched = true
				break
			}
		}
	}

	return matched
}

//If matched pass substring of matched sequence
func findLCSInSequenceRec(s, s1 string, m, n int) int {
	if m == 0 || n == 0 {
		return 1
	}
	if s[m-1] == s1[n-1] {
		return 1 + findLCSInSequenceRec(s, s1, m-1, n-1)
	} else {
		return max(findLCSInSequenceRec(s, s1, m, n-1), findLCSInSequenceRec(s, s1, m-1, n))
	}
}

func findLCSInSequenceDP(s, s1 string, m, n int) int {
	lookup := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		lookup[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			lookup[i][j] = 0
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				lookup[i][j] = 0
			} else if s[i-1] == s1[j-1] {
				lookup[i][j] = lookup[i-1][j-1] + 1
			} else {
				lookup[i][j] = max(lookup[i-1][j], lookup[i][j-1])
			}
		}
	}
	fmt.Print(lookup[m-1][n-1])
	return lookup[m-1][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func CountSubArrayHavingSumDivisibleByK(in []int, k int) int {
	size := len(in)
	tmp := make([]int, size)
	sum := 0
	for i := 0; i < size; i++ {
		sum += in[i]
		tmp[((sum%k)+k)%k]++
	}
	res := 0
	for i := 0; i < k; i++ {
		if tmp[i] > 0 {
			res += (tmp[i] * (tmp[i] - 1)) / 2
		}
	}
	res += tmp[0]
	fmt.Println(tmp)

	return res
}

func CountMaxSumOfSubSubArray(in []int) int {
	maxEnd, maxSoFor := in[0], in[0]
	for i := 1; i < len(in); i++ {
		tmp := in[i]
		maxEnd = max(tmp, maxEnd+tmp)
		maxSoFor = max(maxSoFor, maxEnd)
	}

	return maxSoFor
}

func FindNumberOfSubArrayWhichMaxElementGreaterThanGiven(in []int, k int) int {
	size, i, sum := 0, 0, len(in)
	for i < size {
		if in[i] > k {
			i++
			continue
		}
		count := 0
		for i < size && in[i] <= k {
			i++
			count++
		}
		sum += ((count * (count + 1)) / 2)
	}

	return ((size*(size+1))/2 - sum)
}

func FindMaxIntegerOccurredInRange(start, end []int) int {
	size := 100
	tmp := make([]int, size)
	for i := 0; i < len(start); i++ {
		tmp[start[i]] = tmp[start[i]] + 1
		tmp[end[i]+1] -= 1
	}
	maxSum, index := tmp[0], 0
	for i := 1; i < size; i++ {
		tmp[i] += tmp[i-1]
		if maxSum < tmp[i] {
			maxSum = tmp[i]
			index = i
		}
	}

	return index
}

func Find3MForGivenInput(in []int) int {
	size := len(in)
	var sum int
	inputMap := make(map[int]int, 0)
	for i := 0; i < size; i++ {
		k := in[i]
		sum += k
		if inputMap[k] >= 1 {
			t := inputMap[k]
			t++
			inputMap[k] = t
		} else {
			inputMap[k] = 1
		}
	}
	fmt.Println(fmt.Sprintf("%.1f", float64(sum)/float64(size)))
	mxOccurrence, mOccurrenceValue := 0, 0
	for key, value := range inputMap {
		if value > mxOccurrence {
			mxOccurrence = value
			mOccurrenceValue = key
		}
	}
	for key, value := range inputMap {
		if value == mxOccurrence && key < mOccurrenceValue {
			mOccurrenceValue = key
		}
	}
	fmt.Println(mOccurrenceValue)

	sorting.QuickSort(in)
	var median float64
	if size%2 == 0 {
		tmpIndex := size / 2
		tmpSum := in[tmpIndex-1] + in[tmpIndex]
		median = (float64(tmpSum) / float64(2))
	} else {
		tmpIndex := size / 2
		median = float64(in[tmpIndex-1])
	}
	fmt.Println(median)

	return 0
}

func FindQuartilesGivenInput(in []int) (median float64, tmpIndex int) {
	size := len(in)
	if size%2 == 0 {
		tmpIndex = size / 2
		tmpSum := in[tmpIndex-1] + in[tmpIndex]
		median = (float64(tmpSum) / float64(2))
	} else {
		tmpIndex = size / 2
		median = float64(in[tmpIndex])
	}
	fmt.Println(in)
	return
}

func FindSDForGivenArray(in []int) {
	size := len(in)
	var sum int
	for i := 0; i < size; i++ {
		k := in[i]
		sum += k
	}
	mean := int(float64(sum) / float64(size))
	sum = 0
	for i := 0; i < size; i++ {
		k := (mean - in[i]) * (mean - in[i])
		sum += k
	}
	sQuerSum := float64(sum) / float64(size)
	stDeviation := math.Sqrt(sQuerSum)

	fmt.Println(fmt.Sprintf("%.1f", stDeviation))
}

//Equal Stacks
func PlayGameOfStone(tCase int) {
	turn := true
	canRemoved := false
	numberOfWinByP1 ,numberOfWinByP2 := 0, 0
	for i := 0; i < tCase; i++ {
		var k int
		_, _ = fmt.Scan(&k)
		if k%2 == 0 || k%3 == 0 || k%5 == 0 {
			canRemoved = true
		}
		if canRemoved {
			if turn {
				fmt.Println("First")
				numberOfWinByP1++
			} else {
				fmt.Println("Second")
				numberOfWinByP2++
			}
		} else {
			if turn {
				fmt.Println("First")
				turn = false
			} else {
				fmt.Println("Second")
				turn = true
			}
		}
	}
}



