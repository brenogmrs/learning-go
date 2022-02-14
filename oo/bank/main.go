package main

import "fmt"

type CheckingAccount struct {
	titular 	  string
	agencyNumber  int
	accountNumber int
	balance 	  float64
}

func (a *CheckingAccount) withdraw(value float64) string {
	canWithdraw := value > 0 && value <= a.balance

	if canWithdraw {
		a.balance -= value

		return "Value withdrawed successfully!"
	} else {
		return "Insuficient funds"
	}

}


func main() {
	ccAccount := CheckingAccount{
		"Breno",
		234,
		123456,
		500,
	}

	withdrawValue := -100.
	
	fmt.Println(ccAccount.withdraw(withdrawValue))
	fmt.Println(ccAccount)
}
