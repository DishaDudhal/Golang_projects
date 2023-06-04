package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//This is an example of how to use struct inside a struct. Note that we can also use contactInfo
	//struct freely without having to create any person struct.
	jim := person{
		firstName: "Jim",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "jima@gmail.com",
			zipCode: 12345,
		},
	}

	fmt.Printf("%+v", jim)
}
