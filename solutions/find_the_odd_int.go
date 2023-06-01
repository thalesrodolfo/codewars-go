package solutions

import "fmt"

func FindOdd(seq []int) int {
	control := make(map[int]int)

	for i := 0; i < len(seq); i++ {
		if value, exists := control[seq[i]]; exists {
			control[seq[i]] = value + 1
		} else {
			control[seq[i]] = 1
		}
	}

	fmt.Println(control)

	for i := 0; i < len(seq); i++ {
		value := control[seq[i]]

		//fmt.Println(value)

		if value == 1 || value%2 != 0 {
			return seq[i]
		}
	}

	return 0
}
