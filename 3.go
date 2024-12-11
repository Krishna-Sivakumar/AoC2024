package main

import (
	"common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// quite the ugly set of solutions, these are
// just had to use regex to find the matching patterns.

// I use some common utilities here are there to make fetching files and reading inputs easier

func PartOne(filename string) {
	handler, err := common.FetchFile(filename)
	if err != nil {
		panic(err)
	}
	defer handler.Cleanup()

	lines := make([]string, 4)
	for {
		input, _, err := handler.Reader.ReadLine()
		if err != nil {
			break
		}
		lines = append(lines, string(input))
	}
	input := strings.Join(lines, "")
	// fmt.Println(input)

	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	results := re.FindAllStringSubmatchIndex(input, -1)
	sum := 0
	if results != nil {
		for _, result := range results {
			// fmt.Println(input[result[0]:result[1]])
			pairString := input[result[0]:result[1]]
			pair := pairString[4 : len(pairString)-1]
			intStrings := strings.Split(string(pair), ",")
			argOne, err := strconv.ParseInt(intStrings[0], 10, 64)
			if err != nil {
				panic(err)
			}
			argTwo, err := strconv.ParseInt(intStrings[1], 10, 64)
			if err != nil {
				panic(err)
			}
			sum += int(argOne * argTwo)
		}
	}
	fmt.Println(sum)
}

func mostRecent(matches [][]int, currentIndex int) (retVal int) {
	retVal = -1
	// fmt.Println(matches, currentIndex)
	for _, match := range matches {
		if match[0] < currentIndex {
			retVal = match[0]
		} else {
			break
		}
	}
	return
}

func PartTwo(filename string) {
	handler, err := common.FetchFile(filename)
	if err != nil {
		panic(err)
	}
	defer handler.Cleanup()

	lines := make([]string, 4)
	for {
		input, _, err := handler.Reader.ReadLine()
		if err != nil {
			break
		}
		lines = append(lines, string(input))
	}
	input := strings.Join(lines, "")
	// fmt.Println(input)

	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	dontre := regexp.MustCompile(`don't\(\)`)
	dore := regexp.MustCompile(`do\(\)`)

	results := re.FindAllStringSubmatchIndex(input, -1)
	donts := dontre.FindAllStringSubmatchIndex(input, -1)
	dos := dore.FindAllStringSubmatchIndex(input, -1)

	sum := 0

	if results != nil {
		for _, result := range results {
			pairString := input[result[0]:result[1]]
			pair := pairString[4 : len(pairString)-1]
			intStrings := strings.Split(string(pair), ",")
			argOne, err := strconv.ParseInt(intStrings[0], 10, 64)
			if err != nil {
				panic(err)
			}
			argTwo, err := strconv.ParseInt(intStrings[1], 10, 64)
			if err != nil {
				panic(err)
			}

			recentDo := mostRecent(dos, result[0])
			recentDont := mostRecent(donts, result[0])
			if recentDo == -1 && recentDont == -1 {
				// multiplication is enabled by default
				sum += int(argOne * argTwo)
			} else if recentDo > recentDont {
				// multiplication is enabled now, since there was a more recent do call than a don't.
				sum += int(argOne * argTwo)
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	filename := common.GetProblemFile(3)
	fmt.Println("part one:")
	PartOne(filename)
	fmt.Println("part two:")
	PartTwo(filename)
}
