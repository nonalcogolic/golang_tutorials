package main

import (
	"fmt"
	"math"
)

func getPolinomeFN(values []float64) func(float64) float64 {
	maxLen := len(values)

	functionSlice := make([]func(float64) float64, 0, maxLen)

	fmt.Println("Here functionSlice")
	for i := 0; i < maxLen; i++ {
		a := values[i]
		power := maxLen - (i + 1)
		currentFunction := func(x float64) float64 {
			return a * math.Pow(x, float64(power))
		}

		fmt.Println(currentFunction(1), "===><===")
		functionSlice = append(functionSlice, currentFunction)
	}

	fmt.Println("Here returnFunction")

	returnFunction := func(x float64) float64 {
		fmt.Println(len(functionSlice))
		fmt.Println(functionSlice)
		summ := .0
		for _, fn := range functionSlice {
			summ += fn(x)
		}
		return summ
	}

	return returnFunction
}

func main() {

	var slise_a []float64 = []float64{1, 2, 3}
	function := getPolinomeFN(slise_a)
	result := function(1)
	fmt.Println(result)
	fmt.Println("Finish")

}
