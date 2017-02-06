package sorting

import (
	"fmt"
	"math/rand"
)

func InsertionByShift(in []int) []int {
	size := len(in)
	pivot := in[size-1]
	in[size-1] = 0
	for i := size - 2; i >= 0; i-- {
		if in[i] > pivot {
			in[i+1] = in[i]
			printInput(in)
			fmt.Println()
		} else if in[i] < pivot {
			in[i+1] = pivot
			break
		}
	}
	if in[0] == in[1] {
		in[0] = pivot
	}
	return in
}
func InsertionByShiftAll(in []int) []int {
	size := len(in)
	for i := 1; i < size; i++ {
		for j := i; j > 0; j-- {
			if in[j] < in[j-1] {
				in[j], in[j-1] = in[j-1], in[j]
			}
		}
		printInput(in)
		fmt.Println()
	}
	return in
}

func printInput(in []int) {
	for i := 0; i < len(in); i++ {
		fmt.Printf("%d ", in[i])
	}
}

func CountNumberOfOccurence(in []int) []int {
	size := len(in)
	out := make([]int, size)
	countMap := make(map[int]int, size)
	for i := 0; i < size; i++ {
		if countMap[in[i]] != 0 {
			count := countMap[in[i]]
			count += 1
			countMap[in[i]] = count
		} else {
			countMap[in[i]] = 1
		}
	}

	for i := 0; i < size; i++ {
		if countMap[i] != 0 {
			out[i] = countMap[i]
		} else {
			out[i] = 0
		}
	}
	return out
}

func QuickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1
	pivotIndex := rand.Int() % len(a)
	a[pivotIndex], a[right] = a[right], a[pivotIndex]
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]

	QuickSort(a[:left])
	QuickSort(a[left+1:])
	return a
}

func DivideAndConqur(in []int) {
	size := len(in)
	pivot := in[0]
	l := make([]int, 0)
	r := make([]int, 0)
	e := make([]int, 0)
	for i := 0; i < size; i++ {
		if in[i] < pivot {
			l = append(l, in[i])
		} else if in[i] > pivot {
			r = append(r, in[i])
		} else if in[i] == pivot {
			e = append(e, in[i])
		}
	}
	printInput(l)
	printInput(e)
	printInput(r)
}

func QuickSortPartition(in []int) []int {
	size := len(in)
	left := make([]int, 0)
	right := make([]int, 0)
	pivot := 0
	if size < 2 {
		return in
	} else {
		pivot = in[0]
		for i := 1; i < size; i++ {
			if in[i] < pivot {
				left = append(left, in[i])
			} else if in[i] > pivot {
				right = append(right, in[i])
			}
		}
	}

	left = QuickSortPartition(left)
	right = QuickSortPartition(right)
	left = append(left, pivot)
	left = append(left, right...)

	printInput(left)
	fmt.Println()

	return left
}

func QuickSortInPlace(in []int, low, high int) {
	l, r := low, high
	if l < r {
		pIndex := LoumoloPartitionUtil(in, l, r)
		printInput(in)
		fmt.Println()
		QuickSortInPlace(in, l, pIndex-1)
		QuickSortInPlace(in, pIndex+1, r)
	}
}

func LoumoloPartitionUtil(in []int, l, r int) int {
	i := l - 1
	pivot := in[r]
	for j := l; j <= r-1; j++ {
		if in[j] <= pivot {
			i++
			in[i], in[j] = in[j], in[i]
		}
	}
	in[i+1], in[r] = in[r], in[i+1]

	return i + 1
}

func QuickSortInPlaceWithPartition(in []int, low, high int) {
	l, r := low, high
	if l < r {
		pIndex := CommonPartitionUtil(in, l, r)
		QuickSortInPlaceWithPartition(in, l, pIndex-1)
		QuickSortInPlaceWithPartition(in, pIndex+1, r)
	}

	printInput(in)
	fmt.Println()
}

func CommonPartitionUtil(in []int, l, r int) int {
	pivot := in[l]
	i, j := l, r
	for i < j {
		for in[i] <= pivot && i < j {
			i++
		}
		for in[j] > pivot {
			j--
		}
		if i < j {
			in[i], in[j] = in[j], in[i]
		}
	}
	in[l], in[j] = in[j], in[l]

	return j
}

func CountSort(in []int) {
	size := len(in)
	tmp := make([]int, size)
	for i := 0; i < size; i++ {
		if tmp[in[i]] >= 1 {
			tmp[in[i]]++
			continue
		}
		tmp[in[i]] = 1
	}
	for  i := 0; i < 100; i++ {
		for tmp[i]>= 0{
			fmt.Printf("%v ", i)
			tmp[i]--
		}
	}
	fmt.Println()
}

func CountSortHashMap(in []int) {
	size := len(in)
	tmp := make(map[int]int, size)
	for i := 0; i < size; i++ {
		tmp[in[i]]++
	}

	for  i := 0; i < 100; i++ {
		for tmp[i]>= 0{
			fmt.Printf("%v ", i)
			tmp[i]--
		}
	}

	fmt.Println()
}
