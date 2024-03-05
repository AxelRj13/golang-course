package main

import "fmt"

func main() {
	/**
	* CREATING MAP SYNTAX
	* map[x]y
	* x => the data type of key
	* y => the data type of value
	* NOTE: all of keys / values must be in the same type each. So if the key is string, then all of the rest of keys must be string. Same goes for values.
	* But the key and value doesn't have to be the same type, e.g. the key is string and the value is int, this one's allowed
	 */

	// 1st way to declare map
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// declare map 2
	// var colors map[string]string

	// declare map 3
	// colors := make(map[int]string)
	// colors[10] = "#ffffff" // assign a new value to map
	// delete(colors, "green") // delete a key in map

	printMap(colors)
	fmt.Println(colors)
}

// iterate over a map
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
