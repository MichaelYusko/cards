package main



func main()  {
	cards := deck{"Ace of diamonds", newCard()}
	cards.print()
}

func newCard() string{
	return "Five of diamonds"
}