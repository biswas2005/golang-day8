package pointers

import "fmt"

type BankAccount struct {
	Name    string
	Balance float64
}

func (a *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		fmt.Printf("\nAmount of %.3f has been deposited to %s's account.", amount, a.Name)
	}

}

func (a *BankAccount) Withdraw(amount float64) {
	if amount > 0 && amount <= a.Balance {
		fmt.Printf("\n%.2f has be deducted from %s's account.", amount, a.Name)
	} else {
		fmt.Println("Insufficient Balance.")
		fmt.Printf("\nAvailable balance %.3f", a.Balance)
	}
}

func (a *BankAccount) Display() {
	fmt.Printf("Account Holder:%s Balance:%.3f", a.Name, a.Balance)
}

func BankAcc() {
	account := BankAccount{Name: "Biswas", Balance: 2000.0}

	account.Display()

	account.Deposit(500)
	account.Withdraw(1000)
}
