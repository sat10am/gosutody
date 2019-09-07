package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Answer struct {
	Numbers string
	Input   string
	Ball    int
	Strike  int
	Out     int
}

func (a *Answer) SetInput(input string) {
	a.Input = input
}

func (a *Answer) initializeState() {
	a.Ball = 0
	a.Strike = 0
	a.Out = 0
}

func (a *Answer) Judge() {
	a.initializeState()
	for index, rune_ := range a.Input {
		key := string(rune_)
		i := strings.Index(a.Numbers, key)
		if i == -1 {
			a.Out += 1
		} else if index == i {
			a.Strike += 1
		} else {
			a.Ball += 1
		}
	}
	fmt.Println(a.Strike)
}

func NewAnswer(params ...int) *Answer {
	numbers := ""
	if len(params) != 4 {
		numbers = GenerateRandomNumbers()
	} else {
		tmp := []string{}
		for _, number := range params {
			tmp = append(tmp, strconv.Itoa(number))
		}
		numbers = strings.Join(tmp, "")
	}
	return &Answer{
		Numbers: numbers,
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
