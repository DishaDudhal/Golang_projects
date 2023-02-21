package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create a new type(data-structure) of deck
// which is a slice of strings
type deck []string

// func funcName([arg: Type]) Returntype
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(card, i)
	}
}

// function deal will return 2 values both of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// function to convert string to byte string for easy writing to a file
func (d deck) toString() string {
	return strings.Join([]string(d), "\n")
}

// d deck is a receiver, 0666 is file permissions viz linux, filename is param to func, error is return object
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//if we were successful convert the byteSlice into string, then convert this string into
	// slice of strings using inbuilt function Split.
	sSlice := strings.Split(string(byteSlice), "\n")

	//Convert the stringSlice into type deck and return
	return deck(sSlice)
}
