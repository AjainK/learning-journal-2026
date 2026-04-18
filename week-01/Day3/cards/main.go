package main

func main() {

	cards := newDeckFromFile("my_cards.txt")
	cards.shuffle()
	//hand, remainingCards := deal(cards, 3)
	//hand.print()
	//remainingCards.print()
	//cards.saveTofile("my_cards.txt")
	cards.print()
}
