package main

func main() {
	cards := newDeck()
	cards.shuffle()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	cards.print()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
