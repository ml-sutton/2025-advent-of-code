package test

import (
	"testing"

	"github.com/ml-sutton/2025-advent-of-code/day-4/two"
)

func TestPartTwo(t *testing.T) {
	var expectedValue int = 43
	var actualValue int = two.Solve("test_input.txt")
	if !(expectedValue == actualValue) {
		t.Error("[TEST FAILED!] Expected result : ", expectedValue, " . Actual Result : ", actualValue)
	}

}
