package one

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ml-sutton/2025-advent-of-code/day-6/models"
	"github.com/ml-sutton/2025-advent-of-code/day-6/utils"
)

func PartOne(path string) int {
	var lines = loadFile(path)
	var ingestBus chan models.Problem = make(chan models.Problem)
	var egressBus chan int = make(chan int)
	go Reader(ingestBus, lines)
	go Consumer(ingestBus, egressBus)
	for item := range egressBus {
		return item
	}

	return -1
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
