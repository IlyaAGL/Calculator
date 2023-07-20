package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	add := flag.Bool("add", false, "add two numbers")
	sub := flag.Bool("sub", false, "subtract two numbers")
	mult := flag.Bool("mult", false, "multiply two numbers")
	div := flag.Bool("div", false, "divide two numbers")
	ex := flag.Bool("ex", false, "exit the app")
	res := 0.0

	flag.Parse()

	var num1, num2 float64
	fmt.Println("Enter 1 number:")
	fmt.Scan(&num1)
	fmt.Println("Enter 2 number:")
	fmt.Scan(&num2)

	switch {
	case *add:
		res = num1 + num2
		fmt.Println("Result is ", res)
	case *sub:
		res = num1 - num2
		fmt.Println("Result is ", res)

	case *mult:
		res = num1 * num2
		fmt.Println("Result is ", res)
	case *div:
		if num2 == 0 {
			fmt.Println("Cannot divide by zero!")
			os.Exit(1)
		} else {
			res = num1 / num2
			fmt.Println("Result is ", res)
		}
	case *ex:
		os.Exit(1)
	default:
		fmt.Println("Try again!")
		os.Exit(1)
	}

	for {
		var num3 float64
		fmt.Println("Enter next number:")
		fmt.Scan(&num3)
		switch {
		case *add:
			res += num3
			fmt.Println("Result is ", res)
		case *sub:
			res -= num3
			fmt.Println("Result is ", res)

		case *mult:
			res *= num3
			fmt.Println("Result is ", res)
		case *div:
			if num3 == 0 {
				fmt.Println("Cannot divide by zero!")
				os.Exit(1)
			} else {
				res /= num3
				fmt.Println("Result is ", res)
			}
		case *ex:
			os.Exit(1)
		default:
			fmt.Println("Try again!")
			os.Exit(1)
		}
	}

}

//fmt.Printf("Type is %T", *add)
