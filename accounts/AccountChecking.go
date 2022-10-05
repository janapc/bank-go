package accounts

import "bank-go/clients"

type AccountChecking struct {
	Owner         clients.Client
	AgencyNumber  int
	AccountNumber int
	balance       float64
}

func (a *AccountChecking) WithdrawMoney(withdrawValue float64) string {
	hasMoney := withdrawValue > 0 && withdrawValue <= a.balance
	if hasMoney {
		a.balance -= withdrawValue
		return "withdraw realized successfully"
	}
	return "Balance insufficient"
}

func (a *AccountChecking) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		a.balance += depositValue
		return "Deposit realized successfully", a.balance
	}
	return "Value less than zero", a.balance
}

func (a *AccountChecking) Transfer(transferValue float64, destinationAccount *AccountChecking) bool {
	if transferValue < a.balance && transferValue > 0 {
		a.balance -= transferValue
		destinationAccount.Deposit(transferValue)
		return true
	}
	return false
}

func (a *AccountChecking) GetBalance() float64 {
	return a.balance
}
