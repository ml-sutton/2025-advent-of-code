package one

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ml-sutton/2025-advent-of-code/day-5/utils"
)

type Range struct {
	low  int
	high int
}

func PartOne(path string) int {
	var fresh []int = []int{}
	var (
		ranges []Range
		ids    []int
	)
	ranges, ids = parseInput(path)
	for _, rng := range ranges {
		for _, id := range ids {
			if rng.low <= id && rng.high >= id {
				fresh = append(fresh, id)
			}
		}
	}

	return len(removeDuplicateInt(fresh))
}

func removeDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
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

func parseInput(path string) ([]Range, []int) {
	var ranges []Range = []Range{}
	var ids []int = []int{}
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
		} else {
			idAsInt, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
			}
			ids = append(ids, idAsInt)
		}
	}
	return ranges, ids

}
