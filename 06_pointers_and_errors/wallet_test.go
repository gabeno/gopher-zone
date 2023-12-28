package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10.0)

	got := wallet.Balance()
	want := 10.0

	fmt.Printf("address of balance in test is %p \n", &wallet.balance)

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
