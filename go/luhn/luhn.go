package luhn

import (
	"strings"
	"unicode"
)

func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	if len(input) <= 1 {
		return false
	}

	sum := 0
	shouldDouble := false
	for i := len(input) - 1; i >= 0; i-- {
		if !unicode.IsDigit(rune(input[i])) { // check if not a digit
			return false
		}

		num := int(input[i] - '0') // convert to integer

		if shouldDouble {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}

		sum += num
		shouldDouble = !shouldDouble // toggle the double flag
	}

	return sum%10 == 0
}
