package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(0.0)}
		wallet.Deposit(Bitcoin(10.0))

		got := wallet.Balance()
		want := Bitcoin(10.0)

		fmt.Printf("address of balance in test is %p \n", &wallet.balance)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20.0)}
		wallet.Withdraw(Bitcoin(5.0))

		got := wallet.Balance()
		want := Bitcoin(15.0)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
