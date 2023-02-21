package main

func main() {
	// var card string = "Ace of Spades"   //You are telling Go compiler whats the type
	cards := newDeck()
	//cards.print()
	hand, _ := deal(cards, 5)
	hand.print()
	//remainingCards.print()
	cards.print()
}
