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

	rules := make(map[int64][]int64)
	manuals := make([][]int64, 0)

	// rules section
	for {
		line, err := handler.GetDelimitedLine("|")
		if err != nil {
			break
		}

		if err != nil || len(line) == 0 {
			// we reached the end of this section, or we encountered EOF
			break
		}

		rule := common.StringsToInts(line)
		if _, ok := rules[rule[0]]; !ok {
			rules[rule[0]] = []int64{rule[1]}
		} else {
			rules[rule[0]] = append(rules[rule[0]], rule[1])
		}
	}

	// manual section
	for {
		line, err := handler.GetDelimitedLine(",")
		if err != nil {
			break
		}

		manuals = append(manuals, common.StringsToInts(line))
	}

	middleCount := int64(0)
	for _, manual := range manuals {
		positions := make(map[int64]int)
		for idx, m := range manual {
			positions[m] = idx
		}

		isValid := true
		for idx, m := range manual {
			if _, ok := rules[m]; ok {
				for _, nextPage := range rules[m] {
					if nextPagePosition, ok := positions[nextPage]; ok {
						if nextPagePosition <= idx {
							isValid = false
						}
					}
				}
			}
		}
		if isValid {
			middleCount += manual[len(manual)/2]
			// fmt.Println(manual)
		}
	}
	fmt.Println(middleCount)
}

func PartTwo(filename string) {
	handler, err := common.FetchFile(filename)
	if err != nil {
		panic(err)
	}
	defer handler.Cleanup()

	rules := make(map[int64][]int64)
	manuals := make([][]int64, 0)

	// rules section
	for {
		line, err := handler.GetDelimitedLine("|")
		if err != nil {
			break
		}

		if err != nil || len(line) == 0 {
			// we reached the end of this section, or we encountered EOF
			break
		}

		rule := common.StringsToInts(line)
		if _, ok := rules[rule[0]]; !ok {
			rules[rule[0]] = []int64{rule[1]}
		} else {
			rules[rule[0]] = append(rules[rule[0]], rule[1])
		}
	}

	// manual section
	for {
		line, err := handler.GetDelimitedLine(",")
		if err != nil {
			break
		}

		manuals = append(manuals, common.StringsToInts(line))
	}

	middleCount := int64(0)
	for _, manual := range manuals {
		positions := make(map[int64]int)
		for idx, m := range manual {
			positions[m] = idx
		}

		isValid := true
		for idx, m := range manual {
			if _, ok := rules[m]; ok {
				for _, nextPage := range rules[m] {
					if nextPagePosition, ok := positions[nextPage]; ok {
						if nextPagePosition <= idx {
							isValid = false
						}
					}
				}
			}
		}
		if !isValid {
			// sort the manual properly and add up the middle number
			slices.SortFunc(manual, func(m, n int64) int {
				if ruleset, ok := rules[m]; ok {
					for _, nextPage := range ruleset {
						if nextPage == n {
							return -1
						}
					}
				}
				if ruleset, ok := rules[n]; ok {
					for _, nextPage := range ruleset {
						if nextPage == m {
							return 1
						}
					}
				}
				return 0
			})
			// fmt.Println("sorted manual:", manual)
			middleCount += manual[len(manual)/2]
		}
	}
	fmt.Println(middleCount)
}

func main() {
	filename := common.GetProblemFile(5)
	fmt.Println("part one:")
	PartOne(filename)
	fmt.Println("part two:")
	PartTwo(filename)
}
