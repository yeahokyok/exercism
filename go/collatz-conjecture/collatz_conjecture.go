package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be greater than zero")
	} else if n == 1 {
		return 0, nil
	} else if n%2 == 0 {
		res, err := CollatzConjecture(n / 2)
		if err != nil {
			return 0, err
		}
		return 1 + res, nil
	} else {
		res, err := CollatzConjecture(3*n + 1)
		if err != nil {
			return 0, err
		}
		return 1 + res, nil
	}
}
