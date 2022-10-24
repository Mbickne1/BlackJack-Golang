package game

import (
	"fmt"

	"blackjack/deck"

	"time"
)

type Suit int

type Card = deck.Card

const (
	Hearts Suit = iota
	Diamonds
	Clubs
	Spades
)

//Struct for Player/Dealer
type Player struct {
	hand      []Card
	handValue int
	isDealer  bool
}

//Betting Controls
type BetControl struct {
	playerChipTotal int
	minBet          int
	playerBet       int
}

// Game Controls
type GameControl struct {
	dealersPlay bool
	gameOver    bool
	gameResult  int
}

//Player & Dealer
var player Player
var dealer Player

var _betControl BetControl
var _gameControl GameControl

//Creates the Game Decks with desired 
func CreateGame(chipPool int, minimumBet int) {
	//Set Bet and Game Control Objects
	_betControl = BetControl{chipPool, minimumBet, 0}
	_gameControl = GameControl{false, false, -1}
	player.isDealer = false
	dealer.isDealer = true
	//Create, Extend, and Shuffle the Game Deck -- 3 Decks Combined
	deck.New([]int{11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10})
	deck.Extend()
	deck.Extend()

	deck.Shuffle()
}

//Reset the Game by Shuffling the Decks and Setting Hand Values to Defaults
func ResetGame() {
	deck.Shuffle()
	player.hand = make([]Card, 0)
	dealer.hand = make([]Card, 0)

	player.handValue = 0
	dealer.handValue = 0

	_gameControl = GameControl{false, false, -1}
}

//Main Gameplay Loop
func PlayGame() {
	//Game Opening
	fmt.Printf("Total Chip Pool is: %d | Minimum Bet is: %d\n", _betControl.playerChipTotal, _betControl.minBet)
	//Obtain Player Bet and Deal Initial Hands
	bet()
	deal()

	//Check for player or dealer blackjack
	playerBlackjack := player.isBlackJack()
	dealerBlackjack := dealer.isBlackJack()

	if playerBlackjack && dealerBlackjack {
		_gameControl.gameResult = 0
		_gameControl.gameOver = true
	} else if playerBlackjack {
		_gameControl.gameResult = 3
		_gameControl.gameOver = true
	} else if dealerBlackjack {
		_gameControl.gameResult = 4
		_gameControl.gameOver = true
	}
	
	for !_gameControl.gameOver {
		time.Sleep(1 * time.Second)

		var decision int
		//Print the Game UI
		fmt.Printf("------ BET: %d ------\n", _betControl.playerBet)

		dealer.state()
		player.state()

		fmt.Printf("------ PLAYER TOTAL: %d ------\n\n", _betControl.playerChipTotal)
		if !_gameControl.dealersPlay {
			//Get Player Choice
			fmt.Println("[0]: Hit | [1]: Stay")
			fmt.Scanln(&decision)

			if decision == 0 {
				//Give Player a New Card
				player.hit()
				//Check For Bust
				if player.isBust() {
					_gameControl.gameOver = true
					_gameControl.gameResult = 5
				}
				//Switch to dealer play
			} else if decision == 1 {
				_gameControl.dealersPlay = true
			}
		} else if _gameControl.dealersPlay {
			//Get Dealer Choice
			if !dealerChoice() {
				_gameControl.gameOver = true
			}

			if dealer.isBust() {
				_gameControl.gameOver = true
				_gameControl.gameResult = 6
			}
		}

		fmt.Print("\n\n")
	}
	//If game is over but no result was decided
	if _gameControl.gameResult == -1 {
		_gameControl.gameResult = compareHandValues()
	} 
	
	//The Game didn't end by the Player Standing Up
	if _gameControl.gameResult != 7 {
		time.Sleep(1 * time.Second)

		fmt.Println("------ FINAL HANDS ------")

		dealer.state()
		player.state()

		endOfGame()
	}	
}

func deal() {
	//Create the initial player and dealer hand
	playerHand := make([]Card, 0)
	dealerHand := make([]Card, 0)
	//Draw Cards Sequentially
	playerCardOne := deck.Draw()
	dealerCardOne := deck.Draw()
	playerCardTwo := deck.Draw()
	dealerCardTwo := deck.Draw()
	//Add cards to the players hand and calculate its value
	playerHand = append(playerHand, playerCardOne)
	playerHand = append(playerHand, playerCardTwo)
	player.hand = playerHand
	player.calculateHandValue()

	//Add Cards to the dealers hand
	dealerHand = append(dealerHand, dealerCardOne)
	dealerHand = append(dealerHand, dealerCardTwo)
	dealer.hand = dealerHand
	dealer.calculateHandValue()
}

