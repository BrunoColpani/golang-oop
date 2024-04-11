package accounts

import "golang-oop/clientes"

type SavingsAccount struct {
	Holder                                 clientes.Holder
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (c *SavingsAccount) Withdraw(withdrawalValue float64) string {
	canWithdraw := withdrawalValue > 0 && withdrawalValue <= c.balance
	if canWithdraw {
		c.balance -= withdrawalValue
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *SavingsAccount) Deposit(depositAmount float64) (string, float64) {
	canDeposit := depositAmount > 0
	if canDeposit {
		c.balance += depositAmount
		return "Deposito realizado com sucesso", c.balance
	} else {
		return "Não foi possível realizar o deposito", c.balance
	}
}

func (c *SavingsAccount) GetBalance() float64 {
	return c.balance
}
