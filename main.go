package main

import "fmt"

// func main() {
// 	// var card string = "Hello World Variables in GO"
// 	card := 123 // "Hello World Variables in GO"
// 	fmt.Println(card)
// 	card = 456 // "Hello World Variables in GO after modified"
// 	fmt.Println(card)
// }

func main() {

	cards := deck{cardName(), cardNumber()}
	cards = append(cards, "CSV: 123")
	cards.print()
	fmt.Println(cards)
}

func cardName() string {
	return "Master Card"
}

func cardNumber() string {
	return "1000236547"
}
