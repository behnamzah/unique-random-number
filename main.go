package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	max, _ := strconv.Atoi(os.Args[1])
	var us []int

loop:
	for len(us) < max {
		n := rand.Intn(max) + 1
		fmt.Print(n, " ")

		for _, u := range us {
			if u == n {
				continue loop
			}
		}

		us = append(us, n)
	}

	fmt.Println("\n uniques :", us)
	sort.Ints(us)
	fmt.Println("\n sorted  :", us)

}
