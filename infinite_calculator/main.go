package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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

	firstNumber, err := decimal.NewFromString(strings.Split(op, " ")[0])
	if err != nil {
		log.Fatal("Not a number!")
	}

	secondNumber, err := decimal.NewFromString(strings.Split(op, " ")[2])
	if err != nil {
		log.Fatal("Not a number!")
	}

	operation := strings.Split(op, " ")[1]

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

		if next == "over" {
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
	//fmt.Printf("Type is %T", firstNumber)
}
