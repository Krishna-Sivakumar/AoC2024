package main

import (
	"common"
	"fmt"
)

func PartOne(filename string) {
	handler, err := common.FetchFile(filename)
	if err != nil {
		panic(err)
	}
	defer handler.Cleanup()

	safeCount := 0
	for {
		numbers, err := handler.GetDelimitedLine(" ")
		if err != nil {
			break
		}
		report := common.StringsToInts(numbers)

		ascending := false
		descending := false

		var last int64 = -1
		unsafe := false
		for _, rep := range report {
			if last == -1 {
				last = rep
				continue
			}
			diff := rep - last
			if diff < 0 {
				diff = -diff
			}

			if diff > 3 || diff < 1 {
				// is unsafe
				unsafe = true
				break
			}

			if !ascending && !descending {
				if rep-last > 0 {
					ascending = true
				} else if rep-last < 0 {
					descending = true
				}
			} else if ascending && rep-last < 0 {
				unsafe = true
				break
			} else if descending && rep-last > 0 {
				unsafe = true
				break
			}

			last = rep
		}

		if !unsafe {
			safeCount += 1
		} else {
			// fmt.Println("report", report, "was unsafe")
		}
	}

	fmt.Println(safeCount)
}

/*
Check if a report is unsafe, but ignore the element marked by the parameter ignore
*/
func CheckIfUnsafe(report []int64, ignore int) bool {
	var last int64 = -1
	ascending := false
	descending := false
	unsafe := false
	for idx, rep := range report {
		if idx == ignore {
			// fmt.Println("we skip", idx, "with value", rep)
			continue
		}

		if last == -1 {
			last = rep
			continue
		}

		diff := rep - last
		if diff < 0 {
			diff = -diff
		}

		if diff > 3 || diff < 1 {
			// is unsafe
			unsafe = true
			break
		}

		if !ascending && !descending {
			if rep-last > 0 {
				ascending = true
			} else if rep-last < 0 {
				descending = true
			}
		} else if ascending && rep-last < 0 {
			unsafe = true
			break
		} else if descending && rep-last > 0 {
			unsafe = true
			break
		}

		last = rep
	}

	return unsafe
}

func PartTwo(filename string) {
	handler, err := common.FetchFile(filename)
	if err != nil {
		panic(err)
	}
	defer handler.Cleanup()

	safeCount := 0
	for {
		numbers, err := handler.GetDelimitedLine(" ")
		if err != nil {
			break
		}
		report := common.StringsToInts(numbers)

		ascending := false
		descending := false

		var last int64 = -1
		unsafe := false
		for _, rep := range report {
			if last == -1 {
				last = rep
				continue
			}
			diff := rep - last
			if diff < 0 {
				diff = -diff
			}

			if diff > 3 || diff < 1 {
				// is unsafe
				unsafe = true
				break
			}

			if !ascending && !descending {
				if rep-last > 0 {
					ascending = true
				} else if rep-last < 0 {
					descending = true
				}
			} else if ascending && rep-last < 0 {
				unsafe = true
				break
			} else if descending && rep-last > 0 {
				unsafe = true
				break
			}

			last = rep
		}

		if unsafe {
			// brute force of shame...
			for i := 0; i < len(report); i++ {
				// we just check if any version of the report is safe
				if !CheckIfUnsafe(report, i) {
					unsafe = false
					break
				}
			}
		}

		if !unsafe {
			safeCount += 1
		}
	}

	fmt.Println(safeCount)
}

func main() {
	filename := common.GetProblemFile(2)

	fmt.Println("part one:")
	PartOne(filename)
	fmt.Println("part two:")
	PartTwo(filename)
}
