package main

import "fmt"

type CheckingAccount struct {
	owner         string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func (a *CheckingAccount) withdraw(value float64) string {
	canWithdraw := value > 0 && value <= a.balance

	if canWithdraw {
		a.balance -= value

		return "Value withdrawed successfully!"
	} else {
		return "Insuficient funds!"
	}

}

func (a *CheckingAccount) deposit(amount float64) (string, float64) {
	if amount > 0 {
		a.balance += amount

		return "Amount deposited successfully", a.balance
	} else {
		return "Invalid deposit amount", a.balance
	}
}

func (a *CheckingAccount) transfer(amount float64, destinationAccount *CheckingAccount) bool {
	if amount > 0 && amount < a.balance {
		a.balance -= amount
		destinationAccount.deposit(amount)

		return true
	}

	return false
}

func main() {
	ccAccount := CheckingAccount{
		owner:   "Breno",
		balance: 300,
	}

	ccAccount2 := CheckingAccount{
		owner:   "Teste2",
		balance: 100,
	}

	status := ccAccount.transfer(200., &ccAccount2)
	fmt.Println(status)
	fmt.Println(ccAccount, ccAccount2)
}
