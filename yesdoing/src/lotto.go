package main

import (
	"errors"
	"math/rand"
)

type Lotto struct {
	numbers []int
}

func (l *Lotto) WriteNumbers(numbers []int) error {

	if len(numbers) > 7 {
		return errors.New("over length")
	}

	for _, i := range numbers {
		if i > 45 || i < 1 {
			return errors.New("Number limit over!!")
		}
	}

	l.numbers = numbers

	return nil
}

func (l *Lotto) GetNumbers() []int {
	return l.numbers
}

func GetRamdomTicket() *Lotto {
	numbers := RamdomMachine()
	ticket := &Lotto{}
	ticket.WriteNumbers(numbers)
	return ticket
}

func RamdomMachine() []int {
	list := rand.Perm(45)

	for i, _ := range list {
		list[i]++
	}

	return list[:7]
}
