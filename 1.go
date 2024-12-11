package main

import (
	"common"
	"fmt"
	"slices"
)

func PartOne(filename string) {
	handler, err := common.FetchFile(filename)
	if err != nil {
		panic(err)
	}
	defer handler.Cleanup()

	leftList, rightList := make([]int64, 0), make([]int64, 0)
	for {
		numbers, err := handler.GetDelimitedLine(" ")
		if err != nil {
			break
		}

		leftAndRight := common.StringsToInts(numbers)

		leftList = append(leftList, leftAndRight[0])
		rightList = append(rightList, leftAndRight[1])
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	differences := int64(0)
	for i := 0; i < len(leftList); i++ {
		diff := leftList[i] - rightList[i]
		if diff > 0 {
			differences += diff
		} else {
			differences += -diff
		}
	}
	fmt.Println("differences added up:", differences)
}

func PartTwo(filename string) {
	handler, err := common.FetchFile(filename)
	if err != nil {
		panic(err)
	}
	defer handler.Cleanup()

	leftList, rightList := make([]int64, 0), make([]int64, 0)
	for {
		numbers, err := handler.GetDelimitedLine(" ")
		if err != nil {
			break
		}
		leftAndRight := common.StringsToInts(numbers)

		leftList = append(leftList, leftAndRight[0])
		rightList = append(rightList, leftAndRight[1])
	}

	countMap := make(map[int64]int)
	for _, right := range rightList {
		countMap[right] += 1
	}

	score := int64(0)
	for _, left := range leftList {
		score += left * int64(countMap[left])
	}
	fmt.Println("total score:", score)
}

func main() {
	filename := common.GetProblemFile(1)

	fmt.Println("part one:")
	PartOne(filename)
	fmt.Println("part two:")
	PartTwo(filename)
}
