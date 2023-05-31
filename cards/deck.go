package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

// (d deck) refers to the receiver of the function print, aka any receiver "d" of type deck
// will be able to call the function print() as -> d.print()
// If compared to an object oriented language it is sort of similar to an object's method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(card, i)
	}
}

// function deal will return 2 values both of type deck
// Here we dont have any receiver but we just have a function deal which takes the deck as an argument
// along with the size of the hand to be dealt. It also specifis the return type of the function - which
// is 2 decks.
// This particular design choice was made to make it less confusing about what deal actually does.
// aka d.deal(5) does not convey proper info as to what is 5 or what are we dealing.
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

func (d deck) shuffle() {
	//Go rand always uses same seed to randomize, so we need to figure out how to randomize the seed
	//First we need to create a new source object using rand.NewSource function which takes a seed value in
	//in the form of int64. Then we use this source object to create a new rand object called r.
	//To create a unique int64 seed everytime, we use the time package, Now function which gives current time
	//and then apply UnixNano function on it which returns an int64 value of current time.
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		pos := r.Intn(len(d) - 1)   //find a random card pos and swap current card with that card
		d[i], d[pos] = d[pos], d[i] //swap
	}
}
