package main


func main()  {
	cards := newDeck()
	cards.shuffle()
	cards.saveToFile("my_cards")
	cards = newDeckFromFile("my_cards")
	cards.print()
}