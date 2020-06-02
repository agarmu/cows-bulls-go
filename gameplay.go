package main

import (
	"fmt"
	"syscall"
  "strings"
	"github.com/agarmu/cows-bulls-go/analyze"
	"golang.org/x/crypto/ssh/terminal"
)


//Game is a Basic game of cows/bulls
type Game struct {
	target     []string
	maxGuesses int
	guesses    []guess
}

type guess struct {
	word  []string
	cows  int
	bulls int
}

func initGame() *Game {

	var target []string
	var maxGuesses int = 6
	game:= Game{}
	for {
		fmt.Printf("-----------------------------------\n")
		fmt.Printf("Enter target word and press enter: ")
		a, err := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Printf("\n")
		if err != nil {
			fmt.Printf("Error: ")
			fmt.Println(err)
			fmt.Printf("Please try again:\n")
			continue
		}
		target, err = analyze.Validate(string(a))
		if err != nil {
			fmt.Printf("Error: ")
			fmt.Println(err)
			fmt.Printf("Please try again:\n")
			continue
		}
		break
	}
  game.target = target
  game.maxGuesses = maxGuesses
  return &game
}
func (game *Game) guess(guessesLeft int) {
  g := guess{}
  var word string
  var err error
  for{
  fmt.Printf("-----------------------------------\n")
  fmt. Printf("Guesses left: %d\n", guessesLeft)
  fmt.Printf("Enter Guess: ")
	fmt.Scanf("%s", &word)
	g.word, err = analyze.Validate(word)
	if err != nil {
    fmt.Printf("Error: ")
		fmt.Println(err)
    fmt.Println("Please try again")
    continue
	}
  if len(g.word) != len(game.target) {
    fmt.Printf("Guess length is not the same as target length.\nPlease try again.\n")
    continue
  }
  g.cows, g.bulls = analyze.CowsBulls(game.target, g.word)
  break
  }
  game.guesses = append(game.guesses, g)
  fmt.Printf("Cows: %d ... Bulls: %d\n", g.cows, g.bulls)
}
func (game *Game) checkWinner() (bool) {
  l := len(game.target)
  lastGuess := game.guesses[len(game.guesses)-1]
  if lastGuess.bulls == l {
    return true
  }
  return false
}

func (game *Game) play() {
  target := strings.Join(game.target, string(""))
  for {
    guessesLeft := game.maxGuesses - len(game.guesses)
    if guessesLeft == 0 {
      fmt.Println("-----------------------------------\n")
      fmt.Println("Sorry, but you have no guesses left.")
      fmt.Printf("The word you were trying to guess was %s\n", target)
      break
    }
    game.guess(guessesLeft)
    won := game.checkWinner()
    if won {
      fmt.Println("-----------------------------------")
      fmt.Printf("You won with %d guesses remaining\n", guessesLeft-1)
      break
    }
  }
}

