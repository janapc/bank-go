package operations

import (
	"bank-go/accounts"
	"bank-go/clients"
	"bank-go/payments"
	"fmt"
	"os"
)

func AccountCheckingOperation(client clients.Client) {
	account := getAccountChecking(client)
	for {
		showMenuAccountChecking()
		typeOperation := getCommandAccountChecking()
		if typeOperation == 0 {
			os.Exit(0)
		}
		operationsAccountChecking(typeOperation, &account)
	}
}

func getAccountChecking(client clients.Client) accounts.AccountChecking {
	var agencyNumber, accountNumber int
	fmt.Println("What your agency number:")
	fmt.Scan(&agencyNumber)
	fmt.Println("What your account number:")
	fmt.Scan(&accountNumber)
	fmt.Println("")
	account := accounts.AccountChecking{Owner: client, AgencyNumber: agencyNumber, AccountNumber: accountNumber}
	return account
}

func showMenuAccountChecking() {
	fmt.Println("What kind of operation do you want to do:")
	fmt.Println("5 - Payment Boleto")
	fmt.Println("4 - Deposit")
	fmt.Println("3 - Withdraw")
	fmt.Println("2 - Transfer")
	fmt.Println("1 - Balance")
	fmt.Println("0 - Closed Program")
	fmt.Println("")
}

func getCommandAccountChecking() int {
	var command int
	fmt.Scan(&command)
	return command
}

func operationsAccountChecking(command int, account *accounts.AccountChecking) {
	switch command {
	case 1:
		fmt.Println("Total:", account.GetBalance())
		fmt.Println("")
	case 2:
		fmt.Println("What value to the transfer:")
		var value float64
		fmt.Scan(&value)
		var agencyNumber, accountNumber int
		fmt.Println("What agency number:")
		fmt.Scan(&agencyNumber)
		fmt.Println("What account number:")
		fmt.Scan(&accountNumber)
		accountTransfer := accounts.AccountChecking{Owner: clients.Client{}, AgencyNumber: agencyNumber, AccountNumber: accountNumber}
		fmt.Println(account.Transfer(value, &accountTransfer))
		fmt.Println("")
	case 3:
		fmt.Println("What value of the withdrawing:")
		var value float64
		fmt.Scan(&value)
		fmt.Println(account.WithdrawMoney(value))
		fmt.Println("")
	case 4:
		fmt.Println("What value of the deposit:")
		var value float64
		fmt.Scan(&value)
		fmt.Println(account.Deposit(value))
		fmt.Println("")
	case 5:
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
