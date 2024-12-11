// I genuinely hope I never have to write code like this ever again.
// This problem fills me with a lot of rage and sadness.
package main

import (
	"common"
	"fmt"
)

type Direction struct {
	X, Y int
}

func Dfs(searchSpace *[]string, currX, currY, depth int, direction Direction) bool {
	if !common.ValidGridIndex(currX, currY, len(*searchSpace), len((*searchSpace)[0])) {
		return false
	}

	switch depth {
	case 0:
		if (*searchSpace)[currX][currY] == 'X' {
			return Dfs(searchSpace, currX+direction.X, currY+direction.Y, depth+1, direction)
		} else {
			return false
		}
	case 1:
		if (*searchSpace)[currX][currY] == 'M' {
			return Dfs(searchSpace, currX+direction.X, currY+direction.Y, depth+1, direction)
		} else {
			return false
		}
	case 2:
		if (*searchSpace)[currX][currY] == 'A' {
			return Dfs(searchSpace, currX+direction.X, currY+direction.Y, depth+1, direction)
		} else {
			return false
		}
	case 3:
		if (*searchSpace)[currX][currY] == 'S' {
			return true
		} else {
			return false
		}
	}
	// can't get to this point but gotta deal with the go compiler
	return true
}

func PartOne(filename string) {
	handler, err := common.FetchFile(filename)
	if err != nil {
		panic(err)
	}
	defer handler.Cleanup()

	searchSpace := make([]string, 0, 140)
	for {
		line, _, err := handler.Reader.ReadLine()
		if err != nil {
			break
		}
		searchSpace = append(searchSpace, string(line))
	}

	count := 0
	for i := 0; i < len(searchSpace); i++ {
		for j := 0; j < len(searchSpace[i]); j++ {
			localCount := 0
			tries := make([]bool, 0, 8)
			tries = append(tries, Dfs(&searchSpace, i, j, 0, Direction{X: 1, Y: 0}))  // down
			tries = append(tries, Dfs(&searchSpace, i, j, 0, Direction{X: 0, Y: 1}))  // right
			tries = append(tries, Dfs(&searchSpace, i, j, 0, Direction{X: -1, Y: 0})) // left
			tries = append(tries, Dfs(&searchSpace, i, j, 0, Direction{X: 0, Y: -1})) // top
			tries = append(tries, Dfs(&searchSpace, i, j, 0, Direction{X: 1, Y: 1}))
			tries = append(tries, Dfs(&searchSpace, i, j, 0, Direction{X: 1, Y: -1}))
			tries = append(tries, Dfs(&searchSpace, i, j, 0, Direction{X: -1, Y: 1}))
			tries = append(tries, Dfs(&searchSpace, i, j, 0, Direction{X: -1, Y: -1}))

			for _, try := range tries {
				if try {
					localCount += 1
				}
			}

			count += localCount
		}
	}

	fmt.Println(count)
}

type Location struct {
	X, Y int
}

func CheckCell(searchSpace *[]string, i, j int, orientation int) bool {
	// we check if there's an X-MAS by validating a kernel
	// since there could be different rotations of the pattern, we go through each one
	center := Location{i, j}
	bottomLeft := Location{i + 1, j - 1}
	bottomRight := Location{i + 1, j + 1}
	topLeft := Location{i - 1, j - 1}
	topRight := Location{i - 1, j + 1}

	isValid := func(i, j int) bool {
		return common.ValidGridIndex(i, j, len(*searchSpace), len((*searchSpace)[0]))
	}

	originalOrientation := []Location{bottomLeft, bottomRight, topRight, topLeft}

	actualOrientation := common.Rotate(originalOrientation, orientation)
	mOne := actualOrientation[0]
	mTwo := actualOrientation[1]
	sOne := actualOrientation[2]
	sTwo := actualOrientation[3]

	if !isValid(bottomLeft.X, bottomLeft.Y) || !isValid(bottomRight.X, bottomLeft.Y) || !isValid(topLeft.X, topLeft.Y) || !isValid(topRight.X, topRight.Y) {
		return false
	}

	if (*searchSpace)[center.X][center.Y] == 'A' && (*searchSpace)[mOne.X][mOne.Y] == 'M' && (*searchSpace)[mTwo.X][mTwo.Y] == 'M' && (*searchSpace)[sOne.X][sOne.Y] == 'S' && (*searchSpace)[sTwo.X][sTwo.Y] == 'S' {
		return true
	} else {
		return false
	}
}

func PartTwo(filename string) {
	handler, err := common.FetchFile(filename)
	if err != nil {
		panic(err)
	}
	defer handler.Cleanup()

	searchSpace := make([]string, 0, 140)
	for {
		line, _, err := handler.Reader.ReadLine()
		if err != nil {
			break
		}
		searchSpace = append(searchSpace, string(line))
	}

	count := 0
	for i := 0; i < len(searchSpace); i++ {
		for j := 0; j < len(searchSpace[i]); j++ {
			for orientation := 0; orientation < 4; orientation++ {
				if CheckCell(&searchSpace, i, j, orientation) {
					count += 1
					break
				}
			}
		}
	}

	fmt.Println(count)
}

func main() {
	filename := common.GetProblemFile(4)
	fmt.Println("part one:")
	PartOne(filename)
	fmt.Println("part two:")
	PartTwo(filename)
}
