package sorting

import "errors"

const EmptySliceErrorMsg = "empty slice not allowed"

func SelectionSort(numbers []int) ([]int, error) {
	SIZE := len(numbers)
	if SIZE == 0 {
		return numbers, errors.New(EmptySliceErrorMsg)
	}

	ret := make([]int, SIZE)

	for index, number := range numbers {
		min := number
		ret[index] = min
		for _, innerNumber := range numbers {
			if innerNumber < min {
				ret[index] = innerNumber
			}
		}
	}
	return ret, nil
}
