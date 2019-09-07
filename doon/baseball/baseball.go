package baseball

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Answer struct {
	Numbers string
	Ball    int
	Strike  int
	Out     int
}

func NewAnswer() *Answer {
	return &Answer{
		Numbers: GenerateRandomNumbers(),
		Ball:    0,
		Strike:  0,
		Out:     0,
	}
}

func GetPlayerInput(scanner *bufio.Scanner) (string, error) {
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return scanner.Text(), nil
}

func GenerateRandomNumbers() string {
	numberArray := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Seed(time.Now().UnixNano())
	numbers := ""
	for i := 0; i < 4; i++ {
		rand.Shuffle(len(numberArray), func(i, j int) {
			numberArray[i], numberArray[j] = numberArray[j], numberArray[i]
		})
		numbers += strconv.Itoa(numberArray[0])
		numberArray = numberArray[1:]
	}
	return numbers
}

// func CheckPlyerInput(input string, answer Answer) Answer {
// }

func main() {
	prompt := "Guess What?: "
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(prompt)
	line, err := GetPlayerInput(scanner)
	if err != nil {
		fmt.Println("Input values are invalid")
	}
	fmt.Println(line)
}
