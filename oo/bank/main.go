package main

import "fmt"

type CheckingAccount struct {
	titular       string
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

func main() {
	ccAccount := CheckingAccount{
		"Breno",
		234,
		123456,
		500,
	}

	fmt.Println(ccAccount.deposit(5000.))
}
