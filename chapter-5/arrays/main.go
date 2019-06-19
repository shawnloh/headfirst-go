package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}

	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}

	for i, value := range notes {
		fmt.Println(i, value)
	}
}
