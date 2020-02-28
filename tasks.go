package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatal("Incorrect number of params")
	}

	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatal("Incorrect A param")
	}

	b, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		log.Fatal("Incorrect B param")
	}

	c, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		log.Fatal("Incorrect C param")
	}

	p := (a + b + c) / 2

	if p < a || p < b || p < c {
		log.Fatal("Invalid triangle")
	}

	S := math.Sqrt(p * (p - a) * (p - b) * (p - c))

	fmt.Println("Result", S)
}
