package operations

import (
	"bank-go/accounts"
	"bank-go/clients"
	"bank-go/payments"
	"fmt"
	"os"
)

func AccountSavingOperation(client clients.Client) {
	account := getAccountSaving(client)
	for {
		showMenuAccountSaving()
		typeOperation := getCommandAccountSaving()
		if typeOperation == 0 {
			os.Exit(0)
		}
		operationsAccountSaving(typeOperation, &account)
	}
}

func getAccountSaving(client clients.Client) accounts.AccountSaving {
	var agencyNumber, accountNumber, operator int
	fmt.Println("What your agency number:")
	fmt.Scan(&agencyNumber)
	fmt.Println("What your account number:")
	fmt.Scan(&accountNumber)
	fmt.Println("What your operator:")
	fmt.Scan(&operator)
	fmt.Println("")
	account := accounts.AccountSaving{Owner: client, AgencyNumber: agencyNumber, AccountNumber: accountNumber, Operator: operator}
	return account
}

func showMenuAccountSaving() {
	fmt.Println("What kind of operation do you want to do:")
	fmt.Println("4 - Payment boleto")
	fmt.Println("3 - Deposit")
	fmt.Println("2 - Withdraw")
	fmt.Println("1 - Balance")
	fmt.Println("0 - Closed Program")
	fmt.Println("")
}

func getCommandAccountSaving() int {
	var command int
	fmt.Scan(&command)
	return command
}

func operationsAccountSaving(command int, account *accounts.AccountSaving) {
	switch command {
	case 1:
		fmt.Println("Total:", account.GetBalance())
		fmt.Println("")
	case 2:
		fmt.Println("What value of the withdrawing:")
		var value float64
		fmt.Scan(&value)
		fmt.Println(account.WithdrawMoney(value))
		fmt.Println("")
	case 3:
		fmt.Println("What value of the deposit:")
		var value float64
		fmt.Scan(&value)
		message, balance := account.Deposit(value)
		fmt.Println(message, ":", balance)
		fmt.Println("")
	case 4:
		fmt.Println("What value of the payment:")
		var value float64
		fmt.Scan(&value)
		fmt.Println(payments.PaymentBoleto(account, value))
		fmt.Println("")
	default:
		fmt.Println("Operation invalid")
		os.Exit(-1)
	}
}
