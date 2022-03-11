package accounts

import (
	"github.com/oo/bank/customers"
)

type CheckingAccount struct {
	Owner                       customers.AccountOwner
	AgencyNumber, AccountNumber int
	balance                     float64
}

func (a *CheckingAccount) Withdraw(amount float64) string {
	canWithdraw := amount > 0 && amount <= a.balance

	if canWithdraw {
		a.balance -= amount

		return "Value withdrawed successfully!"
	} else {
		return "Insuficient funds!"
	}

}

func (a *CheckingAccount) Deposit(amount float64) (string, float64) {
	if amount > 0 {
		a.balance += amount

		return "Amount deposited successfully", a.balance
	} else {
		return "Invalid deposit amount", a.balance
	}
}

func (a *CheckingAccount) Transfer(amount float64, destinationAccount *CheckingAccount) bool {
	if amount > 0 && amount < a.balance {
		a.balance -= amount
		destinationAccount.Deposit(amount)

		return true
	}

	return false
}

func (a *CheckingAccount) GetBalance() float64 {
	return a.balance
}
