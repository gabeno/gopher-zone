package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(0.0)}
		wallet.Deposit(Bitcoin(10.0))

		assertBalance(t, wallet, Bitcoin(10.0))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20.0)}
		wallet.Withdraw(Bitcoin(5.0))

		assertBalance(t, wallet, Bitcoin(15.0))
	})
}
