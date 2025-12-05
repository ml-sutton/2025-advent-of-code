package one

import (
	"fmt"

	"github.com/ml-sutton/2025-advent-of-code/day-3/utils"
)

func Solve(line string) int64 {
	var numbers []int64 = []int64{}
	var max int64 = 0
	var maxIndex int = -1
	var secondNum int64 = 0
	for _, number := range line {
		numbers = append(numbers, utils.ParseInt(string(number)))
	}
	for index, number := range numbers {
		if number > max {
			max = number
			maxIndex = index
		}
		if index+1 == len(numbers)-1 {
			break
		}
	}
	for index, number := range numbers {

		if number > secondNum && index > maxIndex {
			secondNum = number
		}
	}
	var joltageAsString = fmt.Sprintf("%d%d", max, secondNum)
	return utils.ParseInt(joltageAsString)
}
