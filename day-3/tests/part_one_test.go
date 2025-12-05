package tests

import (
	"testing"

	"github.com/ml-sutton/2025-advent-of-code/day-3/one"
)

func TestPartOne(t *testing.T) {
	var expectedValue int64 = 357
	var actualValue int64 = one.PartOne("test_input.txt")
	if !(expectedValue == actualValue) {
		t.Error("[TEST FAILED!] Expected result : ", expectedValue, " . Actual Result : ", actualValue)
	}
}
