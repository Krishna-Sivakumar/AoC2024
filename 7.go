package main

import (
	"common"
	"fmt"
	"math"
	"strconv"
	"sync"
)

type bfsNode struct {
	operatorStack []byte
	currentValue  int64
}

func bfs(expression []int64, targetValue int64) bool {
	queue := make([]bfsNode, 0, 100)
	queue = append(queue, bfsNode{operatorStack: []byte{}, currentValue: 0})

	for len(queue) > 0 {
		searchItem := queue[0]
		queue = queue[1:]
		// fmt.Println(string(searchItem.operatorStack), searchItem.currentValue, targetValue)

		if len(searchItem.operatorStack) == len(expression)-1 {
			if searchItem.currentValue == targetValue {
				// fmt.Println(targetValue, "found a match")
				return true
			}
		} else {
			plusNode := bfsNode{searchItem.operatorStack, searchItem.currentValue}
			starNode := bfsNode{searchItem.operatorStack, searchItem.currentValue}

			if len(starNode.operatorStack) == 0 {
				// fmt.Println('*', expression[len(searchItem.operatorStack)]*expression[len(searchItem.operatorStack)+1])
				starNode.currentValue = expression[len(searchItem.operatorStack)] * expression[len(searchItem.operatorStack)+1]
				plusNode.currentValue = expression[len(searchItem.operatorStack)] + expression[len(searchItem.operatorStack)+1]
			} else {
				starNode.currentValue *= expression[len(searchItem.operatorStack)+1]
				plusNode.currentValue += expression[len(searchItem.operatorStack)+1]
			}

			starNode.operatorStack = append(starNode.operatorStack, '*')
			plusNode.operatorStack = append(plusNode.operatorStack, '+')

			if searchItem.currentValue <= targetValue {
				queue = append(queue, plusNode)
			}
			if searchItem.currentValue <= targetValue {
				queue = append(queue, starNode)
			}
		}
	}

	return false
}

func combineNumbers(a, b int64) int64 {
	if a == 0 {
		return b
	} else {
		digits := int(math.Log10(float64(b))) + 1
		return int64(math.Pow10(digits))*a + b
	}
}

func bfsWithOr(expression []int64, targetValue int64) bool {
	queue := make([]bfsNode, 0, 100)
	queue = append(queue, bfsNode{operatorStack: []byte{}, currentValue: 0})

	for len(queue) > 0 {
		searchItem := queue[0]
		queue = queue[1:]

		if len(searchItem.operatorStack) == len(expression)-1 {
			if searchItem.currentValue == targetValue {
				return true
			}
		} else {
			plusNode := bfsNode{searchItem.operatorStack, searchItem.currentValue}
			starNode := bfsNode{searchItem.operatorStack, searchItem.currentValue}
			orNode := bfsNode{searchItem.operatorStack, searchItem.currentValue}

			if len(plusNode.operatorStack) == 0 {
				plusNode.currentValue = expression[len(searchItem.operatorStack)] + expression[len(searchItem.operatorStack)+1]
			} else {
				plusNode.currentValue += expression[len(searchItem.operatorStack)+1]
			}

			if len(starNode.operatorStack) == 0 {
				starNode.currentValue = expression[len(searchItem.operatorStack)] * expression[len(searchItem.operatorStack)+1]
			} else {
				starNode.currentValue *= expression[len(searchItem.operatorStack)+1]
			}

			if len(orNode.operatorStack) == 0 {
				orNode.currentValue = combineNumbers(expression[len(searchItem.operatorStack)], expression[len(searchItem.operatorStack)+1])
			} else {
				orNode.currentValue = combineNumbers(orNode.currentValue, expression[len(searchItem.operatorStack)+1])
			}

			plusNode.operatorStack = append(plusNode.operatorStack, '+')
			starNode.operatorStack = append(starNode.operatorStack, '*')
			orNode.operatorStack = append(orNode.operatorStack, '|')

			if plusNode.currentValue <= targetValue {
				queue = append(queue, plusNode)
			}
			if starNode.currentValue <= targetValue {
				queue = append(queue, starNode)
			}
			if orNode.currentValue <= targetValue {
				queue = append(queue, orNode)
			}
		}
	}

	return false
}

func PartOne(filename string) {
	handler, err := common.FetchFile(filename)
	if err != nil {
		panic(err)
	}
	defer handler.Cleanup()

	resultant := make(map[int]int64) // line number to resultant
	incompleteExpressions := make([][]int64, 0)

	lineNumber := 0
	for {
		line, err := handler.GetDelimitedLine(" ")
		if err != nil {
			break
		}

		lhs, err := strconv.ParseInt(line[0][:len(line[0])-1], 10, 64)
		if err != nil {
			panic(err)
		}

		resultant[lineNumber] = lhs

		incompleteExpressions = append(incompleteExpressions, common.StringsToInts(line[1:]))

		lineNumber++
	}

	// we bfs through each line
	count := 0
	for idx, expression := range incompleteExpressions {
		if bfs(expression, resultant[idx]) {
			count += int(resultant[idx])
		}
	}
	fmt.Println(count)
}

func PartTwo(filename string) {
	handler, err := common.FetchFile(filename)
	if err != nil {
		panic(err)
	}
	defer handler.Cleanup()

	resultant := make(map[int]int64) // line number to resultant
	incompleteExpressions := make([][]int64, 0)

	lineNumber := 0
	for {
		line, err := handler.GetDelimitedLine(" ")
		if err != nil {
			break
		}

		lhs, err := strconv.ParseInt(line[0][:len(line[0])-1], 10, 64)
		if err != nil {
			panic(err)
		}

		resultant[lineNumber] = lhs

		incompleteExpressions = append(incompleteExpressions, common.StringsToInts(line[1:]))

		lineNumber++
	}

	// we bfs through each line

	count := int64(0)
	results := make(chan int64, 850)
	var wg sync.WaitGroup

	segmentSize := 10

	for i := 0; i < len(incompleteExpressions); i += segmentSize {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var localCount int64 = 0
			for exprIdx := i; exprIdx < min(len(incompleteExpressions), i+segmentSize); exprIdx++ {
				if bfsWithOr(incompleteExpressions[exprIdx], resultant[exprIdx]) {
					localCount += resultant[exprIdx]
				}
			}
			results <- localCount
		}()
	}

	wg.Wait()
	close(results)

	for val := range results {
		count += val
	}

	fmt.Println(count)
}

func main() {
	filename := common.GetProblemFile(7)
	fmt.Println("part one:")
	PartOne(filename)
	fmt.Println("part two:")
	PartTwo(filename)
}
