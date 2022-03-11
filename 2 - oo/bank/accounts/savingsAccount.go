package accounts

import (
	"github.com/oo/bank/customers"
)

type SavingsAccount struct {
	Owner                                  customers.AccountOwner
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (a *SavingsAccount) Withdraw(amount float64) string {
	canWithdraw := amount > 0 && amount <= a.balance

	if canWithdraw {
		a.balance -= amount

		return "Value withdrawed successfully!"
	} else {
		return "Insuficient funds!"
	}

}

func (a *SavingsAccount) Deposit(amount float64) (string, float64) {
	if amount > 0 {
		a.balance += amount

		return "Amount deposited successfully", a.balance
	} else {
		return "Invalid deposit amount", a.balance
	}
}

func (a *SavingsAccount) GetBalance() float64 {
	return a.balance
}
