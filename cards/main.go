package main

func main() {
	//###Initialization
	//var card string = "Ace of spades"

	//Or
	//card := "Ace of spades"
	//card = "Five of Diamonds"

	//card := newCard()
	//fmt.Println(card)

	//calling other file in same package 'main'
	//printState()

	//Array vs slicing (dynamic array for same type of datatype)
	//cards := deck{"Five of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades")

	//call other file function
	//cards.print()

	//loop array
	//for i, card := range cards {
	//	fmt.Println(i, card)
	//}

	//create newdeck object
	//cards := newDeck()
	//hand, remainingCards := deal(cards, 3)
	//hand.print()
	//remainingCards.print()

	//g := "hi"
	//fmt.Println([]byte(g))

	//fmt.Println(cards.toString())

	//cards.saveToFile("myCards")

	cards := newDeckFromFile("myCards")
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
