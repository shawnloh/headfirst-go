package main

import (
	"fmt"
	"github.com/shawnloh/headfirstgo/chapter-8/magazine"
)

func main() {
	//subscriber := magazine.Subscriber{Name: "Aman, Singh", Rate:4.99, Active:true}
	//fmt.Println(subscriber)
	employee := magazine.Employee{Name: "Joy Carr"}
	employee.Address.Street = "456 Kim St."
	employee.Address.City = "Portland"
	employee.Address.State = "OR"
	employee.Address.PostalCode = "97222"

	fmt.Printf("%#v", employee)

}
