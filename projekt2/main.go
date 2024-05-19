// W pliku main.go

package main

import (
	"fmt"
	"os/exec"
	"projekt2/modules"
	"strconv"
	"strings"
)

func forest_structure(m int, n int, density int) float64 {
	generator := &modules.RandomForestGenerator{
		Width:   m,
		Height:  n,
		Density: density,
	}

	forest := generator.Generate()

	// fmt.Println("Forest structure before lightning strike:")
	// displayForest(forest)

	forest.SimulateLightningStrikeRecursive()

	// fmt.Println("Forest structure after lightning strike:")
	// displayForest(forest)

	burnedTrees := forest.CalculateBurnedTreesPercentage()

	return burnedTrees

}

func displayForest(forest *modules.Forest) {
	for _, row := range forest.Trees {
		for _, obj := range row {
			switch obj.(type) {
			case *modules.Tree:
				if obj.(*modules.Tree).IsBurning {
					fmt.Print("🔥")
				} else {
					fmt.Print("🌲")
				}
			case *modules.Shrub:
				fmt.Print("🌳")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	meanTestValues := make(map[int]float64)

	for i := 5; i <= 95; i += 5 {
		sumTestValues := 0.0

		for j := 0; j < 1000; j++ {
			unburnedTrees := (100 - forest_structure(10, 50, i))
			sumTestValues += unburnedTrees
		}

		meanTestValue := sumTestValues / 1000
		meanTestValues[i] = meanTestValue

		fmt.Printf("Mean test value for density %d: %.2f\n", i, meanTestValue)
	}

	meanTestRatios := make(map[int]float64)
	for meanTestValuesKey, meanTestValuesValue := range meanTestValues {
		meanTestRatios[meanTestValuesKey] = meanTestValuesValue * float64(meanTestValuesKey)
	}

	densities := make([]string, 0, len(meanTestRatios))
	ratios := make([]string, 0, len(meanTestRatios))
	fmt.Println("Mean test ratios:")
	for meanTestRatiosKey, meanTestRatiosValue := range meanTestRatios {
		densities = append(densities, strconv.Itoa(meanTestRatiosKey))
		ratios = append(ratios, fmt.Sprintf("%.2f", meanTestRatiosValue))
		fmt.Printf("Density: %d, Ratio: %.2f\n", meanTestRatiosKey, meanTestRatiosValue)
	}

	cmd := exec.Command("python3", "graphs.py", strings.Join(densities, ","), strings.Join(ratios, ","))
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	// forest_structure(20, 20, 60)

}
