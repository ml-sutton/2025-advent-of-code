package main

import (
	"github.com/ml-sutton/2025-advent-of-code/day-2/parts/one"
	"github.com/ml-sutton/2025-advent-of-code/day-2/parts/two"
)

// ------------------------------ ID RULES ------------------------------
// Since the young Elf was just doing silly patterns,
// you can find the invalid IDs by looking for any ID,
// which is made only of some sequence of digits repeated twice.
// So, 55 (5 twice), 6464 (64 twice), and 123123 (123 twice)
// would all be invalid IDs. None of the numbers have leading zeroes;
// 0101 isn't an ID at all. (101 is a valid ID that you would ignore.)

// Algorithm find range -> binary search -> find invalid

// ----------------------------------------------------------------------

func main() {
	var partOneValue int64 = one.PartOne("real_input.txt")
	println("part 1 : ", partOneValue)
	var partTwoValue int64 = two.PartTwo("real_input.txt")
	println("part 2 : ", partTwoValue)
}
