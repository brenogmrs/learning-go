package accounts

type CheckingAccount struct {
	Owner         string
	AgencyNumber  int
	AccountNumber int
	Balance       float64
}

func (a *CheckingAccount) Withdraw(amount float64) string {
	canWithdraw := amount > 0 && amount <= a.Balance

	if canWithdraw {
		a.Balance -= amount

		return "Value withdrawed successfully!"
	} else {
		return "Insuficient funds!"
	}

}

func (a *CheckingAccount) Deposit(amount float64) (string, float64) {
	if amount > 0 {
		a.Balance += amount

		return "Amount deposited successfully", a.Balance
	} else {
		return "Invalid deposit amount", a.Balance
	}
}

func (a *CheckingAccount) Transfer(amount float64, destinationAccount *CheckingAccount) bool {
	if amount > 0 && amount < a.Balance {
		a.Balance -= amount
		destinationAccount.Deposit(amount)

		return true
	}

	return false
}
