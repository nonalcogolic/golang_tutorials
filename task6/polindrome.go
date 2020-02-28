package main

import (
	"fmt"
	"log"
	"os"
)

func existMirror(s string, length int) bool {

	fmt.Println(s, len(s), length)
	result := true
	for start, rstart := 0, length*2-1; start < rstart; {
		fmt.Println(s[start], s[rstart])
		if s[start] != s[rstart] {
			result = false
			break
		}
		start++
		rstart--
	}

	return result
}

func main() {

	if len(os.Args) != 2 {
		log.Fatal("Incorrect number of params")
	}

	s := os.Args[1]

	maxLen := 0

	for i := 0; i < len(s); i++ {
		length := 0
		testString := s[i:]
		for j := 1; len(testString) >= j*2; j++ {
			if existMirror(testString, j) {
				length = j
			}
		}
		if maxLen < length {
			maxLen = length
		}
	}

	fmt.Println(maxLen)

}
