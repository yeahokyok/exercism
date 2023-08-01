package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

func is_numeric(word string) bool {
	return regexp.MustCompile(`\d`).MatchString(word)
}

func Valid(id string) bool {

	id = strings.ReplaceAll(id, " ", "")
	if !is_numeric(id) {
		return false
	}

	if len(id) <= 1 {
		return false
	}

	doubled := make([]int, len(id))
	isOddID := false
	if len(id)%2 != 0 {
		isOddID = true
	}

	for i := len(doubled) - 1; i >= 0; i-- {
		if isOddID {
			if i%2 == 0 {
				val, err := strconv.Atoi(string(id[i]))
				if err != nil {
					return false
				}
				doubled[i] = val
				continue
			} else {
				val, err := strconv.Atoi(string(id[i]))
				if err != nil {
					return false
				}
				val *= 2
				if val > 9 {
					val -= 9
				}
				doubled[i] = val
				continue
			}

		}
		if i%2 == 0 {
			val, err := strconv.Atoi(string(id[i]))
			if err != nil {
				return false
			}
			val *= 2
			if val > 9 {
				val -= 9
			}
			doubled[i] = val
			continue
		}
		val, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return false
		}
		doubled[i] = val
	}

	sum := 0
	for _, num := range doubled {
		sum += int(num)
	}

	return sum%10 == 0
}
