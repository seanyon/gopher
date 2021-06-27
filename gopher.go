package main

/*func main() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%d - 偶数\n", i)
		} else {
			fmt.Printf("%d - 奇数\n", i)
		}
	}
}*/

/*func main() {
	for a := 1; a <= 100; a++ {
		switch {
		case
			a % 2 == 0:
			fmt.Printf("%d - 偶数\n", a)
		default:
			fmt.Printf("%d - 奇数\n", a)
		}
	}
}*/

import (
	"math/rand"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	rand.Seed(t)
	s := rand.Intn(10)
	println(s)
}
