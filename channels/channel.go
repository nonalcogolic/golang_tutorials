package main

import "fmt"

var a int = 0

func f(c chan int) {
	a := 0
	for i := 0; i < 10_000; i++ {
		a++
	}
	c <- a
}

func main() {
	c := make(chan int, 100)
	for i := 0; i < 100; i++ {
		go f(c)
	}

	for i := 0; i < 100; i++ {
		a += <-c
	}

	fmt.Println(a)
}
