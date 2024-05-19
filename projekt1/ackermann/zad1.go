package main

import (
	"fmt"
	"math/big"
	"time"
)

func ackermann(m, n *big.Int) *big.Int {
	zero := big.NewInt(0)
	one := big.NewInt(1)

	if m.Cmp(zero) == 0 {
		return new(big.Int).Add(n, one)
	} else if n.Cmp(zero) == 0 {
		return ackermann(new(big.Int).Sub(m, one), one)
	} else {
		return ackermann(new(big.Int).Sub(m, one), ackermann(m, new(big.Int).Sub(n, one)))
	}
}

func ackermann2(m, n int) int {
	if m == 0 {
		return n + 1
	} else if n == 0 {
		return ackermann2(m-1, 1)
	} else {
		return ackermann2(m-1, ackermann2(m, n-1))
	}
}

func main() {
	start := time.Now()
	// m := big.NewInt(4) // Change these values as needed
	// n := big.NewInt(1) // Change these values as needed
	// result := ackermann(m, n)
	ackermann2(3, 1)
	elapsed := time.Since(start)

	// fmt.Printf("Ackermann function result: ", result)
	fmt.Printf("Time taken to compute: %s\n", elapsed)
}

// potrzebuje zrobić jakieś przybliżenie ile mogłaby się wykonywać funkcja Ackermanna dla m=261, n=18. Dla m=3 n=5 czas=2.25275ms, dla m=3 n=6 czas=9.63875ms, dla m=4
