package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

const fileName = "balance.txt"

func main() {

	var option int
	balance, err := readBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------------------")
		//panic("Sorry can not continue")
	}


	fmt.Println("Welcome to Go Banking System")

	//infinite loop made with an empty for
	for  {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		fmt.Print("Please select an option: ")
		fmt.Scanln(&option)
		fmt.Println("You selected option ", option)

		switch option {
			case 1:
				fmt.Println("Your balance is ", balance)
			case 2:
				fmt.Println("How much do you want to deposit?")
				var deposit float64
				fmt.Scanln(&deposit)
				if deposit <= 0 {
					fmt.Println("Invalid amount")
					continue
				}
				balance += deposit
				writeBalanceToFile(balance)
				fmt.Println("Your new balance is ", balance)
			case 3:
				fmt.Println("How much do you want to withdraw?")
				var withdraw float64
				fmt.Scanln(&withdraw)
				if withdraw <= 0 {
					fmt.Println("Invalid amount")
					continue
				}
				if withdraw > balance {
					fmt.Println("You don't have enough funds")
					continue
				} 
				balance -= withdraw
				writeBalanceToFile(balance)
				fmt.Println("Your new balance is ", balance)	
			default:
				fmt.Println("Exiting...!")
				fmt.Println("Thank you for using the Go Banking System. Goodbye!")
				return //return for switch
				//break for if
		}
	
	}

}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 10000, errors.New("Error reading file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 10000, errors.New("Error parsing balance")
	}
	return balance, nil
}