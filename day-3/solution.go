package main

import (
	"fmt"

	"github.com/ml-sutton/2025-advent-of-code/day-3/one"
	"github.com/ml-sutton/2025-advent-of-code/day-3/two"
)

func main() {
	var partOneValue = one.PartOne("real_input.txt")
	fmt.Println("PART ONE : ", partOneValue)
	var partTwoValue = two.PartTwo("real_input.txt")
	fmt.Println("PART TWO : ", partTwoValue)
}
