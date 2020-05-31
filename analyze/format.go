package analyze

import (
	"errors"
	"strings"
)

func split(str string) []string {
	var arr []string
	for _, r := range str {
		arr = append(arr, string(r))
	}
	return arr
}

func isUnique(i []string) bool {

	check := make(map[string]int)
	res := make([]string, 0)
	for _, val := range i {
		check[val] = 1
	}

	for letter := range check {
		res = append(res, letter)
	}

	return (len(res) == len(i))
}

//Validate validates an input and returns the
func Validate(str string) ([]string, error) {
	sliced := split(strings.ToLower(str))
	if !isUnique(sliced) {
		return sliced, errors.New("Input not distinct/unique")
	}
	return sliced, nil
}
