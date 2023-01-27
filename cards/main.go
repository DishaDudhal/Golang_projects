package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"   //You are telling Go compiler whats the type
	card := "Ace of Spades" // the go compiler needs to figure out the type of variable
	// := is only used while first initialization, after than you dont
	//have to do ":=" you can just use "="
	fmt.Println(card)
}

func newCard() {
	return "Five of Diamonds"
}
