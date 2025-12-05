package one

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ml-sutton/2025-advent-of-code/day-2/utils"
)

func PartOne(path string) int64 {
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
	var inputIds []string = utils.ReadFile(buffer)
	file.Close()

	var sumOfInvalids int64 = Solve(inputIds)

	return sumOfInvalids
}
