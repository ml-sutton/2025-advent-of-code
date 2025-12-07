package one

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ml-sutton/2025-advent-of-code/day-4/utils"
)

// problem rules
// If there are fewer than 4 rolls of paper in the 8 adjacent cells then you can access the roll.
// a simple ceres search on a matrix would work well

func Solve(path string) int {
	var total = 0
	var lines = loadFile(path)
	var matrix [][]string = populateMatrix(lines)
	for yIndex, row := range matrix {
		for xIndex, column := range row {
			if column == "@" {
				if checkAdjacentNodes(matrix, xIndex, yIndex) {
					total += 1
				}
			}
		}
	}
	return total
}

// helper methods are defined after this
func loadFile(path string) []string {
	var (
		file *os.File
		err  error
	)
	file, err = os.Open(path)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var buffer = bufio.NewReader(file)
	var lines []string = utils.ReadFile(buffer)
	file.Close()
	return lines
}

func populateMatrix(fileLines []string) [][]string {
	var matrix [][]string = [][]string{}
	for _, line := range fileLines {
		var row []string = []string{}
		for _, char := range line {
			row = append(row, string(char))
		}
		matrix = append(matrix, row)
	}
	return matrix
}
func checkAdjacentNodes(matrix [][]string, xCoordinate int, yCoordinate int) bool {
	var upOk bool = yCoordinate-1 >= 0
	var downOk bool = yCoordinate+1 <= len(matrix)-1
	var leftOk bool = xCoordinate-1 >= 0
	var rightOk bool = xCoordinate+1 <= len(matrix[yCoordinate])-1
	var numRolls = 0
	var found []string
	if upOk {
		found = append(found, matrix[yCoordinate-1][xCoordinate])
	}
	if downOk {
		found = append(found, matrix[yCoordinate+1][xCoordinate])
	}
	if leftOk {
		found = append(found, matrix[yCoordinate][xCoordinate-1])
	}
	if rightOk {
		found = append(found, matrix[yCoordinate][xCoordinate+1])
	}
	if upOk && leftOk {
		found = append(found, matrix[yCoordinate-1][xCoordinate-1])
	}
	if upOk && rightOk {
		found = append(found, matrix[yCoordinate-1][xCoordinate+1])
	}
	if downOk && leftOk {
		found = append(found, matrix[yCoordinate+1][xCoordinate-1])
	}
	if downOk && rightOk {
		found = append(found, matrix[yCoordinate+1][xCoordinate+1])
	}
	for _, item := range found {
		if item == "@" {
			numRolls += 1
		}
	}
	return numRolls < 4
}
