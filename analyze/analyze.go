package analyze

import (
	"fmt"
)

//CowsBulls returns cows and bulls for a specific guess
func CowsBulls(target, guess []string) (int, int) {
	cows, bulls := 0, 0

	if len(target) != len(guess) {
		fmt.Printf("Target Length (%d) not equal to guess length (%d)", len(target), len(guess))
	}
	for i := 0; i < len(target); i++ {
		for j := 0; j < len(guess); j++ {
			if target[i] == guess[j] {
				if i == j {
					bulls++
				} else {
					cows++
				}
			}
		}
	}
	return cows, bulls
}
