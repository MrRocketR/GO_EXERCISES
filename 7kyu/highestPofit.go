package _kyu

func MinMax(arr []int) [2]int {

	minimum, maximum := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maximum {
			maximum = arr[i]
		} else if arr[i] < minimum {
			minimum = arr[i]
		}
	}
	return [2]int{minimum, maximum}
}
