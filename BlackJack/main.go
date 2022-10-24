package main

import (

	"fmt"

	"blackjack/game"

)

//Loop for multiple rounds
func main() {
    var play int 

	fmt.Println("------ WELCOME TO BLACKJACK ------")
	fmt.Println("[0]:  BET & PLAY | [1]: EXIT")
	fmt.Scanln(&play) 

	game.CreateGame(1000, 10)
	for play == 0 {
		game.PlayGame()

		fmt.Println("\n\nKeep Playing?")
		fmt.Println("[0]: BET & PLAY | [1]: EXIT")
		fmt.Scanln(&play)

		game.ResetGame()
	}
}