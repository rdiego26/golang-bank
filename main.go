package main

import (
	"fmt"
	c "golang-bank/accounts"
	"golang-bank/customers"
)

func namedParameters() {
	var name = "Ramos"
	var branchCode = 589
	var accountNumber = 123456
	var balance = 125.50

	newAccount := c.CheckingAccount{Balance: 100.00}

	fmt.Println(newAccount)
	fmt.Println(c.CheckingAccount{Owner: customers.AccountOwner{
		Name: name, NationalID: "000.000.000-31", Occupation: "Accountant"},
	}, accountNumber, branchCode, balance)
}

func comparisonTypes() {
	var crisAccount *c.CheckingAccount
	crisAccount = new(c.CheckingAccount)
	crisAccount.Owner.Name = "Cris"
	crisAccount.Balance = 500

	var crisAccount2 *c.CheckingAccount
	crisAccount2 = new(c.CheckingAccount)
	crisAccount2.Owner.Name = "Cris"
	crisAccount2.Balance = 500

	fmt.Println(*crisAccount)

	// Comparing references(not values)
	fmt.Println("crisAccount == crisAccount2 ->", crisAccount == crisAccount2)

	// Comparing values
	fmt.Println("*crisAccount == *crisAccount2 ->", *crisAccount == *crisAccount2)

}

func main() {
	firstCustomer := customers.AccountOwner{Name: "Ramos", NationalID: "000.000.000-30", Occupation: "Developer"}
	firstAccount := c.CheckingAccount{Owner: firstCustomer, BranchCode: 123, AccountNumber: 9837, Balance: 30.5}

	secondCustomer := customers.AccountOwner{Name: "Silvia", NationalID: "000.000.000-20", Occupation: "Developer"}
	secondAccount := c.CheckingAccount{Owner: secondCustomer, BranchCode: 123, AccountNumber: 9838, Balance: 130.5}

	fmt.Printf("Initial balance(%s) -> %.2f\n", firstAccount.Owner.Name, firstAccount.Balance)
	fmt.Printf("Initial balance(%s) -> %.2f\n", secondAccount.Owner.Name, secondAccount.Balance)

	fmt.Println("")
	fmt.Printf("(%s) Trying withdraw operation with 13...\n", firstAccount.Owner.Name)
	fmt.Println(firstAccount.Withdraw(13))
	fmt.Println("")
	fmt.Printf("(%s) Trying withdraw operation with 130...\n", firstAccount.Owner.Name)
	fmt.Println(firstAccount.Withdraw(130))

	fmt.Println("")
	fmt.Printf("(%s) Trying deposit operation with 130...\n", firstAccount.Owner.Name)
	fmt.Println(firstAccount.Deposit(130))

	fmt.Println("")
	fmt.Printf("Actual balance(%s) -> %.2f\n", secondAccount.Owner.Name, secondAccount.Balance)
	fmt.Printf("(%s) Trying transfer operation with 100 to (%s)...\n", secondAccount.Owner.Name,
		firstAccount.Owner.Name)
	fmt.Println(secondAccount.Transfer(100, &firstAccount))
}
