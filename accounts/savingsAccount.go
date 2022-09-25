package accounts

import (
	"fmt"
	"golang-bank/customers"
)

type SavingAccount struct {
	Owner                                customers.AccountOwner
	BranchCode, AccountNumber, Operation int
	balance                              float64
}

func (account *SavingAccount) messageAfterOperation() {
	fmt.Printf("(%s) balance after operation -> %.2f\n", account.Owner.Name, account.balance)
}

func (account *SavingAccount) GetBalance() float64 {
	return account.balance
}

func (account *SavingAccount) Withdraw(requestedValue float64) Message {
	timeToCalculations()

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

func (account *SavingAccount) Deposit(requestedValue float64) Message {
	timeToCalculations()

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

func (account *SavingAccount) Transfer(requestedValue float64, destinationAccount *SavingAccount) Message {
	timeToCalculations()

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
