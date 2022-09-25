package accounts

import (
	"fmt"
	"golang-bank/customers"
	"time"
)

type CheckingAccount struct {
	Owner         customers.AccountOwner
	BranchCode    int
	AccountNumber int
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

func messageAfterOperation(account *CheckingAccount) {
	fmt.Printf("(%s) balance after operation -> %.2f\n", account.Owner.Name, account.balance)
}

func (account *CheckingAccount) GetBalance() float64 {
	timeToCalculations()

	return account.balance
}

func (account *CheckingAccount) Withdraw(requestedValue float64) Message {
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

func (account *CheckingAccount) Deposit(requestedValue float64) Message {
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

func (account *CheckingAccount) Transfer(requestedValue float64, destinationAccount *CheckingAccount) Message {
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
