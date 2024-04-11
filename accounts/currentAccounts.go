package accounts

import "golang-oop/clientes"

type CurrentAccount struct {
	Holder                      clientes.Holder
	AgencyNumber, AccountNumber int
	balance                     float64
}

func (c *CurrentAccount) Withdraw(withdrawalValue float64) string {
	canWithdraw := withdrawalValue > 0 && withdrawalValue <= c.balance
	if canWithdraw {
		c.balance -= withdrawalValue
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *CurrentAccount) Deposit(depositAmount float64) (string, float64) {
	canDeposit := depositAmount > 0
	if canDeposit {
		c.balance += depositAmount
		return "Deposito realizado com sucesso", c.balance
	} else {
		return "Não foi possível realizar o deposito", c.balance
	}
}

func (c *CurrentAccount) Transferir(transferValue float64, destinationAccount *CurrentAccount) bool {
	if transferValue > 0 && transferValue < c.balance {
		c.balance -= transferValue
		destinationAccount.Deposit(transferValue)
		return true
	} else {
		return false
	}
}

func (c *CurrentAccount) GetBalance() float64 {
	return c.balance
}
