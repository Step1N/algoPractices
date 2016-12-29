package consolereader

import (
	"bufio"
	"fmt"
	"os"
)

//To read input as sentence
func takeStringInput() string {
	consolereader := bufio.NewReader(os.Stdin)
	input, _ := consolereader.ReadString('\n')
	fmt.Println(input)
	return input
}

//To read input as space seprated number
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

//To read input for matrix
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
