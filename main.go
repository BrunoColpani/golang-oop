package main

import (
	"fmt"
	"golang-oop/accounts"
)

type AccountVerify interface {
	Withdraw(valor float64) string
}

func main() {
	brunoAccount := accounts.CurrentAccount{}
	brunoAccount.Deposit(100)
	payBill(&brunoAccount, 60)

	fmt.Println(brunoAccount.GetBalance())
}

func payBill(account AccountVerify, boletoValor float64) {
	account.Withdraw(boletoValor)
}
