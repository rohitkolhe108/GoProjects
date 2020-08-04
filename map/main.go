package main

import "fmt"

func main() {
	colors := map[string]string{
		"a": "1",
		"b": "2",
	}

	printmap(colors)
}

func printmap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for ", color, " is ", hex)
	}
}
