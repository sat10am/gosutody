package main

import (
	"reflect"
	"testing"
)

func TestGetLottoTicket(t *testing.T) {
	t.Run("get 7 Number in Lotto Ticket", func(t *testing.T) {
		lotto := Lotto{}
		lotto.WriteNumbers([]int{1, 2, 3, 4, 5, 6, 7})

		got := lotto.GetNumbers()
		want := []int{1, 2, 3, 4, 5, 6, 7}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("create 7 length Random Numbers ", func(t *testing.T) {
		random := RamdomMachine()

		length := len(random)

		if length != 7 {
			t.Errorf("length %d is not length 7", length)
		}
	})

	t.Run("create Random Lotto ticket", func(t *testing.T) {
		ramdomNumberTicket := GetRamdomTicket()

		if len(ramdomNumberTicket.GetNumbers()) != 7 {
			t.Errorf("got Ticket %v", ramdomNumberTicket)
		}
	})

}
