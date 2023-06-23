package main

import (
	"log"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, expect Bitcoin) {
		t.Helper()
		response := wallet.Balance()

		if response != expect {
			t.Errorf("response: %s,\nexpect: %s\n", response, expect)

		}

	}

	t.Run("Deposite", func(t *testing.T) {
		wallet := Wallet{}
		amount := 10
		wallet.Deposite(Bitcoin(amount))
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 30}
		amount := 10
		err := wallet.Withdraw(Bitcoin(amount))
		if err != nil {
			log.Fatal(err)
		}
		assertBalance(t, wallet, Bitcoin(20))
	})

	t.Run("Withdraw fail", func(t *testing.T) {
		wallet := Wallet{balance: 30}
		err := wallet.Withdraw(Bitcoin(40))
		if err == nil {

		}
	})
}
