package main

import (
	"common"
	"fmt"
)

type Object int

const (
	Free Object = iota
	Obstacle
)

type Direction struct {
	X, Y int
}

func (d *Direction) turnRight() {
	if d.X == -1 && d.Y == 0 {
		d.X = 0
		d.Y = 1
	} else if d.X == 0 && d.Y == 1 {
		d.X = 1
		d.Y = 0
	} else if d.X == 1 && d.Y == 0 {
		d.X = 0
		d.Y = -1
	} else if d.X == 0 && d.Y == -1 {
		d.X = -1
		d.Y = 0
	}
}

type Location struct {
	X, Y int
}

func PartOne(filename string) {
	handler, err := common.FetchFile(filename)
	if err != nil {
		panic(err)
	}
	defer handler.Cleanup()

	grid := make([][]Object, 0)
	var guardX, guardY int
	rowPos := 0
	for {
		line, _, err := handler.Reader.ReadLine()
		if err != nil {
			break
		}

		row := make([]Object, 0)
		for colPos, item := range line {
			if item == '#' {
				row = append(row, Obstacle)
			} else if item == '.' {
				row = append(row, Free)
			} else if item == '^' {
				guardX, guardY = rowPos, colPos
				row = append(row, Free)
			}
		}
		grid = append(grid, row)
		rowPos += 1
	}

	visited := make(map[Location]bool)
	// fmt.Println(grid)

	direction := Direction{X: -1, Y: 0}
	for common.ValidGridIndex(guardX, guardY, len(grid), len(grid[0])) {
		if !common.ValidGridIndex(guardX, guardY, len(grid), len(grid[0])) {
			break
		}

		// fmt.Println(guardX, guardY, grid[guardX][guardY])
		visited[Location{guardX, guardY}] = true

		if common.ValidGridIndex(guardX+direction.X, guardY+direction.Y, len(grid), len(grid[0])) && grid[guardX+direction.X][guardY+direction.Y] == Obstacle {
			// fmt.Println("turn right")
			direction.turnRight()
		} else {
			guardX += direction.X
			guardY += direction.Y
		}
	}

	for loc, _ := range visited {
		grid[loc.X][loc.Y] = 2
	}

	fmt.Println(len(visited))
}

func checkLoop(grid [][]Object, newObstacleX, newObstacleY int, guardX, guardY int) bool {
	direction := Direction{X: -1, Y: 0}
	log := make([]Location, 0)
	visited := make(map[Location]int)
	for common.ValidGridIndex(guardX, guardY, len(grid), len(grid[0])) {
		if !common.ValidGridIndex(guardX, guardY, len(grid), len(grid[0])) {
			break
		}

		// fmt.Println(guardX, guardY)
		visited[Location{guardX, guardY}] += 1
		log = append(log, Location{guardX, guardY})

		if common.ValidGridIndex(guardX+direction.X, guardY+direction.Y, len(grid), len(grid[0])) && grid[guardX+direction.X][guardY+direction.Y] == Obstacle {
			// fmt.Println("turn right")
			direction.turnRight()
		} else if guardX+direction.X == newObstacleX && guardY+direction.Y == newObstacleY {
			// is a perceived obstacle
			direction.turnRight()
		} else {
			guardX += direction.X
			guardY += direction.Y
		}

		// if we cross a location more than 20 times, it's probably a loop
		// scuffed solution but it works like a charm!
		// I would prefer to figure it out by matching the logs but oh well

		if visited[Location{guardX, guardY}] >= 20 {
			return true
		}
	}
	return false
}

func PartTwo(filename string) {
	handler, err := common.FetchFile(filename)
	if err != nil {
		panic(err)
	}
	defer handler.Cleanup()

	grid := make([][]Object, 0)
	var guardX, guardY int
	rowPos := 0
	for {
		line, _, err := handler.Reader.ReadLine()
		if err != nil {
			break
		}

		row := make([]Object, 0)
		for colPos, item := range line {
			if item == '#' {
				row = append(row, Obstacle)
			} else if item == '.' {
				row = append(row, Free)
			} else if item == '^' {
				guardX, guardY = rowPos, colPos
				row = append(row, Free)
			}
		}
		grid = append(grid, row)
		rowPos += 1
	}

	count := 0

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if x == guardX && y == guardY {
				continue
			}
			if checkLoop(grid, x, y, guardX, guardY) {
				// found a loop
				count += 1
			}
		}
	}

	fmt.Println(count)
}

func main() {
	filename := common.GetProblemFile(6)
	fmt.Println("part one:")
	PartOne(filename)
	fmt.Println("part two:")
	PartTwo(filename)
}
