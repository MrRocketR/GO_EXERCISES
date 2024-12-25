package main

import "fmt"

func main() {
	fmt.Println(MinMax([]int{1, 2, 3, 4, 5}))
	fmt.Println(MinMax([]int{2334454, 5}))
}

/*
[1,2,3,4,5] --> [1,5]
[2334454,5] --> [5,2334454]
[1]         --> [1,1]
*/

func MinMax(arr []int) [2]int {
	maximum := arr[0]
	minimum := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > maximum {
			maximum = arr[i]
		}
		if arr[i] < minimum {
			minimum = arr[i]
		}
	}
	return [2]int{minimum, maximum}
}
