package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(target)
	fmt.Println("Please guess the number I've generated")
	for i := 0; i < 10; i++ {
		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		userInput = strings.TrimSpace(userInput)
		userGuess, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please entered a valid number", err)
		}

		if userGuess < target {
			fmt.Println("Oops. Your guess was LOW")
		} else if userGuess > target {
			fmt.Println("Oops. Your guess was HIGH")
		} else {
			fmt.Println("Good job! You guessed it!")
			os.Exit(0)
		}
	}

	fmt.Println("Sorry, you didn't guess my number, it was: ", target)

}
