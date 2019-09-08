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
	Numbers    string
	Input      string
	Ball       int
	Strike     int
	Out        int
	IsCorrect  bool
	MaxRetries int
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
	a.IsCorrect = a.Strike == 4
}

func (a *Answer) Response() {
	fmt.Printf("%dO %dS %dB\n", a.Out, a.Strike, a.Ball)
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
		Numbers:    numbers,
		Ball:       0,
		Strike:     0,
		Out:        0,
		IsCorrect:  false,
		MaxRetries: 15,
	}
}

func GetPlayerInput(scanner *bufio.Scanner) (string, error) {
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return "", err
	}
	return scanner.Text(), nil
}

func GenerateRandomNumbers() string {
	numberArray := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(numberArray), func(i, j int) {
		numberArray[i], numberArray[j] = numberArray[j], numberArray[i]
	})
	return strings.Join(numberArray[:4], "")
}

func main() {
	prompt := "Guess What?: "
	scanner := bufio.NewScanner(os.Stdin)
	answer := NewAnswer()
	tries := 0

	for {
		tries += 1
		fmt.Print(prompt)
		input, err := GetPlayerInput(scanner)
		if err != nil {
			break
		}
		answer.SetInput(input)
		answer.Judge()
		answer.Response()
		if answer.IsCorrect == true {
			fmt.Println("🎉 Congratulations 🎉")
			break
		}
		if answer.MaxRetries == tries {
			fmt.Println("🐶 What a noob! The answer is ", answer.Numbers)
			break
		}
	}
}
