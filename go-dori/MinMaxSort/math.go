package math

func Min(numbers []int) (minVal int) {
	for idx, number := range numbers {
		if idx == 0 || number < minVal {
			minVal = number
		}
	}
	return
}

func Max(numbers []int) (maxVal int) {
	for idx, number := range numbers {
		if idx == 0 || maxVal < number {
			maxVal = number
		}
	}
	return
}

func BubbleSort(arr []int) (sortedArr []int) {
	size := len(arr)
	for i := size - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j+1] < arr[j] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	sortedArr = append(arr)
	return
}

func SelectionSort(arr []int) (sortedArr []int) {
	size := len(arr)
	min := Min(arr)
	for i := 0; i < size-1; i++ {
		if min < arr[i] {
			temp := arr[i]
			arr[i] = temp
			temp = min
		} else {
			min = Min(arr[i:])
		}
	}
	sortedArr = append(arr)
	return
}
