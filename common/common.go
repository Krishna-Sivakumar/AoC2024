// A bunch of common utilities used throughout my solutions for Advent of Code 2024
// Author: Krishna Sivakumar

package common

import (
	"bufio"
	"flag"
	"os"
	"strconv"
	"strings"
)

/*
Encapsulates a file and a buffered reader.
Exposes a bunch of functions for easy input.
*/
type InputHandler struct {
	fileHandle *os.File
	Reader     *bufio.Reader
}

/*
Populates the InputHandler struct with a file handle and a Buffered Reader over the file.

filePath is the path to the file in question.

Remember to call Cleanup() once you're done parsing input.
*/
func FetchFile(filePath string) (InputHandler, error) {
	var handler InputHandler
	var err error

	handler.fileHandle, err = os.Open(filePath)
	if err != nil {
		return InputHandler{}, err
	}

	handler.Reader = bufio.NewReader(handler.fileHandle)

	return handler, nil
}

/*
Gets a list of strings delimited by spaces, with empty strings removed.

The Buffered Reader advances once due to this operation.
*/
func (handler *InputHandler) GetDelimitedLine(sep string) ([]string, error) {
	line, _, err := handler.Reader.ReadLine()
	if err != nil {
		return []string{}, err
	}
	strings := strings.Split(string(line), sep)

	results := make([]string, 0, len(strings))
	for _, str := range strings {
		if len(str) > 0 {
			results = append(results, str)
		}
	}

	return results, nil
}

/*
Returns the file's contents as one string.

Since the contents are read line-by-line, a separator to join the lines together needs to be provided.
*/
func (handler *InputHandler) GetAllContents(separator string) string {
	lines := make([]string, 0)
	for line, _, err := handler.Reader.ReadLine(); err == nil; line, _, err = handler.Reader.ReadLine() {
		lines = append(lines, string(line))
	}
	return strings.Join(lines, separator)
}

func (handler *InputHandler) GetAllContentsArray() [][]byte {
	result := make([][]byte, 0)
	for {
		line, _, err := handler.Reader.ReadLine()
		if err != nil {
			break
		}
		result = append(result, line)
	}
	return result
}

/*
Close the file descriptor properly.
*/
func (handler *InputHandler) Cleanup() {
	handler.fileHandle.Close()
}

/*
Converts an array of strings to integers.
*/
func StringsToInts(strings []string) (result []int64) {
	for _, plausibleNumber := range strings {
		parsed, err := strconv.ParseInt(plausibleNumber, 10, 64)
		if err != nil {
			panic(err)
		}
		result = append(result, parsed)
	}
	return
}

/*
Check if the coordinate (i,j) falls within the bounds of the grid of a certain length and width.
The length is the size of the outer array of the grid.
The width is the size of an inner array within the grid.
*/
func ValidGridIndex(i, j int, length, width int) bool {
	return !(i < 0 || j < 0 || i >= length || j >= width)
}

/*
Rotates an array by n steps.
*/
func Rotate[E any](array []E, steps int) []E {
	steps = steps % len(array)
	newArray := make([]E, 0)
	newArray = append(newArray, array[len(array)-steps:]...)
	newArray = append(newArray, array[:len(array)-steps]...)
	return newArray
}

/*
Gets the problem set file for the easy version of the problem or the hard version of the problem.

problemNumber is the index of the problem.
*/
func GetProblemFile(problemNumber int) string {
	easyFlag := flag.Bool("easy", false, "Test the easy version of the problem.")
	flag.Parse()

	if *easyFlag {
		return "inputs/" + strconv.Itoa(problemNumber) + "a.in"
	} else {
		return "inputs/" + strconv.Itoa(problemNumber) + "b.in"
	}
}

/*
Gets the problem set file for the easy version of the problem or the hard version of the problem.

tag is the name of the problem.

Example:
file, err := GetProblemFileWithTag("19-2015") // fetches inputs/19-2015a.in, inputs/19-2015b.in
*/
func GetProblemFileWithTag(tag string) string {
	easyFlag := flag.Bool("easy", false, "Test the easy version of the problem.")
	flag.Parse()

	if *easyFlag {
		return "inputs/" + tag + "a.in"
	} else {
		return "inputs/" + tag + "b.in"
	}
}
