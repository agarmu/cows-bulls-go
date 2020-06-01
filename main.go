package main

import (
	"fmt"

	"syscall"

	"github.com/agarmu/cows-bulls-go/analyze"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {

	fmt.Printf("Target Word (will not be visible): ")
	a, _ := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println('\n')
	inA, errA := analyze.Validate(string(a))
	if errA != nil {
		fmt.Printf("error: ")
		fmt.Print(errA)
		fmt.Printf("\n")
	}

	fmt.Printf("Enter Guess: ")
	var b string
	fmt.Scanf("%s", &b)
	inB, errB := analyze.Validate(b)
	if errB != nil {
		fmt.Println(errB)
	}
	cows, bulls := analyze.CowsBulls(inA, inB)
	fmt.Printf("Cows %d, Bulls %d\n", cows, bulls)
}
