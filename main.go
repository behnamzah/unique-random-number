package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	const max = 5
	var us [max]int

loop:
	for found := 0; found < max; {
		n := rand.Intn(max) + 1
		fmt.Print(n, " ")

		for _, u := range us {
			if u == n {
				continue loop
			}
		}

		us[found] = n
		found++
	}

	fmt.Println("\n\n uniques :", us)

}
