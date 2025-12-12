package two

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/ml-sutton/2025-advent-of-code/day-5/utils"
)

type Range struct {
	low  int
	high int
}

func PartTwo(path string) int {
	count := 0
	var (
		ranges []Range
	)
	ranges = parseInput(path)
	for _, rng := range mergeRanges(ranges) {
		count += rng.high - rng.low + 1
	}

	return count
}

func loadFile(path string) []string {
	var (
		file *os.File
		err  error
	)
	file, err = os.Open(path)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var buffer = bufio.NewReader(file)
	var lines []string = utils.ReadFile(buffer)
	file.Close()
	return lines
}
func mergeRanges(ranges []Range) []Range {
	if len(ranges) == 0 {
		return nil
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].low < ranges[j].low
	})

	merged := []Range{ranges[0]}

	for _, curr := range ranges[1:] {
		last := &merged[len(merged)-1]
		if curr.low <= last.high+1 {
			if curr.high > last.high {
				last.high = curr.high
			}
		} else {
			merged = append(merged, curr)
		}
	}

	return merged
}
func parseInput(path string) []Range {
	var ranges []Range = []Range{}
	var lines = loadFile(path)
	for _, line := range lines {
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			lowAsInt, lowErr := strconv.Atoi(parts[0])
			if lowErr != nil {
				fmt.Println(lowErr)
			}
			highAsInt, highErr := strconv.Atoi(parts[1])
			if highErr != nil {
				fmt.Println(highErr)
			}
			var newRange Range = Range{
				low:  lowAsInt,
				high: highAsInt,
			}
			ranges = append(ranges, newRange)
		}
	}
	return ranges

}
