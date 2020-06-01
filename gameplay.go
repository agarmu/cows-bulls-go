package main

import (
	"fmt"
	"syscall"

	"github.com/agarmu/cows-bulls-go/analyze"
	"golang.org/x/crypto/ssh/terminal"
)

/*import  (
	"fmt"

	"syscall"

	"github.com/agarmu/cows-bulls-go/analyze"
	"golang.org/x/crypto/ssh/terminal"
)*/

//Game is a Basic game of cows/bulls
type Game struct {
	target     []string
	maxGuesses int
	guesses    []guess
}

type guess struct {
	word  string
	cows  int
	bulls int
}

func initGame() Game {

	var target []string
	var maxGuesses int = 6
	game := Game{}
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

}
