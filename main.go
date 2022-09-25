package main

import "fmt"

type Account struct {
	holderName    string
	branchCode    int
	accountNumber int
	balance       float64
}

func main() {
	var holderName string = "Ramos"
	var branchCode int = 589
	var accountNumber int = 123456
	var balance float64 = 125.50

	newAccount := Account{balance: 100.00}

	var crisAccount *Account
	crisAccount = new(Account)
	crisAccount.holderName = "Cris"
	crisAccount.balance = 500

	var crisAccount2 *Account
	crisAccount2 = new(Account)
	crisAccount2.holderName = "Cris"
	crisAccount2.balance = 500

	fmt.Println(newAccount)
	fmt.Println(Account{holderName, accountNumber, branchCode, balance})
	fmt.Println(*crisAccount)

	// Comparing references(not values)
	fmt.Println("crisAccount == crisAccount2 ->", crisAccount == crisAccount2)

	// Comparing values
	fmt.Println("*crisAccount == *crisAccount2 ->", *crisAccount == *crisAccount2)

}
