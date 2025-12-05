package test

import (
	"testing"

	"github.com/ml-sutton/2025-advent-of-code/day-2/parts/one"
)

func TestPartOne(t *testing.T) {
	var expectedOutput int64 = 1227775554
	var actualOutput = one.PartOne("test_input.txt")
	if !(expectedOutput == actualOutput) {
		t.Error("[TEST FAILED!] Expected result : ", expectedOutput, " . Actual Result : ", actualOutput)
	}
}
