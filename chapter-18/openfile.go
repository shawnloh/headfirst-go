package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("aardvark.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.FileMode(0600))
	check(err)
	defer file.Close()
	_, err = file.Write([]byte("amazing!\n"))
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
