// forest.module.go

package modules

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var totalTrees int
var burnedTrees int
var imageCounter int

type Tree struct {
	IsBurning bool
}

type Shrub struct{}

type Forest struct {
	Width, Height int
	Trees         [][]interface{}
}

type ForestGenerator interface {
	Generate() *Forest
}

type RandomForestGenerator struct {
	Width, Height, Density int
}

func (rfg *RandomForestGenerator) Generate() *Forest {
	rand.Seed(time.Now().UnixNano())

	forest := &Forest{
		Width:  rfg.Width,
		Height: rfg.Height,
		Trees:  make([][]interface{}, rfg.Width),
	}

	for i := range forest.Trees {
		forest.Trees[i] = make([]interface{}, rfg.Height)
		for j := range forest.Trees[i] {
			if rand.Intn(100) < rfg.Density {
				forest.Trees[i][j] = &Tree{}
			} else {
				forest.Trees[i][j] = &Shrub{}
			}
		}
	}

	return forest
}

func (f *Forest) SimulateLightningStrikeRecursive() {
	rand.Seed(time.Now().UnixNano())
	os.RemoveAll("images")

	// SaveImage(*f, imageCounter)
	imageCounter++

	x := rand.Intn(f.Width)
	y := rand.Intn(f.Height)

	if tree, ok := f.Trees[x][y].(*Tree); ok {
		tree.IsBurning = true
	}

	// SaveImage(*f, imageCounter)
	imageCounter++

	// fmt.Println("Lightning strike at:", x, y)

	totalTrees = f.Width * f.Height
	burnedTrees = 0

	f.simulateBurning(x, y)

	// SaveImage(*f, imageCounter)

	// burnedTreesPercentage := f.CalculateBurnedTreesPercentage()
	// fmt.Printf("Percentage of burned trees: %.2f%%\n", burnedTreesPercentage)
}

func (f *Forest) simulateBurning(x, y int) {
	os.MkdirAll("images", os.ModePerm)
	if tree, ok := f.Trees[x][y].(*Tree); ok {
		// fmt.Println("Checking tree at:", x, y)
		if tree.IsBurning {
			burnedTrees++
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if i == 0 && j == 0 {
						continue
					}
					newX, newY := x+i, y+j
					if newX >= 0 && newX < f.Width && newY >= 0 && newY < f.Height {
						if neighbor, ok := f.Trees[newX][newY].(*Tree); ok && !neighbor.IsBurning {
							// fmt.Println("Neighbor at:", newX, newY, "is now burning")
							neighbor.IsBurning = true

							if burnedTrees%(totalTrees/20) == 0 && imageCounter < 20 {
								// SaveImage(*f, imageCounter)
								imageCounter++
							}

							f.simulateBurning(newX, newY)
						}
					}
				}
			}
		} else {
			fmt.Println("Tree at:", x, y, "is not burning, stopping recursion.")
			return
		}
	}
}

func displayForest(forest Forest) {
	for _, row := range forest.Trees {
		for _, obj := range row {
			switch obj.(type) {
			case *Tree:
				if obj.(*Tree).IsBurning {
					fmt.Print("🔥")
				} else {
					fmt.Print("🌲")
				}
			case *Shrub:
				fmt.Print("🌳")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (f *Forest) CalculateBurnedTreesPercentage() float64 {
	burnedTrees := 0
	for _, row := range f.Trees {
		for _, obj := range row {
			if tree, ok := obj.(*Tree); ok && tree.IsBurning {
				burnedTrees++
			}
		}
	}
	return float64(burnedTrees) / float64(f.Width*f.Height) * 100
}

const outputFolder = "images/"

func SaveImage(forest Forest, iteration int) {
	img := image.NewRGBA(image.Rect(0, 0, forest.Width*20, forest.Height*20))

	treeColor := color.RGBA{0, 128, 0, 255}    // Green
	fireColor := color.RGBA{255, 0, 0, 255}    // Red
	shrubColor := color.RGBA{139, 69, 19, 255} // Brown
	emptyColor := color.RGBA{255, 255, 255, 255}

	for y := 0; y < forest.Height; y++ {
		for x := 0; x < forest.Width; x++ {
			var c color.Color
			switch obj := forest.Trees[x][y].(type) {
			case *Tree:
				if obj.IsBurning {
					c = fireColor
				} else {
					c = treeColor
				}
			case *Shrub:
				c = shrubColor
			default:
				c = emptyColor
			}
			for j := 0; j < 20; j++ {
				for i := 0; i < 20; i++ {
					img.Set(x*20+i, y*20+j, c)
				}
			}
		}
	}

	err := os.MkdirAll(outputFolder, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	fileName := outputFolder + "iteration_" + strconv.Itoa(iteration) + ".png"
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating image file:", err)
		return
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		fmt.Println("Error encoding image:", err)
		return
	}

	fmt.Println("Image saved:", fileName)
}
