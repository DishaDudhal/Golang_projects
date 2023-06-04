package main

import "fmt"

func main() {
	//Adding and deleting from a Map

	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#123345",
		"green": "#098098",
		"black": "#676576",
	}
	//Add an element to map
	colors["white"] = "#345345"

	//Delete an element from map
	delete(colors, "blue")

	//Iterate over the map
	for key, val := range colors {
		fmt.Println("Hex code for color", key, "is :", val)
	}

	// fmt.Println(colors)
}
