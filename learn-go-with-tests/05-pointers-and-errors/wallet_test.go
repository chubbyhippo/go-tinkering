package _5_pointers_and_errors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("should deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		fmt.Printf("address of balance in test is %v \n", &wallet.balance)
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("should withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		fmt.Printf("address of balance in test is %v \n", &wallet.balance)
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
