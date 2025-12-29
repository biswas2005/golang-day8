package pointers

import "fmt"

type BankAccount struct {
	Name    string
	Balance float64
}

//Deposit() adds an amount to Bank Account
func (a *BankAccount) Deposit() {

	var amount float64
	fmt.Println("\nEnter an amount to Deposit:")
	fmt.Scan(&amount)

	if amount > 0 {
		fmt.Printf("\nAmount of %.3f has been deposited to %s's account.", amount, a.Name)
		updatedBalance := a.Balance + amount
		fmt.Printf("\nUpdated Balance:%f", updatedBalance)
	} else {
		fmt.Print("\nEnter valid amount.")
	}

}

//withdraw() an amount from account
func (a *BankAccount) Withdraw() {

	var amount float64
	fmt.Println("\nEnter an amount to withdraw:")
	fmt.Scan(&amount)
	if amount > 0 && amount <= a.Balance {
		fmt.Printf("\n%.2f has be deducted from %s's account.", amount, a.Name)
		updatedBalance := a.Balance - amount
		fmt.Printf("\nUpdated Balance:%f", updatedBalance)
	} else {
		fmt.Println("Insufficient Balance.")
		fmt.Printf("\nAvailable balance %.3f", a.Balance)
	}
}

//Display() the Account details
func (a *BankAccount) Display() {
	fmt.Printf("Account Holder:%s Balance:%.3f", a.Name, a.Balance)
}

func BankAcc() {
	for {
		var n int
		//prints a menu to choose action
		fmt.Println("\n<Choose an option>")
		fmt.Println("1.Diposit Money.")
		fmt.Println("2.Withdraw Money.")
		fmt.Println("3.Dispaly Details.")
		fmt.Println("4.EXIT.")
		fmt.Println("Enter your choice:")
		fmt.Scan(&n)
		account := BankAccount{Name: "Biswas\n", Balance: 2000.0}

		//switch() between cases
		switch n {
		case 1:
			account.Deposit()
		case 2:
			account.Withdraw()
		case 3:

			account.Display()
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid Choice!")

		}

	}

}
