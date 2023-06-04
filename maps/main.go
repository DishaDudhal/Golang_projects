package main

import "fmt"

func main() {
	//Ways of declaring a map
	var colors1 map[string]string

	colors2 := make(map[string]string)

	colors := map[string]string{
		"red":  "#ff0000",
		"blue": "#123345",
	}

	fmt.Println(colors)
}
