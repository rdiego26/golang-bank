package main

import (
	"fmt"
	"time"
)

func namedParameters() {
	var holderName string = "Ramos"
	var branchCode int = 589
	var accountNumber int = 123456
	var balance float64 = 125.50

	newAccount := Account{balance: 100.00}

	fmt.Println(newAccount)
	fmt.Println(Account{holderName, accountNumber, branchCode, balance})
}

func comparisonTypes() {
	var crisAccount *Account
	crisAccount = new(Account)
	crisAccount.holderName = "Cris"
	crisAccount.balance = 500

	var crisAccount2 *Account
	crisAccount2 = new(Account)
	crisAccount2.holderName = "Cris"
	crisAccount2.balance = 500

	fmt.Println(*crisAccount)

	// Comparing references(not values)
	fmt.Println("crisAccount == crisAccount2 ->", crisAccount == crisAccount2)

	// Comparing values
	fmt.Println("*crisAccount == *crisAccount2 ->", *crisAccount == *crisAccount2)

}

type Account struct {
	holderName    string
	branchCode    int
	accountNumber int
	balance       float64
}

type Message string

const (
	OperationSuccessfully Message = "Operation made successfully, thanks for using our services!"
	InsufficientFunds             = "Insufficient funds, we couldn't achieve your request!"
)

func timeToCalculations() {
	time.Sleep(1 * time.Second)
}

func messageAfterOperation(account *Account) {
	fmt.Printf("(%s) Balance after operation -> %.2f\n", account.holderName, account.balance)
}

func (account *Account) withdraw(requestedValue float64) Message {
	timeToCalculations()

	allowed := requestedValue > 0 && requestedValue <= account.balance
	if allowed {
		account.balance -= requestedValue
		messageAfterOperation(account)

		return OperationSuccessfully
	} else {
		fmt.Println("Your balance -> ", account.balance)
		return InsufficientFunds
	}
}

func (account *Account) deposit(requestedValue float64) Message {
	timeToCalculations()

	allowed := requestedValue > 0
	if allowed {
		account.balance += requestedValue
		messageAfterOperation(account)

		return OperationSuccessfully
	} else {
		fmt.Println("Your balance -> ", account.balance)
		return "Insufficient funds, we couldn't achieve your request!"
	}

}

func (account *Account) transfer(requestedValue float64, destinationAccount *Account) Message {
	timeToCalculations()

	allowed := requestedValue > 0 && account.balance >= requestedValue
	if allowed {
		account.balance -= requestedValue
		destinationAccount.balance += requestedValue
		messageAfterOperation(account)

		return OperationSuccessfully
	} else {
		fmt.Println("Your balance -> ", account.balance)
		return InsufficientFunds
	}

}

func main() {
	firstAccount := Account{"Ramos", 123, 9837, 30.5}
	secondAccount := Account{"Silvia", 123, 9838, 130.5}
	fmt.Printf("Initial balance(%s) -> %.2f\n", firstAccount.holderName, firstAccount.balance)
	fmt.Printf("Initial balance(%s) -> %.2f\n", secondAccount.holderName, secondAccount.balance)

	fmt.Println("")
	fmt.Printf("(%s) Trying withdraw operation with 13...\n", firstAccount.holderName)
	fmt.Println(firstAccount.withdraw(13))
	fmt.Println("")
	fmt.Printf("(%s) Trying withdraw operation with 130...\n", firstAccount.holderName)
	fmt.Println(firstAccount.withdraw(130))

	fmt.Println("")
	fmt.Printf("(%s) Trying deposit operation with 130...\n", firstAccount.holderName)
	fmt.Println(firstAccount.deposit(130))

	fmt.Println("")
	fmt.Printf("Actual balance(%s) -> %.2f\n", secondAccount.holderName, secondAccount.balance)
	fmt.Printf("(%s) Trying transfer operation with 100 to (%s)...\n", secondAccount.holderName,
		firstAccount.holderName)
	fmt.Println(secondAccount.transfer(100, &firstAccount))
}
