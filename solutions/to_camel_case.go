package solutions

import "strings"

func ToCamelCase(s string) string {
	var newString strings.Builder
	var uppercaseNext bool = false

	for i := 0; i < len(s); i++ {
		var char = string(s[i])

		if char == "-" || char == "_" {
			uppercaseNext = true
		} else {
			if uppercaseNext {
				newString.WriteString(strings.ToUpper(char))
				uppercaseNext = false
			} else {
				newString.WriteString(char)
			}

		}
	}

	return newString.String()
}
