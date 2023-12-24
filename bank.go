package main

import (
	"fmt"
	"github.com/thegera4/dummy_banking/file_operations" //add the path in your .mod file and then the name of your package
	"github.com/Pallinder/go-randomdata" // 3rd party package
)

const fileName = "balance.txt"

func main() {

	var option int
	balance, err := file_operations.GetFloatFromFile(fileName, 10000)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------------------")
		//panic("Sorry can not continue")
	}

	fmt.Println("Welcome to Go Banking System")
	fmt.Println("Reach us 24/7 at: ", randomdata.PhoneNumber())

	//infinite loop made with an empty for
	for  {
		
		presentMenu()

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
				file_operations.WriteFloatToFile(fileName, balance)
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
				file_operations.WriteFloatToFile(fileName, balance)
				fmt.Println("Your new balance is ", balance)	
			default:
				fmt.Println("Exiting...!")
				fmt.Println("Thank you for using the Go Banking System. Goodbye!")
				return //return for switch
				//break for if
		}
	
	}

}