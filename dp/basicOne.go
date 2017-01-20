package dp

func sherlockSum(input []int, sum, i, j, pass int) int {
	if len(input) <= 3 {
		return absDiff(input[1], input[0])
	}
	mid := (i + j) / 2
	pass++
	t := sherlockSum(input[:mid], sum, 0, mid, pass) + sherlockSum(input[mid+1:], sum, mid+1, len(input), pass)
	sum += t

	return sum * pass
}

func findBestCombination(in []int) {
	size := len(in)
	count := 0
	for i := 2; i < size; i++ {
		if in[i] > in[i-1] {
			count++
			continue
		}
	}

	if count == size/2 {
		for i := 2; i < size; i++ {
			if in[i] > in[i-1] {
				in[i-1] = 1
				continue
			}
		}
	}
}

func absDiff(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}
