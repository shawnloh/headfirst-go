package main

import (
	"github.com/shawnloh/headfirstgo/chapter-4/keyboard"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"log"
)

func main() {
	fmt.Print("Enter a temperatures in fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees celsius\n", celsius)
}
