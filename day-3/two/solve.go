package two

import (
	"fmt"

	"github.com/ml-sutton/2025-advent-of-code/day-3/utils"
)

func Solve(line string) int64 {
	var stringNum string = ""
	var k int = 12
	var numbers []int64 = []int64{}
	var solution []int64 = []int64{}
	for _, number := range line {
		numbers = append(numbers, utils.ParseInt(string(number)))
	}
	var length int = len(numbers)
	var R int = length - k
	for _, number := range numbers {
		for R > 0 && len(solution) > 0 && solution[len(solution)-1] < number {
			solution = solution[:len(solution)-1]
			R--
		}
		solution = append(solution, number)

	}
	if R > 0 {
		solution = solution[:len(solution)-R]
	}
	for _, number := range solution[:k] {
		stringNum = fmt.Sprintf("%s%d", stringNum, number)
	}
	return utils.ParseInt(stringNum)
}
