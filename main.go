package main

import (
	"fmt"
)

// func main() {
// 	// var card string = "Hello World Variables in GO"
// 	card := 123 // "Hello World Variables in GO"
// 	fmt.Println(card)
// 	card = 456 // "Hello World Variables in GO after modified"
// 	fmt.Println(card)
// }

func add(x, y int) int {
	return x + y
}

func main() {

	// cards := deck{cardName(), cardNumber()}
	// cards = append(cards, "CSV: 123")
	// cards.print()
	// fmt.Println(cards)
	forFace1()
	forFace2()

}

func forFace1() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func forFace2() {
	sum := 1
	for sum < 100 {
		sum += sum
		if sum == 64 {
			fmt.Println("if success", sum)
			break
		} else {
			fmt.Println("if failed", sum)
		}
	}
	fmt.Println(sum)
}

func cardName() string {
	return "Master Card"
}

func cardNumber() string {
	return "1000236547"
}
