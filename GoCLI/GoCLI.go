package main

import (
	"fmt"
)

func main() {
	var userName string

	for {
		fmt.Print("What is your name? ")
		fmt.Scan(&userName)
		if userName != "steve" {
			if userName == "Steven" {
				fmt.Println("OoOoOh I'm fancy my name is", userName)
			}
			fmt.Println("Get outta here", userName, "!!!")
		} else {
			fmt.Println("Welcome", userName)
		}
	}
}
