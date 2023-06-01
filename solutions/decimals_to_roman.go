package solutions

import "strings"

func Solution(number int) string {
	var builder strings.Builder

	for number > 0 {
		if number/1000 > 0 {
			builder.WriteString("M")
			number = number - 1000
		} else if number/900 > 0 {
			builder.WriteString("CM")
			number = number - 900
		} else if number/500 > 0 {
			builder.WriteString("D")
			number = number - 500
		} else if number/400 > 0 {
			builder.WriteString("CD")
			number = number - 400
		} else if number/100 > 0 {
			builder.WriteString("C")
			number = number - 100
		} else if number/90 > 0 {
			builder.WriteString("XC")
			number = number - 90
		} else if number/50 > 0 {
			builder.WriteString("L")
			number = number - 50
		} else if number/40 > 0 {
			builder.WriteString("XL")
			number = number - 40
		} else if number/10 > 0 {
			builder.WriteString("X")
			number = number - 10
		} else if number/9 > 0 {
			builder.WriteString("IX")
			number = number - 9
		} else if number/5 > 0 {
			builder.WriteString("V")
			number = number - 5
		} else if number/4 > 0 {
			builder.WriteString("IV")
			number = number - 4
		} else {
			builder.WriteString("I")
			number = number - 1
		}
	}

	return builder.String()
}
