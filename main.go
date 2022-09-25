package main

import (
	"fmt"
	c "golang-bank/accounts"
)

func namedParameters() {
	var holderName = "Ramos"
	var branchCode = 589
	var accountNumber = 123456
	var balance = 125.50

	newAccount := c.CheckingAccount{Balance: 100.00}

	fmt.Println(newAccount)
	fmt.Println(c.CheckingAccount{holderName, accountNumber, branchCode, balance})
}

func comparisonTypes() {
	var crisAccount *c.CheckingAccount
	crisAccount = new(c.CheckingAccount)
	crisAccount.HolderName = "Cris"
	crisAccount.Balance = 500

	var crisAccount2 *c.CheckingAccount
	crisAccount2 = new(c.CheckingAccount)
	crisAccount2.HolderName = "Cris"
	crisAccount2.Balance = 500

	fmt.Println(*crisAccount)

	// Comparing references(not values)
	fmt.Println("crisAccount == crisAccount2 ->", crisAccount == crisAccount2)

	// Comparing values
	fmt.Println("*crisAccount == *crisAccount2 ->", *crisAccount == *crisAccount2)

}

func main() {
	firstAccount := c.CheckingAccount{HolderName: "Ramos", BranchCode: 123, AccountNumber: 9837, Balance: 30.5}
	secondAccount := c.CheckingAccount{HolderName: "Silvia", BranchCode: 123, AccountNumber: 9838, Balance: 130.5}
	fmt.Printf("Initial balance(%s) -> %.2f\n", firstAccount.HolderName, firstAccount.Balance)
	fmt.Printf("Initial balance(%s) -> %.2f\n", secondAccount.HolderName, secondAccount.Balance)

	fmt.Println("")
	fmt.Printf("(%s) Trying withdraw operation with 13...\n", firstAccount.HolderName)
	fmt.Println(firstAccount.Withdraw(13))
	fmt.Println("")
	fmt.Printf("(%s) Trying withdraw operation with 130...\n", firstAccount.HolderName)
	fmt.Println(firstAccount.Withdraw(130))

	fmt.Println("")
	fmt.Printf("(%s) Trying deposit operation with 130...\n", firstAccount.HolderName)
	fmt.Println(firstAccount.Deposit(130))

	fmt.Println("")
	fmt.Printf("Actual balance(%s) -> %.2f\n", secondAccount.HolderName, secondAccount.Balance)
	fmt.Printf("(%s) Trying transfer operation with 100 to (%s)...\n", secondAccount.HolderName,
		firstAccount.HolderName)
	fmt.Println(secondAccount.Transfer(100, &firstAccount))
}
