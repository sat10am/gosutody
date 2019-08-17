package main

func Min(numbers []int) int {
	index := numbers[0]

	for _, number := range numbers {
		if index > number {
			index = number
		}
	}

	return index
}

func Max(numbers []int) int {
	index := numbers[0]

	for _, number := range numbers {
		if index < number {
			index = number
		}
	}

	return index
}
