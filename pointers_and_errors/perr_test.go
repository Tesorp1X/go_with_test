package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("one deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("two deposits", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(20)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
