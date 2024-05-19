package main

import (
	"fmt"
	"math/rand"
)

func game(n int, strategy bool) {
	wins := 0

	for round := 0; round < n; round++ {
		winning_box := rand.Intn(3)
		player_box := rand.Intn(3)

		var revealed_box int
		for {
			revealed_box = rand.Intn(3)
			if revealed_box != winning_box && revealed_box != player_box {
				break
			}
		}

		switch strategy {
		case true:
			for i := 0; i < 3; i++ {
				if i != player_box && i != revealed_box {
					player_box = i
					break
				}
			}
		}

		if player_box == winning_box {
			wins += 1
		}
	}

	fmt.Printf("With strategy %v, won %d out of %d games\n", strategy, wins, n)
	win_percentage := float64(wins) / float64(n) * 100
	fmt.Printf("Win percentage: %.2f%%\n", win_percentage)
}

func game_simulation(n int, strategy bool, total_boxes int, boxes_to_open int) {
	wins := 0

	for round := 0; round < n; round++ {
		winning_box := rand.Intn(total_boxes)
		player_box := rand.Intn(total_boxes)

		revealed_boxes := make(map[int]bool)
		for len(revealed_boxes) < boxes_to_open {
			revealed_box := rand.Intn(total_boxes)
			if revealed_box != winning_box && revealed_box != player_box {
				revealed_boxes[revealed_box] = true
			}
		}

		if strategy {
			free_indexes := make([]int, 0)
			for i := 0; i < total_boxes; i++ {
				if !revealed_boxes[i] && i != player_box {
					free_indexes = append(free_indexes, i)
				}
			}
			random_index := rand.Intn(len(free_indexes))
			player_box = free_indexes[random_index]

		}
		if player_box == winning_box {
			wins++
		}
	}

	fmt.Printf("With strategy %v, won %d out of %d games\n", strategy, wins, n)
	win_percentage := float64(wins) / float64(n) * 100
	fmt.Printf("Win percentage: %.2f%%\n", win_percentage)
}
func main() {
	game(1000, false)
	game(1000, true)

	fmt.Println("-------------------")

	game_simulation(10000, false, 100, 1)
	game_simulation(10000, true, 100, 1)

}
