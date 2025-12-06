package two

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ml-sutton/2025-advent-of-code/day-3/utils"
)

func PartTwo(path string) int64 {
	var (
		file *os.File
		err  error
	)
	file, err = os.Open(path)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var buffer *bufio.Reader = bufio.NewReader(file)
	var lines []string = utils.ReadFile(buffer)
	file.Close()
	var sum int64 = 0
	for _, line := range lines {
		sum += Solve(line)
	}

	return sum
}
