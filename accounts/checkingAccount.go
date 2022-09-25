package accounts

import (
	"fmt"
	"time"
)

type CheckingAccount struct {
	HolderName    string
	BranchCode    int
	AccountNumber int
	Balance       float64
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
	fmt.Printf("(%s) Balance after operation -> %.2f\n", account.HolderName, account.Balance)
}

func (account *CheckingAccount) Withdraw(requestedValue float64) Message {
	timeToCalculations()

	allowed := requestedValue > 0 && requestedValue <= account.Balance
	if allowed {
		account.Balance -= requestedValue
		messageAfterOperation(account)

		return OperationSuccessfully
	} else {
		fmt.Println("Your Balance -> ", account.Balance)
		return InsufficientFunds
	}
}

func (account *CheckingAccount) Deposit(requestedValue float64) Message {
	timeToCalculations()

	allowed := requestedValue > 0
	if allowed {
		account.Balance += requestedValue
		messageAfterOperation(account)

		return OperationSuccessfully
	} else {
		fmt.Println("Your Balance -> ", account.Balance)
		return "Insufficient funds, we couldn't achieve your request!"
	}

}

func (account *CheckingAccount) Transfer(requestedValue float64, destinationAccount *CheckingAccount) Message {
	timeToCalculations()

	allowed := requestedValue > 0 && account.Balance >= requestedValue
	if allowed {
		account.Balance -= requestedValue
		destinationAccount.Balance += requestedValue
		messageAfterOperation(account)

		return OperationSuccessfully
	} else {
		fmt.Println("Your Balance -> ", account.Balance)
		return InsufficientFunds
	}

}