func (_player *Player) hit() {
	newCard := deck.Draw()
	_player.hand = append(_player.hand, newCard)
	_player.calculateHandValue()
}

func (_player *Player) isBust() bool {
	if _player.handValue > 21 {
		return true
	}
	return false
}

func (_player *Player) calculateHandValue() {
	var val int = 0

	for _, card := range _player.hand {
		//The card is an Ace
		if card.CardValue() == 11 {
			//If the ace would put the hand value above 21 use as a 1
			if val + card.CardValue() > 21 {
				val += 1 
				continue
			}
		}

		val += card.CardValue()
	}

	_player.handValue = val
}
//The dealer hits unless cards are 17 or greater
func dealerChoice() bool {
	if dealer.handValue < 17 {
		dealer.hit()
		return true
	}
	return false
}

func (_player *Player) isBlackJack() bool {
	if _player.handValue == 21 {
		return true
	}
	return false
}

func compareHandValues() int {
	if player.handValue == dealer.handValue {
		return 0
	}
	if player.handValue > dealer.handValue {
		return 1
	}
	if player.handValue < dealer.handValue {
		return 2
	}

	return -1
}

func bet() {
	var input int
	var betting bool = true

	for betting {
		fmt.Print("Enter Bet Amount or [-1] to Stand-Up: \t")
		if _betControl.playerChipTotal <= 0 {
			fmt.Println("You're outta chips! Get Outta Here!")
		}
		fmt.Scanln(&input)

		if input > _betControl.playerChipTotal {
			fmt.Println("You don't have enough chips for that. Enter a new bet or stand-up.")
			continue
		}

		if input < _betControl.minBet && input > 0{
			fmt.Printf("You gotta bet more than that cheapo! Minimum bet is: %d\n" , _betControl.minBet)
			continue
		}

		if input == 0 {
			fmt.Println("Changed your mind?")
			continue
		}

		if input == -1 {
			_gameControl.gameOver = true
			_gameControl.gameResult = 7
		}

		_betControl.playerBet = input;
		betting = false
	}
}

func (_player *Player) state() {
	var isDealer = _player.isDealer
	if !isDealer {
		for _, card := range _player.hand {
			fmt.Printf(card.CardString() + " ")
		}
		fmt.Print("\n")

		fmt.Printf("Player Hand Value : %d\n", _player.handValue)
	} else if isDealer && _gameControl.dealersPlay || _gameControl.gameOver {
		for _, card := range _player.hand {
			fmt.Print(card.CardString() + " ")
		}

		fmt.Print("\n")

		fmt.Printf("Dealer Hand Value : %d\n", _player.handValue)
	} else {

		fmt.Printf(_player.hand[0].CardString() + " " + "HIDDEN" + "\n")

		fmt.Printf("Dealer Hand Value : %d\n\n", _player.hand[0].CardValue())
	}
}

func endOfGame() {
	switch _gameControl.gameResult {
		case 0:
			fmt.Println("------ THE GAME ENDED IN A DRAW ------")
		case 1:
			fmt.Println("------ THE PLAYER WINS ------")
			_betControl.playerChipTotal += _betControl.playerBet
		case 2:
			fmt.Println("------ THE DEALER WINS ------")
			_betControl.playerChipTotal -= _betControl.playerBet
		case 3:
			fmt.Println("------ PLAYER WINS WITH A BLACKJACK! WOW! ------")
			_betControl.playerChipTotal += _betControl.playerBet + _betControl.playerBet / 2
		case 4:
			fmt.Println("------ DEALER WINS WITH A BLACKJACK! UNFORTUNATE! ------")
			_betControl.playerChipTotal -= _betControl.playerBet
		case 5:
			fmt.Println("------ THE PLAYERS BUSTS. DEALER WINS ------")
			_betControl.playerChipTotal -= _betControl.playerBet
		case 6:
			fmt.Println("------ THE DEALER BUSTS. PLAYER WINS ------")
			_betControl.playerChipTotal += _betControl.playerBet
	}

	_betControl.playerBet = 0
}