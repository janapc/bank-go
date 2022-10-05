package main

import (
	"bank-go/clients"
	"bank-go/operations"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the Bank")
	showMenuTypeAccount()
	typeAccount := getCommand()
	client := getClient()
	if typeAccount == 1 {
		operations.AccountCheckingOperation(client)
	} else if typeAccount == 2 {
		operations.AccountSavingOperation(client)
	}
	os.Exit(0)
}

func showMenuTypeAccount() {
	fmt.Println("What type of account you can open:")
	fmt.Println("2 - account saving")
	fmt.Println("1 - account checking")
	fmt.Println("0 - closed program")
}

func getCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}

func getClient() clients.Client {
	var name, cpf, profession string
	fmt.Println("What your name:")
	fmt.Scan(&name)
	fmt.Println("What your cpf:")
	fmt.Scan(&cpf)
	fmt.Println("What your profession:")
	fmt.Scan(&profession)
	client := clients.Client{
		Name:       name,
		CPF:        cpf,
		Profession: profession,
	}
	return client
}
