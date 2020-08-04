package main

import "fmt"

func main() {

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Printf("%d", number)
			fmt.Println(" is even")
		} else {
			fmt.Printf("%d", number)
			fmt.Println(" is odd")
		}

	}
}
