package analyze

import (
	"log"
)

//CowsBulls returns cows and bulls for a specific guess
func CowsBulls(target, guess []string) (int, int) {
	cows, bulls := 0, 0

	if len(target) != len(guess) {
		log.Panicf("Target Length (%d) not equal to guess length (%d)", len(target), len(guess))
	}
	for i := 0; i < len(target)-2; i++ {
		for j := 0; j < len(guess)-2; j++ {
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
