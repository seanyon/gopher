package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	rand.Seed(t)
	s := rand.Intn(6)
	n := s + 1

	switch n {
	case 6:
		fmt.Printf("%d-大吉\n", n)
	case 5, 4:
		fmt.Printf("%d-中吉\n", n)
	case 3, 2:
		fmt.Printf("%d-小吉\n", n)
	default:
		fmt.Printf("%d-凶\n", n)
	}

}
