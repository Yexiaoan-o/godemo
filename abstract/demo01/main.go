package main
import (
	"fmt"
)

// Define an Account struct
type Account struct {
	Username string
	Password string
	Balance float64
}

func (account *Account) Deposit (money float64, password string){
	if (password == account.Password){
		account.Balance += money
		fmt.Println("Current balance is", account.Balance)
	} else {
		fmt.Println("Password incorrect")
	}
}

func main(){
	bank1 := Account{ }
	bank1.Username = "yexiaoan"
	bank1.Password = "1@gotest1"
	bank1.Balance = 167079.93

	bank1.Deposit(10000, "1@gotest")

	fmt.Printf("Account name is %v, Password is %v, and Balance is %v", bank1.Username, bank1.Password, bank1.Balance)
}