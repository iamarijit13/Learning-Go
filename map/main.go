package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff000",
		"black": "#000000",
		"white": "#ffffff",
	}

	// var colors map[string]string

	// colors := make(map[string]string)

	// colors["white"] = "#ffffff"
	// colors["black"] = "#000000"

	// fmt.Println(colors)

	// delete(colors, "white")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("The hex code of %v is %v.\n", color, hex)
	}
}
