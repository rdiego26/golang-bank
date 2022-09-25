package main

import (
	"fmt"
	"golang-bank/accounts"
	"golang-bank/customers"
)

func namedParameters() {
	var name = "Ramos"
	var branchCode = 589
	var accountNumber = 123456

	newAccount := accounts.CheckingAccount{AccountNumber: accountNumber}

	fmt.Println(newAccount)
	fmt.Println(accounts.CheckingAccount{Owner: customers.AccountOwner{
		Name: name, NationalID: "000.000.000-31", Occupation: "Accountant"},
	}, accountNumber, branchCode)
}

func comparisonTypes() {
	var crisAccount *accounts.CheckingAccount
	crisAccount = new(accounts.CheckingAccount)
	crisAccount.Owner.Name = "Cris"
	crisAccount.Deposit(500)

	var crisAccount2 *accounts.CheckingAccount
	crisAccount2 = new(accounts.CheckingAccount)
	crisAccount2.Owner.Name = "Cris"
	crisAccount2.Deposit(500)

	fmt.Println(*crisAccount)

	// Comparing references(not values)
	fmt.Println("crisAccount == crisAccount2 ->", crisAccount == crisAccount2)

	// Comparing values
	fmt.Println("*crisAccount == *crisAccount2 ->", *crisAccount == *crisAccount2)

}

func main() {
	firstCustomer := customers.AccountOwner{Name: "Ramos", NationalID: "000.000.000-30", Occupation: "Developer"}
	firstAccount := accounts.CheckingAccount{Owner: firstCustomer, BranchCode: 123, AccountNumber: 9837}
	firstAccount.Deposit(30.5)

	secondCustomer := customers.AccountOwner{Name: "Silvia", NationalID: "000.000.000-20", Occupation: "Developer"}
	secondAccount := accounts.CheckingAccount{Owner: secondCustomer, BranchCode: 123, AccountNumber: 9838}
	secondAccount.Deposit(130.5)

	fmt.Printf("Initial balance(%s) -> %.2f\n", firstAccount.Owner.Name, firstAccount.GetBalance)
	fmt.Printf("Initial balance(%s) -> %.2f\n", secondAccount.Owner.Name, secondAccount.GetBalance)

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
	fmt.Printf("Actual balance(%s) -> %.2f\n", secondAccount.Owner.Name, secondAccount.GetBalance)
	fmt.Printf("(%s) Trying transfer operation with 100 to (%s)...\n", secondAccount.Owner.Name,
		firstAccount.Owner.Name)
	fmt.Println(secondAccount.Transfer(100, &firstAccount))
}
