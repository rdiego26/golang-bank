package accounts

import (
	"fmt"
	"golang-bank/customers"
	"golang-bank/utils"
)

type CheckingAccount struct {
	Owner                     customers.AccountOwner
	BranchCode, AccountNumber int
	balance                   float64
}

func (account *CheckingAccount) messageAfterOperation() {
	fmt.Printf("(%s) balance after operation -> %.2f\n", account.Owner.Name, account.balance)
}

func (account *CheckingAccount) GetBalance() float64 {
	return account.balance
}

func (account *CheckingAccount) Withdraw(requestedValue float64) Message {
	utils.TimeToCalculations()

	allowed := requestedValue > 0 && requestedValue <= account.balance
	if allowed {
		account.balance -= requestedValue
		account.messageAfterOperation()

		return OperationSuccessfully
	} else {
		fmt.Println("Your balance -> ", account.balance)
		return InsufficientFunds
	}
}

func (account *CheckingAccount) Deposit(requestedValue float64) Message {
	utils.TimeToCalculations()

	allowed := requestedValue > 0
	if allowed {
		account.balance += requestedValue
		account.messageAfterOperation()

		return OperationSuccessfully
	} else {
		fmt.Println("Your balance -> ", account.balance)
		return "Insufficient funds, we couldn't achieve your request!"
	}

}

func (account *CheckingAccount) Transfer(requestedValue float64, destinationAccount *CheckingAccount) Message {
	utils.TimeToCalculations()

	allowed := requestedValue > 0 && account.balance >= requestedValue
	if allowed {
		account.balance -= requestedValue
		destinationAccount.balance += requestedValue
		account.messageAfterOperation()

		return OperationSuccessfully
	} else {
		fmt.Println("Your balance -> ", account.balance)
		return InsufficientFunds
	}

}
