package main

import (
	"common"
	"fmt"
	"strconv"
)

type QueueItem struct {
	X, Y int
}

func Bfs(grid [][]int, startX, startY int) int {
	queue := make([]QueueItem, 0)
	queue = append(queue, QueueItem{startX, startY})
	tailends := make(map[QueueItem]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if grid[current.X][current.Y] == 9 {
			tailends[current] = true
		} else {
			var i, j int
			i, j = current.X, current.Y-1
			if common.ValidGridIndex(i, j, len(grid), len(grid[0])) && grid[i][j] == grid[current.X][current.Y]+1 {
				queue = append(queue, QueueItem{i, j})
			}

			i, j = current.X, current.Y+1
			if common.ValidGridIndex(i, j, len(grid), len(grid[0])) && grid[i][j] == grid[current.X][current.Y]+1 {
				queue = append(queue, QueueItem{i, j})
			}

			i, j = current.X-1, current.Y
			if common.ValidGridIndex(i, j, len(grid), len(grid[0])) && grid[i][j] == grid[current.X][current.Y]+1 {
				queue = append(queue, QueueItem{i, j})
			}

			i, j = current.X+1, current.Y
			if common.ValidGridIndex(i, j, len(grid), len(grid[0])) && grid[i][j] == grid[current.X][current.Y]+1 {
				queue = append(queue, QueueItem{i, j})
			}
		}
	}

	return len(tailends)
}

func parseGrid(handler *common.InputHandler) [][]int {
	grid := make([][]int, 0)
	for {
		line, _, err := handler.Reader.ReadLine()
		if err != nil {
			break
		}

		ints := make([]int, 0)
		for _, digit := range line {
			integer, err := strconv.Atoi(string(digit))
			if err != nil {
				panic(err)
			}
			ints = append(ints, integer)
		}
		grid = append(grid, ints)
	}
	return grid
}

func PartOne(filename string) {
	handler, err := common.FetchFile(filename)
	if err != nil {
		panic(err)
	}
	defer handler.Cleanup()

	grid := parseGrid(&handler)
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				localCount := Bfs(grid, i, j)
				// fmt.Println(i, j, localCount)
				count += localCount
			}
		}
	}
	fmt.Println(count)
}

func Bfs2(grid [][]int, startX, startY int) int {
	queue := make([]QueueItem, 0)
	queue = append(queue, QueueItem{startX, startY})
	total := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if grid[current.X][current.Y] == 9 {
			total += 1
		} else {
			var i, j int
			i, j = current.X, current.Y-1
			if common.ValidGridIndex(i, j, len(grid), len(grid[0])) && grid[i][j] == grid[current.X][current.Y]+1 {
				queue = append(queue, QueueItem{i, j})
			}

			i, j = current.X, current.Y+1
			if common.ValidGridIndex(i, j, len(grid), len(grid[0])) && grid[i][j] == grid[current.X][current.Y]+1 {
				queue = append(queue, QueueItem{i, j})
			}

			i, j = current.X-1, current.Y
			if common.ValidGridIndex(i, j, len(grid), len(grid[0])) && grid[i][j] == grid[current.X][current.Y]+1 {
				queue = append(queue, QueueItem{i, j})
			}

			i, j = current.X+1, current.Y
			if common.ValidGridIndex(i, j, len(grid), len(grid[0])) && grid[i][j] == grid[current.X][current.Y]+1 {
				queue = append(queue, QueueItem{i, j})
			}
		}
	}

	return total
}

// funnily enough, part two is an easier version of part 1
func PartTwo(filename string) {
	handler, err := common.FetchFile(filename)
	if err != nil {
		panic(err)
	}
	defer handler.Cleanup()

	grid := parseGrid(&handler)
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				localCount := Bfs2(grid, i, j)
				// fmt.Println(i, j, localCount)
				count += localCount
			}
		}
	}
	fmt.Println(count)
}

func main() {
	filename := common.GetProblemFile(10)
	fmt.Println("part one:")
	PartOne(filename)
	fmt.Println("part two:")
	PartTwo(filename)
}
