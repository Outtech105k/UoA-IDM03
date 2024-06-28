package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var userName string
	fmt.Println("Who are you?")
	fmt.Print("> ")
	fmt.Scan(&userName)
	fmt.Printf("Hello, %s!\n", userName)

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

	if heads > tails {
		fmt.Println(userName, "won!")
	} else {
		fmt.Println(userName, "lost!")
	}
}
