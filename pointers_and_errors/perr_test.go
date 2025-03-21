package pointersanderrors

import (
	"testing"
)

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestWallet(t *testing.T) {
	t.Run("one deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("two deposits", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(20))
	})
}

func TestWallet_Withdraw(t *testing.T) {
	type args struct {
		amount Bitcoin
	}
	tests := []struct {
		name    string
		w       *Wallet
		args    args
		want    Bitcoin
		wantErr bool
		err     string
	}{
		{
			name:    "withdraw: no overdraft",
			w:       &Wallet{Bitcoin(20)},
			args:    args{amount: Bitcoin(10)},
			want:    Bitcoin(10),
			wantErr: false,
			err:     "",
		},
		{
			name:    "withdraw: overdraft scenario",
			w:       &Wallet{Bitcoin(10)},
			args:    args{amount: Bitcoin(20)},
			want:    Bitcoin(10),
			wantErr: true,
			err:     ErrInsufficientFunds.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.w.Withdraw(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("Wallet.Withdraw() error = %v, wantErr %v", err, tt.err)
			}
			assertBalance(t, *tt.w, tt.want)
		})
	}
}
