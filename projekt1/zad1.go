package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"
)

func gen_ascii_for_name(name string, surname string) []int {
	nick := name[:3] + surname[:3]
	nick = strings.ToLower(nick)
	fmt.Println(nick)

	var result []int
	for _, c := range nick {
		result = append(result, int(c))
		fmt.Println(c, result)
	}
	return result
}

func factorial(n int) *big.Int {
	factorial := new(big.Int).MulRange(1, int64(n))
	return factorial
}

func contains_bytes(s string, intArray []int) bool {
	for _, c := range intArray {
		if !strings.Contains(s, strconv.Itoa(c)) {
			return false
		}
	}
	return true
}

func check_if_factorial_contains(intArray []int, n int) int {
	for {
		factorial := factorial(n).String()

		if contains_bytes(factorial, intArray) {
			return n
		} else {
			n++
		}
	}
}

func fibFrequency(n int) int {
	cache[n]++
	// fmt.Println(cache)

	switch n {
	case 0, 1:
		return n
	default:
		return fibFrequency(n-1) + fibFrequency(n-2)
	}
}

func fibFrequencyClosestToStrongNumber(cache map[int]int, strongNumber int) int {
	minDiff := float64(strongNumber - cache[len(cache)-1])
	small := len(cache) - 1

	for key, value := range cache {
		diff := strongNumber - value
		if math.Abs(float64(diff)) < minDiff {
			minDiff = math.Abs(float64(diff))
			small = key
		}
	}
	return small
}

func main() {
	fmt.Println(factorial(261))
	array := gen_ascii_for_name("Oskar", "Gawryszewski")
	fmt.Println(array)
	strong_num := check_if_factorial_contains(array, 1)

	fmt.Println("My strong number is: ", strong_num)

	start := time.Now()
	fibFrequency(39)
	fmt.Println(time.Since(start))

	weak_num := fibFrequencyClosestToStrongNumber(cache, strong_num)
	fmt.Println("My weak number is: ", weak_num)
}

var cache = make(map[int]int)
