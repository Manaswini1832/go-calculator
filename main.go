package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*
		READING OPERANDS
	*/
	// operand 1
	fmt.Print("First number: ")
	reader := bufio.NewReader(os.Stdin)
	sop1, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading first operand:", err)
		return
	}

	// operand 2
	fmt.Print("Second number: ")
	sop2, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading second operand:", err)
		return
	}

	// Remove the newline character from the input
	sop1 = strings.TrimSpace(sop1)
	sop2 = strings.TrimSpace(sop2)

	op1, err := strconv.Atoi(sop1)
	if err != nil {
		panic(err) // Handle error appropriately in production code
	}

	op2, err := strconv.Atoi(sop2)
	if err != nil {
		panic(err) // Handle error appropriately in production code
	}

	/*
		SELECTING THE OPERATOR
	*/
	fmt.Print("Choose an operation: +, -, *, /\n")
	operatorString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading operator:", err)
		return
	}

	// Trim whitespace from the operator input
	operatorString = strings.TrimSpace(operatorString)

	switch operatorString {
	case "+":
		fmt.Printf("Sum = %v\n", op1+op2)
	case "-":
		fmt.Printf("Difference = %v\n", op1-op2)
	case "*":
		fmt.Printf("Product = %v\n", op1*op2)
	case "/":
		if op2 != 0 {
			fmt.Printf("Quotient = %v\n", op1/op2)
			fmt.Printf("Remainder = %v\n", op1%op2)
		} else {
			fmt.Println("Error: Cannot divide by zero.")
		}
	default:
		fmt.Println("Invalid operator")
	}
}
