package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(sub subscriber) {
	fmt.Println("Name:", sub.name)
	fmt.Println("Monthly rate:", sub.rate)
	fmt.Println("Active?", sub.active)
}

func defaultSubscriber(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return s
}
func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func main() {
	subscriber1 := defaultSubscriber("Shawn Loh")
	subscriber1.rate = 4.95
	printInfo(subscriber1)
	subscriber2 := defaultSubscriber("Ryan")
	printInfo(subscriber2)
	applyDiscount(&subscriber2)
	printInfo(subscriber2)
}
