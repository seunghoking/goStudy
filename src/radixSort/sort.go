package main

import "fmt"

func main() {
	numArray := [10]int{7, 4, 1, 2, 5, 8, 9, 5, 2, 1}
	sortArray := [10]int{}

	for i := 0; i < len(numArray); i++ {
		sortArray[numArray[i]]++
	}

	idx := 0
	for i := 0; i < len(sortArray); i++ {
		for j := 0; j < sortArray[i]; j++ {
			numArray[idx] = i
			idx++
		}
	}

	fmt.Println(numArray)
}
