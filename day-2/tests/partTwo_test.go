package test

import (
	"testing"

	"github.com/ml-sutton/2025-advent-of-code/day-2/parts/two"
)

func TestPartTwo(t *testing.T) {
	var expectedOutput int64 = 4174379265
	var actualOutput = two.PartTwo("test_input.txt")
	if !(expectedOutput == actualOutput) {
		t.Error("[TEST FAILED!] Expected result : ", expectedOutput, " . Actual Result : ", actualOutput)
	}
}
