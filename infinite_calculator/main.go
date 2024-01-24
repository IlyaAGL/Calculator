package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"

	"github.com/shopspring/decimal"
)

func main() {
	var res decimal.Decimal
	var op string

	fmt.Println("Actul result is ", res)
	fmt.Print("Enter first operation: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		op = scanner.Text()
	}

	// For catching several operations
	var operation string;
	var ind = -1;
	
	// Finding operation sign
	for i := 0; i < len(op); i++ {
		_, err := strconv.ParseInt(string(op[i]), 10, 64)
		if (err != nil && op[i] != ' ') {
			if ind != -1 {
				log.Fatal("Too many operations")
			}
			operation = string(op[i]);
			ind = i;
		}
	}

	if ind == -1 {
		log.Fatal("No opearations detected")
	}

	firstNumber, err := decimal.NewFromString(strings.Trim(op[0 : ind], " "))
	if err != nil {
		log.Fatal("Not a number")
	}

	secondNumber, err := decimal.NewFromString(strings.Trim(op[(ind + 1):], " "))
	if err != nil {
		log.Fatal("Not a number")
	}


	switch operation {
	case "+":
		res = firstNumber.Add(secondNumber)
	case "-":
		res = firstNumber.Sub(secondNumber)
	case "*":
		res = firstNumber.Mul(secondNumber)
	case "/":
		if secondNumber.Cmp(decimal.NewFromInt(0)) == 0 {
			log.Fatal("Cannot divide by zero!")
		} else {
			res = firstNumber.Div(secondNumber)
		}
	default:
		log.Fatal("Try average operations!")
	}

	for {
		fmt.Println("Next actual result: ", res)

		var next string

		fmt.Print("Enter next operaton or type '!' to exit: ", res, " ")

		nextScanner := bufio.NewScanner(os.Stdin)
		if nextScanner.Scan() {
			next = nextScanner.Text()
		}

		if next == "!" {
			os.Exit(1)
		}

		next = strings.ReplaceAll(next, " ", "")
		nextOp := string(next[0])
		parts := strings.Split(next, string(next[0]))
		nextNum, err := decimal.NewFromString(parts[1])
		if err != nil {
			log.Fatal("Not a number!")
		}

		switch nextOp {
		case "+":
			res = res.Add(nextNum)
		case "-":
			res = res.Sub(nextNum)
		case "*":
			res = res.Mul(nextNum)
		case "/":
			if decimal.Zero.Cmp(nextNum) == 0 {
				log.Fatal("Cannot divide by zero!")
			} else {
				res = res.Div(nextNum)
			}
		default:
			log.Fatal("Try average operations!")
		}

	}
}
