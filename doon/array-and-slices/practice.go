package array_and_slices

import "errors"

func Max(numbers []int) (int, error) {

	if len(numbers) == 0 {
		return -1, errors.New("not allowed empty array")
	}

	ret := numbers[0]
	for _, number := range numbers {
		if number > ret {
			ret = number
		}
	}
	return ret, nil
}

func Min(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return -1, errors.New("not allowed empty array")
	}

	ret := numbers[0]
	for _, number := range numbers {
		if number < ret {
			ret = number
		}
	}

	return ret, nil
}
