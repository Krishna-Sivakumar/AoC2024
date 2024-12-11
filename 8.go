// we solve a parametric equation for this one hehehe
// first math application of this AoC
package main

import (
	"common"
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

type Pair struct {
	X1, Y1, X2, Y2 int
}

func MakePair(pointA, pointB Point) Pair {
	return Pair{
		X1: pointA.X,
		Y1: pointA.Y,
		X2: pointB.X,
		Y2: pointB.Y,
	}
}

func getAntinodes(a, b Point) (a1 Point, a2 Point) {
	var slope float64 = float64(b.Y-a.Y) / float64(b.X-a.X)
	theta := math.Atan(slope)
	d := math.Sqrt(math.Pow((float64(b.Y)-float64(a.Y)), 2) + math.Pow(float64(b.X)-float64(a.X), 2))

	a1.X = int(math.Round(float64(a.X) - d*math.Cos(theta)))
	a1.Y = int(math.Round(float64(a.Y) - d*math.Sin(theta)))
	a2.X = int(math.Round(float64(b.X) + d*math.Cos(theta)))
	a2.Y = int(math.Round(float64(b.Y) + d*math.Sin(theta)))
	return
}

func PartOne(filename string) {
	handler, err := common.FetchFile(filename)
	if err != nil {
		panic(err)
	}
	defer handler.Cleanup()

	points := make(map[byte][]Point)

	antennae := make(map[Point]bool)
	uniques := make(map[Point]bool)
	uniquePairs := make(map[Pair]bool)

	// y := len(buffer) - 1
	buffer := make([][]byte, 0)
	x := 0
	for {
		line, _, err := handler.Reader.ReadLine()
		if err != nil {
			break
		}

		// for x, char := range line {
		for y, char := range line {
			if char != '.' {
				antennae[Point{x, y}] = true
				points[char] = append(points[char], Point{x, y})
			}
		}

		// y -= 1
		x += 1
		buffer = append(buffer, line)
	}

	for _, points := range points {
		for _, pointA := range points {
			for _, pointB := range points {
				if pointA == pointB || uniquePairs[MakePair(pointA, pointB)] || uniquePairs[MakePair(pointB, pointA)] {
					continue
				}

				antiNode1, antiNode2 := getAntinodes(pointA, pointB)
				flippedAntiNode1, flippedAntiNode2 := getAntinodes(pointB, pointA)

				if (antiNode1 == pointA && antiNode2 == pointB) || (antiNode1 == pointB && antiNode2 == pointA) {
					// we need to use flipped antinodes
					if common.ValidGridIndex(flippedAntiNode1.X, flippedAntiNode1.Y, len(buffer[0]), len(buffer)) {
						uniques[flippedAntiNode1] = true
					}
					if common.ValidGridIndex(flippedAntiNode2.X, flippedAntiNode2.Y, len(buffer[0]), len(buffer)) {
						uniques[flippedAntiNode2] = true
					}
				} else {
					if common.ValidGridIndex(antiNode1.X, antiNode1.Y, len(buffer[0]), len(buffer)) {
						uniques[antiNode1] = true
					}
					if common.ValidGridIndex(antiNode2.X, antiNode2.Y, len(buffer[0]), len(buffer)) {
						uniques[antiNode2] = true
					}
				}
				uniquePairs[MakePair(pointA, pointB)] = true
			}
		}
	}
	for x, line := range buffer {
		for y, cell := range line {
			if uniques[Point{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(string(cell))
			}
		}
		fmt.Println()
	}
	fmt.Println(len(uniques))
}

func getAllAntinodes(a, b Point, boundX, boundY int) (rightNodes []Point, leftNodes []Point) {
	var slope float64 = float64(b.Y-a.Y) / float64(b.X-a.X)
	theta := math.Atan(slope)
	d := math.Sqrt(math.Pow((float64(b.Y)-float64(a.Y)), 2) + math.Pow(float64(b.X)-float64(a.X), 2))

	x := int(math.Round(float64(a.X) - d*math.Cos(theta)))
	y := int(math.Round(float64(a.Y) - d*math.Sin(theta)))
	for {
		if common.ValidGridIndex(x, y, boundX, boundY) {
			leftNodes = append(leftNodes, Point{x, y})
		} else {
			break
		}
		x = int(math.Round(float64(x) - d*math.Cos(theta)))
		y = int(math.Round(float64(y) - d*math.Sin(theta)))

	}

	x = int(math.Round(float64(a.X) + d*math.Cos(theta)))
	y = int(math.Round(float64(a.Y) + d*math.Sin(theta)))
	for {
		if common.ValidGridIndex(x, y, boundX, boundY) {
			rightNodes = append(rightNodes, Point{x, y})
		} else {
			break
		}
		x = int(math.Round(float64(x) + d*math.Cos(theta)))
		y = int(math.Round(float64(y) + d*math.Sin(theta)))

	}

	return
}

func PartTwo(filename string) {
	handler, err := common.FetchFile(filename)
	if err != nil {
		panic(err)
	}
	defer handler.Cleanup()

	points := make(map[byte][]Point)

	antennae := make(map[Point]bool)
	uniques := make(map[Point]bool)
	uniquePairs := make(map[Pair]bool)

	// y := len(buffer) - 1
	buffer := make([][]byte, 0)
	x := 0
	for {
		line, _, err := handler.Reader.ReadLine()
		if err != nil {
			break
		}

		// for x, char := range line {
		for y, char := range line {
			if char != '.' {
				antennae[Point{x, y}] = true
				points[char] = append(points[char], Point{x, y})
			}
		}

		// y -= 1
		x += 1
		buffer = append(buffer, line)
	}

	for _, points := range points {
		for _, pointA := range points {
			for _, pointB := range points {
				if pointA == pointB || uniquePairs[MakePair(pointA, pointB)] || uniquePairs[MakePair(pointB, pointA)] {
					continue
				}

				rightNodes, leftNodes := getAllAntinodes(pointA, pointB, len(buffer), len(buffer[0]))
				for _, point := range rightNodes {
					uniques[point] = true
				}
				for _, point := range leftNodes {
					uniques[point] = true
				}

				uniquePairs[MakePair(pointA, pointB)] = true
			}
		}
	}

	// each antenna is also an antinode now
	for antenna, _ := range antennae {
		uniques[antenna] = true
	}

	// printing grid here
	for x, line := range buffer {
		for y, cell := range line {
			if uniques[Point{x, y}] && cell == '.' {
				fmt.Print("#")
			} else {
				fmt.Print(string(cell))
			}
		}
		fmt.Println()
	}
	fmt.Println(len(uniques))
}

func main() {
	filename := common.GetProblemFile(8)
	fmt.Println("part one:")
	PartOne(filename)
	fmt.Println("part two:")
	PartTwo(filename)
}
