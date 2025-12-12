package main

import (
	"fmt"

	"github.com/ml-sutton/2025-advent-of-code/day-4/one"
	"github.com/ml-sutton/2025-advent-of-code/day-4/two"
)

func main() {
	var partOneValue = one.Solve("real_input.txt")
	fmt.Println("PART ONE : ", partOneValue)
	var partTwoValue = two.Solve("real_input.txt")
	fmt.Println("PART ONE : ", partTwoValue)
}
