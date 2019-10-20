package main

import (
	"math/rand"
)

type Lotto struct {
	numbers []int
}

func (l *Lotto) WriteNumbers(numbers []int) {
	l.numbers = numbers
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
