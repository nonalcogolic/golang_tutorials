package main

import (
	"fmt"
)

func printNumber(n int) {

	switch {
	case n%3 == 0 && n%5 == 0:
		fmt.Println("FizzBuzz")
	case n%3 == 0:
		fmt.Println("Fizz")
	case n%5 == 0:
		fmt.Println("Buzz")
	default:
		fmt.Println(n)
	}
}

func main() {

	for i := 0; i < 100; i++ {
		printNumber(i)
	}

}
