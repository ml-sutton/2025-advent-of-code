package test

import (
	"testing"

	"github.com/ml-sutton/2025-advent-of-code/day-5/two"
)

func TestPartTwo(t *testing.T) {
	var expectedValue int = 14
	var actualValue int = two.PartTwo("test_input.txt")
	if !(expectedValue == actualValue) {
		t.Error("[TEST FAILED!] Expected result : ", expectedValue, " . Actual Result : ", actualValue)
	}
}
