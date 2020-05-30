package analyze

import (
	"log"
)

//CowsBulls returns cows and bulls for a specific guess
func CowsBulls(target, guess string) (int, int) {
	cows, bulls := 0, 0
	t, g := SplitStr(target), SplitStr(guess)

	if len(t) != len(g) {
		log.Panicf("Target Length (%d) not equal to guess length (%d)", len(t), len(g))
	}
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(g); j++ {
			if t[i] == g[j] {
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

//SplitStr splits strings
func SplitStr(str string) []string {
	var arr []string
	for _, r := range str {
		arr = append(arr, string(r))
	}
	return arr
}
