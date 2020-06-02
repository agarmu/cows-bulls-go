package analyze

//CowsBulls returns cows and bulls for a specific guess
func CowsBulls(target, guess []string) (int, int) {
	cows, bulls := 0, 0
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

