package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	//Writing the firstname and lastname : is optional, however if not specified the order within
	//the curly braces will be mapped to person struct aka {"Alex", "Anderson"} by default Alex is firstName
	//and Anderson is lastname
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}
