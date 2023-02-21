package main

func main() {
	// var card string = "Ace of Spades"   //You are telling Go compiler whats the type
	//cards := newDeck()
	//cards.print()
	//hand, _ := deal(cards, 5)
	// fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	newcards := readDeckFromFile("my_cards1")
	newcards.print()
}
