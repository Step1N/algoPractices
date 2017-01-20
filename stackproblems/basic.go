package stack

import (
	"fmt"
	s "github.com/Step1N/goUtils/stack"
	"github.com/golang/go/src/pkg/strings"
	"math"
)

func findSolidRectangle(in []int) int {
	maxArea := 0
	size := len(in)
	for i := 0; i < size; i++ {
		tmpArea := in[i] * (size - i)
		if tmpArea > maxArea {
			maxArea = tmpArea
		}
	}

	return maxArea
}

func findNumberOfDayToDiePlants(in []int) int {
	size := len(in)
	stk := s.NewStack()
	tmp := make([]int, size)
	min, max := in[0], 0

	for i := 1; i < size; i++ {
		if in[i] > in[i-1] {
			tmp[i] = 1
		}
		if in[i] < min {
			min = in[i]

		}
		for !stk.IsEmpty() && in[stk.Peek().(int)] >= in[i] {
			if min < in[i] {
				if tmp[i] < tmp[stk.Peek().(int)]+1 {
					tmp[i] = tmp[stk.Peek().(int)] + 1
				}
			}
			stk.Pop()
		}
		fmt.Println("after ", i, " iteration", tmp)
		if max < tmp[i] {
			max = tmp[i]
		}
		stk.Push(i)
	}
	return max
}

func findDiff(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func computeBitwiseEquation(in []int) int {
	s := s.NewStack()
	s.Push(in[0])
	s.Push(in[1])
	max := in[0] ^ in[1]
	for i := 2; i < len(in); i++ {
		for !s.IsEmpty() && s.Peek().(int) >= in[i] {
			c := in[i] ^ s.Peek().(int)
			s.Pop()
			if max < c {
				max = c
			}
		}
		if !s.IsEmpty() {
			c := in[i] ^ s.Peek().(int)
			s.Pop()
			if max < c {
				max = c
			}
		}

		s.Push(in[i])
	}

	return max
}

type element struct {
	data int
	max  int
}

func SimpleStackOperation(in int, s1 *s.Stack) (input int) {
	switch in {
	case 1:
		var t int
		_, _ = fmt.Scan(&t)
		el := element{data: t, max: t}
		if !s1.IsEmpty() {
			p := s1.Peek().(element)
			if el.data < p.data {
				el.max = p.max
			}
		}
		s1.Push(el)
		break
	case 2:
		s1.Pop()
		break
	case 3:
		input = s1.Peek().(element).data
		break
	}

	return
}

func BracketBalanced(in string, s1 *s.Stack) string {
	sAr := strings.Split(in, "")
	for i := 0; i < len(sAr); i++ {
		c := string(in[i])
		if c == "{" || c == "[" || c == "(" {
			s1.Push(c)
		}
		if c == "}" || c == "]" || c == ")" {
			if s1.IsEmpty() {
				return "NO"
			} else {
				t := s1.Peek().(string)
				opp := findClosingOfBrackets(t)
				if c == opp {
					s1.Pop()
				}
			}
		}
	}
	if s1.IsEmpty() {
		return "YES"
	}

	return "NO"
}

func findClosingOfBrackets(in string) string {
	opp := ""
	switch in {
	case "{":
		opp = "}"
		break
	case "[":
		opp = "]"
		break
	case "(":
		opp = ")"
		break
	}

	return opp
}

func SampleEditorWithStack(s string, operation, index int, s1 *s.Stack) {

	switch operation {
	case 1:
		/*var s string
		_, _ = fmt.Scan(&s)*/
		if !s1.IsEmpty() {
			t := s1.Peek().(string)
			t += s
			s1.Push(t)
		} else {
			s1.Push(s)
		}
		break
	case 2:
		/*var index int
		_, _ = fmt.Scan(&index)*/
		t := s1.Peek().(string)
		t = t[:len(t)-index]
		s1.Push(t)
		break
	case 3:
		/*var index int
		_, _ = fmt.Scan(&index)*/
		t := s1.Peek().(string)
		fmt.Printf("%s \n", string(t[index-1]))
		break
	case 4:
		s1.Pop()
		break
	}
}

func LargestRactangl(input []int) int {
	s1 := *s.NewStack()
	maxArea, area, topIndex, i := 0, 0, 0, 0
	for i < len(input) {
		if s1.IsEmpty() || input[(s1.Peek().(int))] <= input[i] {
			s1.Push(i)
			i++
		} else {
			topIndex = s1.Peek().(int)
			s1.Pop()
			if s1.IsEmpty() {
				area = input[topIndex] * i
			} else {
				area = input[topIndex] * (i - s1.Peek().(int) - 1)
			}
			if maxArea < area {
				maxArea = area
			}
		}
	}
	for !s1.IsEmpty() {
		topIndex = s1.Peek().(int)
		s1.Pop()
		if s1.IsEmpty() {
			area = input[topIndex] * i
		} else {
			area = input[topIndex] * (i - s1.Peek().(int) - 1)
		}
		if maxArea < area {
			maxArea = area
		}
	}

	return maxArea
}

func EqualStackSize(in1, in2, in3 []int) int {
	sum := 0

	size1 := len(in1)
	s1, s2, s3 := s.NewStack(), s.NewStack(), s.NewStack()
	for i := size1 - 1; i >= 0; i-- {
		sum += in1[i]
		s1.Push(in1[i])
	}
	s1.Push(sum)

	sum = 0
	size2 := len(in2)
	for i := size2 - 1; i >= 0; i-- {
		sum += in2[i]
		s2.Push(in2[i])
	}
	s2.Push(sum)

	sum = 0
	size3 := len(in3)
	for i := size3 - 1; i >= 0; i-- {
		sum += in3[i]
		s3.Push(in3[i])
	}
	s3.Push(sum)

	a, b, c := s1.Pop().(int), s2.Pop().(int), s3.Pop().(int)
	for !s1.IsEmpty() && !s2.IsEmpty() && !s3.IsEmpty() {
		if a > b {
			a = a - s1.Pop().(int)
		} else if b > c {
			b = b - s2.Pop().(int)
		} else {
			c = c - s3.Pop().(int)
		}
		if a == b && b == c {
			break
		}
	}

	return a
}

func WaiterConfusion(in []int, prime int) {
	inStack := s.NewStack()
	for i := 0; i < len(in); i++ {
		inStack.Push(in[i])
	}
	outStack := s.NewStack()
	tmpStack := s.NewStack()
	primes := FindNPrimeNumber(prime)

	for i := 0; i < len(primes); i++ {
		for !inStack.IsEmpty() {
			el := inStack.Pop().(int)
			if el%primes[i] == 0 {
				tmpStack.Push(el)
			} else {
				outStack.Push(el)
			}
		}
		for !tmpStack.IsEmpty() {
			fmt.Println(tmpStack.Pop().(int))
		}
		inStack = outStack
		outStack = s.NewStack()

	}
	for !inStack.IsEmpty() {
		fmt.Println(inStack.Pop().(int))
	}
}

func FindNPrimeNumber(number int) (primeNumber []int) {
	if number == 1 {
		return []int{2}
	}

	num := 2
	isPrime := true
	for i := 2; i <= number+1; {
		for j := 2; j <= int(math.Sqrt(float64(num))); j++ {
			if num%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primeNumber = append(primeNumber, num)
			i++
		}
		num++
		isPrime = true
	}

	return primeNumber
}
