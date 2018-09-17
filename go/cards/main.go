package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	//--Read write from file--
	//cards.saveToFile("my_cards")
	//cards := newDeckFromFile("my_cards")

}
