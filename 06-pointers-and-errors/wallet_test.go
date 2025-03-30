package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(20))
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func BenchmarkDeposit(b *testing.B) {
	wallet := Wallet{}

	for b.Loop() {
		wallet.Deposit(Bitcoin(10))
	}
}

func BenchmarkWithdraw(b *testing.B) {
	wallet := Wallet{balance: Bitcoin(1000)}

	for b.Loop() {
		wallet.Withdraw(Bitcoin(1))
	}
}

func BenchmarkBalance(b *testing.B) {
	wallet := Wallet{balance: Bitcoin(100)}

	for b.Loop() {
		wallet.Balance()
	}
}

func ExampleWallet_Deposit() {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	fmt.Println(wallet.balance)
	// Output: 10 BTC
}

func ExampleWallet_Withdraw() {
	wallet := Wallet{balance: Bitcoin(20)}
	err := wallet.Withdraw(Bitcoin(10))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(wallet.Balance())
	}
	// Output: 10 BTC
}

func ExampleWallet_Withdraw_insufficientFunds() {
	wallet := Wallet{balance: Bitcoin(5)}
	err := wallet.Withdraw(Bitcoin(10))
	if err != nil {
		fmt.Println(err)
	}
	// Output: cannot withdraw, insufficient funds
}
