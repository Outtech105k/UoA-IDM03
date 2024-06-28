package main

import (
	"fmt"
	"math/rand"
)

func main() {
	heads := 0
	tails := 0

	fmt.Println("Tossing a coin...")
	for i := 1; i <= 3; i++ {
		if rand.Intn(2) == 0 {
			fmt.Printf("Round %d: Heads\n", i)
			heads++
		} else {
			fmt.Printf("Round %d: Tails\n", i)
			tails++
		}
	}
	fmt.Printf("Heads: %d, Tails: %d\n", heads, tails)
}
