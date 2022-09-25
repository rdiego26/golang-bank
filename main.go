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

func timeToCalculations() {
	time.Sleep(1 * time.Second)
}

func (account *Account) withdraw(requestedValue float64) string {
	timeToCalculations()

	allowed := requestedValue > 0 && requestedValue <= account.balance
	if allowed {
		account.balance -= requestedValue
		fmt.Println("Balance after withdraw -> ", account.balance)
		return "Withdrawal made successfully, thanks for using our services!"
	} else {
		fmt.Println("Your balance -> ", account.balance)
		return "Insufficient funds, we couldn't achieve your request!"
	}
}

func (account *Account) deposit(requestedValue float64) string {
	timeToCalculations()

	allowed := requestedValue > 0
	if allowed {
		account.balance += requestedValue
		fmt.Println("Balance after deposit -> ", account.balance)
		return "Deposit made successfully, thanks for using our services!"
	} else {
		fmt.Println("Your balance -> ", account.balance)
		return "Insufficient funds, we couldn't achieve your request!"
	}

}

func main() {
	newAccount := Account{"Ramos", 123, 9837, 30.5}
	fmt.Println("Initial balance -> ", newAccount.balance)

	fmt.Println("Trying withdraw operation with 13...")
	fmt.Println(newAccount.withdraw(13))
	fmt.Println("")
	fmt.Println("Trying withdraw operation with 130...")
	fmt.Println(newAccount.withdraw(130))

	fmt.Println("")
	fmt.Println("Trying deposit operation with 130...")
	fmt.Println(newAccount.deposit(130))

}
