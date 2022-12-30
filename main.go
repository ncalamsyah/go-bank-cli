package main

import (
	"fmt"
)

func main() {
	var choice int
	fmt.Println("_< Welcome to banking portal terminal app >_")
	fmt.Println("0: Login\n1: Create Account")

	fmt.Print(">> ")
	fmt.Scanln(&choice)

	if choice == 0 {
		login()
	} else if choice == 1 {
		register()
	} else {
		fmt.Println("! > Invalid option!")
		login()
	}
}
