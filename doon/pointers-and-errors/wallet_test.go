package pointers_and_errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {
	amount := 10
	wallet := Wallet{}
	wallet.Deposit(amount)

	got := wallet.Balance()
	want := amount

	assert.Equal(t, got, want)
}
