package iteration

func Repeat(char string, count int) (repeated string) {
	for i := 0; i < count; i++ {
		repeated += char
	}
	return
}
