package tests

import (
	"testing"

	"github.com/ml-sutton/2025-advent-of-code/day-3/two"
)

func TestPartTwo(t *testing.T) {
	var expectedValue int64 = 3121910778619
	var actualValue int64 = two.PartTwo("test_input.txt")
	if !(expectedValue == actualValue) {
		t.Error("[TEST FAILED!] Expected result : ", expectedValue, " . Actual Result : ", actualValue)
	}
}
