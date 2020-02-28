package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func split(text string) map[string]int {
	vocabluary := make(map[string]int)
	match, _ := regexp.MatchString("[A-Za-z]+", text)
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))

	for _, value := range match {
		vocabluary[value.toUpper()]++
	}

	return make(map[string]int)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Invalid param")
	}

	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

}
