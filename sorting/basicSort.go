package sorting

import (
	"fmt"
	"math/rand"
	"goUtils/stack"
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
		}
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

func sortUsingTwoStack(in []int)  {
	s1 := stack.NewStack()
	s2 := stack.NewStack()
	s1.Push(in[0])
	s1.Push(in[1])
	for i := 0; i<len(in);i++  {
		p := s2.Peek().(int)
		if in[i] < p{
			s2.Push(in[i])
		}
	}
}