package searching

import (
	"fmt"
)

func CountPairForWhoseDifferByK(in []int, diff int) int {
	pair := 0
	pairMap := make(map[int]int, 0)
	size := len(in)

	for i := 0; i < size; i++ {
		pairMap[in[i]] = 1
	}

	for i := 0; i < size; i++ {
		if pairMap[in[i]-diff] == 1 {
			pair++
		}
	}

	return pair
}

func CountPalindromeInGivenSequence(in string) int {
	count, size := 0, len(in)
	for i := 0; i < size; i++ {
		tmp := ""
		if isPalindrome(tmp) {
			fmt.Println(i, " p sequence found ", i)
			count++
		}
	}

	return count
}

func isPalindrome(s string) bool {
	palindrome := true
	size := len(s)
	i, j := 0, size-1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
			continue
		}
		palindrome = false
		break
	}

	return palindrome
}

func KindergartenProblem(in []int) int {
	size := len(in)
	count, max, minIndex, numberOfElement := 0, 0, -1, 0
	for i := 0; i < size; i++ {
		j := i
		for count < size {
			if in[j] <= count {
				numberOfElement++
			}
			count++
			j = (j + 1) % size
		}
		if max < numberOfElement {
			max = numberOfElement
			minIndex = i
		}
		count = 0
		numberOfElement = 0
	}

	return minIndex + 1
}

func KindergartenProblemOptimal(in []int) int {
	size := len(in)
	count, max, minIndex, numberOfElement := 0, 0, -1, 0
	for i := 0; i < size; i++ {
		j := i
		for count < size {
			if in[j] <= count {
				numberOfElement++
			}
			count++
			j = (j + 1) % size
		}
		if max < numberOfElement {
			max = numberOfElement
			minIndex = i
		}
		count = 0
		numberOfElement = 0
	}

	return minIndex + 1
}

func MaxSubArrayOfGivenSize(in []int, k int) int {
	size := len(in)
	max := 0
	for i := 0; i < size-k; i++ {
		for j := 0; j < k; j++ {
			if max < in[i+j] {
				max = in[i+j]
			}
		}
	}

	return max
}

func BinarySearchForSortedArray(in []int, l, r, x int) int {
	if r >= l {
		mid := l + (r-l)/2
		if in[mid] == x {
			return mid
		}
		if x < in[mid] {
			return BinarySearchForSortedArray(in, l, mid-1, x)
		} else {
			return BinarySearchForSortedArray(in, mid+1, r, x)
		}
	}

	return -1
}
func BinarySearchForSortedArrayIterative(in []int, l, r, x int) int {
	for l <= r {
		mid := l + (r-l)/2
		if in[mid] == x {
			return mid
		}
		if in[mid] < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

func BinarySearchForSortedShiftedArrayIterative(in []int, l, r, x int) int {
	for l <= r {
		mid := l + (r-l)/2
		if in[mid] == x {
			return mid
		}
		if in[mid] < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}









func FindSubArrayInUnsortedArrayWhichMakesArraySorted(in []int) int {
	size := len(in) - 1
	min, minIndex, maxIndex, max := 10000, 0, 0, 0
	for i := 1; i < size; i++ {
		if in[i-1] < in[i] {
			continue
		} else {
			maxIndex = i + 1
			max = in[i+1]
			if in[i] < min {
				minIndex = i-1
				min = in[i-1]
			}
			for j := i - 1; j >= 0; j-- {
				if in[i] < min && min < in[j] {
					minIndex = j-1
					min = in[j-1]
					continue
				} else {
					break
				}
			}
		}
	}
	fmt.Println(minIndex, " is min index for  ", min)
	fmt.Println(maxIndex, " is max index for  ", max)

	return -1
}

func FindKClosestElementToGivenVlaue(in []int, k, element int) int  {
	size := len(in) - 1
	countMap := make(map[int]int, 0)
	for i := 0; i < size; i++ {
		d := diff(element, k)

		if findInMap(countMap, d){

		}else if len(countMap) <=  k{
			countMap[in[i]] = diff(element, in[i])
		}
	}
	return -1
}

func findInMap(m map[int]int, el, newEL int) bool {
	valid :=false
	for key, value :=range m{
		if el < value {
			m[key] = newEL
			valid = true
			break
		}
	}
	return valid

}

func diff(a, b int) int {
	if a > b{
		return a-b
	}else{
		return b-a
	}
}

