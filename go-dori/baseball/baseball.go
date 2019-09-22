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

func getInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	// error handling
	return scanner.Text()
}

func generateRandNum() string {
	randNumStr := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 4; i++ {
		randNum := rand.Intn(10)
		answer := strconv.Itoa(randNum)
		randNumStr += answer
	}
	return randNumStr
}

func checkAnswer(correctAnswer string, userAnswer string) bool {
	return correctAnswer == userAnswer
}

func displayScore(correctAnswer string, userAnswer string) string {
	out := 0
	strike := 0
	ball := 0

	// check numbers
	for i := 0; i < 4; i++ {
		s := userAnswer[i]
		if strings.Contains(correctAnswer, string(s)) {
			out++
		}

	}
	// check digits

	// convert
	score := "o" + strconv.Itoa(out) + " s" + strconv.Itoa(strike) + " b" + strconv.Itoa(ball)

	return score
}

func main() {
	prompt := "Type your answer: "
	scanner := bufio.NewScanner(os.Stdin)
	answer := generateRandNum()
	tries := 0

	for {
		if tries == 15 {
			fmt.Printf("Out of chances! The correct answer is %s bye bye...", answer)
		}
		// New Try
		tries += 1
		fmt.Printf(prompt)
		userAnswer := getInput(scanner)

		// Check Answer
		fmt.Printf("check answer: %t", checkAnswer(answer, userAnswer))

		// fail -> Display score

		// success -> end of the game
	}
}
