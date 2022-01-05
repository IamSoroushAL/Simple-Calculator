package main

import "fmt"

func main() {
	var flag bool
	flag = false

	fmt.Println()
	fmt.Println("Welcome to simple calculator")
	fmt.Println("Please enter a number as input")
	fmt.Println("To add numbers use + operator")
	fmt.Println("To subtract numbers use - operator")
	fmt.Println("To multiply numbers use * operator")
	fmt.Println("To divide numbers use / operator")
	fmt.Println("To get output use = operator")
	fmt.Println("Hope you enjoy it :)")
	fmt.Println("Author : Soroush Ataee <ataeesoroush@gmail.com>")
	fmt.Println()

	var input_number float64 = 0
	var input_operator string
	var output float64 = 0

	counter := 0
	for flag == false {
		counter = counter + 1
		if input_operator != "=" {
			fmt.Println("Input : ")
			fmt.Scanln(&input_number)
		}
		if counter == 1 {
			output = input_number
		}
		if counter == 1 {
			fmt.Println("Operator : ")
			fmt.Scanln(&input_operator)
		}
		if counter == 1 && input_operator != "=" {
			fmt.Println("Input : ")
			fmt.Scanln(&input_number)
		}
		switch {
		case input_operator == "+":
			output = output + input_number
		case input_operator == "-":
			output = output - input_number
		case input_operator == "*":
			output = output * input_number
		case input_operator == "/":
			output = output / input_number
		case input_operator == "=":
			fmt.Println("Output :", output)
			fmt.Println("You can start and calculate again by entering input or you can close the program using 'Ctrl + c'")
			input_number = 0
			output = 0
			counter = 0
			input_operator = ""
		}
		if input_operator != "=" && counter != 0 {
			fmt.Println("Operator : ")
			fmt.Scanln(&input_operator)
		}
	}
}
