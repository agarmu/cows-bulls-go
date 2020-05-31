package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/agarmu/cows-bulls-go/analyze"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Word One: ")
	a, _ := reader.ReadString('\n')
	inA, errA := analyze.Validate(a)
	if errA != nil {
		fmt.Println(errA)
	}
	fmt.Printf("Word Two: ")
	b, _ := reader.ReadString('\n')
	inB, errB := analyze.Validate(b)
	if errB != nil {
		fmt.Println(errA)
	}
	cows, bulls := analyze.CowsBulls(inA, inB)
	fmt.Printf("C: %d ... B: %d\n", cows, bulls)
	os.Exit(0)
}
