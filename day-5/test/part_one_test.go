package test

import (
	"testing"

	"github.com/ml-sutton/2025-advent-of-code/day-5/one"
)

func TestPartOne(t *testing.T) {
	var expectedValue int = 3
	var actualValue int = one.PartOne("test_input.txt")
	if !(expectedValue == actualValue) {
		t.Error("[TEST FAILED!] Expected result : ", expectedValue, " . Actual Result : ", actualValue)
	}
}
