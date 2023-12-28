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

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("did not raise an error")
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

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10.0)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100.0))

		assertError(t, err)
		assertBalance(t, wallet, startingBalance)
	})
}
