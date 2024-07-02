package pointers_errors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Testing deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Testing withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		err := wallet.Withdraw(Bitcoin(5))

		want := Bitcoin(5)

		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("Testing withdrawing more than the wallet's balance", func(t *testing.T) {
		startingBalance := Bitcoin(5)
		toWithdraw := Bitcoin(10)

		wallet := Wallet{balance: startingBalance}

		err := wallet.Withdraw(toWithdraw)

		assertError(t, err, ErrInsufficientFunds.Error())
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, err error, want string) {
	t.Helper()
	if err == nil {
		t.Fatal("withdrawing an amount greater than the wallet's balance should've returned an error")
	}

	if err.Error() != want {
		t.Errorf("got %s want %s", err, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an unexpected error")
	}
}
