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
	fmt.Printf("Word Two: ")
	b, _ := reader.ReadString('\n')
	t, g := analyze.CowsBulls(a, b)
	fmt.Printf("C: %d ... B: %d\n", t, g)
	os.Exit(0)
}
