package accounts

import "bank-go/clients"

type AccountSaving struct {
	Owner                                 clients.Client
	AgencyNumber, AccountNumber, Operator int
	balance                               float64
}

func (a *AccountSaving) WithdrawMoney(withdrawValue float64) string {
	hasMoney := withdrawValue > 0 && withdrawValue <= a.balance
	if hasMoney {
		a.balance -= withdrawValue
		return "Withdraw realized successfully"
	}
	return "Balance insufficient"
}

func (a *AccountSaving) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		a.balance += depositValue
		return "Deposit realized successfully", a.balance
	}
	return "Value less than zero", a.balance
}

func (a *AccountSaving) GetBalance() float64 {
	return a.balance
}
