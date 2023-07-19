package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

func main() {
	add := flag.Bool("add", false, "adding two numbers")
	subs := flag.Bool("subtruct", false, "substract two numbers")
	mult := flag.Bool("multiply", false, "multiple two numbers")
	div := flag.Bool("divide", false, "divide two numbers")
	pow := flag.Bool("pow", false, "first number power to second")
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
		//os.Exit(1)
	}

	var num1, num2 float64
	fmt.Println("Enter 1st Number: ")
	fmt.Scan(&num1)
	fmt.Println("Enter 2nd Number: ")
	fmt.Scan(&num2)

	switch {
	case *add:
		fmt.Println("Result is", num1+num2)
	case *subs:
		fmt.Println("Result is", num1-num2)
	case *mult:
		fmt.Println("Result is", num1*num2)
	case *div:
		if num2 == 0 {
			fmt.Println("Can't divide by 0, try again")
			os.Exit(1)
		} else {
			fmt.Println("Result is", num1/num2)
		}
	case *pow:
		fmt.Println("Result is", math.Pow(num1, num2))
	default:
		fmt.Fprintln(os.Stderr, "Wrong option try with add, subtract, div and multply")
		os.Exit(1)
	}
}
