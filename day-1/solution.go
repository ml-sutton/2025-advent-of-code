package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func getNewDialPositionPartTwo(position int, direction string, command int) (int, int) {
	passed := 0
	for i := 0; i < command; i++ {
		if direction == "R" {
			position = (position + 1) % 100
		} else {
			position = (position - 1 + 100) % 100
		}
		if position == 0 {
			passed++
		}
	}
	return passed, position
}
func getNewDialPositionPartOne(position int, direction string, command int) int {
	var newDialPosition int
	if direction == "R" {
		newDialPosition = (position + command) % 100
	}
	if direction == "L" {
		if 0 <= command || command < 100 {
			newDialPosition = (position - command + 100) % 100
		} else {
			newDialPosition = ((position - command) + 10000) % 100
		}
	}
	return newDialPosition
}
func partTwo(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	Regex := regexp.MustCompile(`[a-zA-Z]+|[0-9]+`)
	timesAtZero := 0
	dialPosition := 50
	direction := "nil"
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if len(line) > 0 {
			stringParts := Regex.FindAllString(line, -1)
			fmt.Println(stringParts)
			direction = stringParts[0]
			parsedInt, error := strconv.Atoi(stringParts[1])
			if error != nil {
				fmt.Println(error)
			}
			if direction == "R" {
				hits, newDialPosition := getNewDialPositionPartTwo(dialPosition, direction, parsedInt)
				timesAtZero = timesAtZero + hits
				dialPosition = newDialPosition
			}
			if direction == "L" {
				hits, newDialPosition := getNewDialPositionPartTwo(dialPosition, direction, parsedInt)
				timesAtZero = timesAtZero + hits
				dialPosition = newDialPosition
			}
		}
		if err != nil {
			break
		}
	}
	fmt.Println(timesAtZero)
}
func partOne(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	Regex := regexp.MustCompile(`[a-zA-Z]+|[0-9]+`)
	timesAtZero := 0
	dialPosition := 50
	newDialPosition := dialPosition
	direction := "nil"
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		stringParts := Regex.FindAllString(line, -1)
		direction = stringParts[0]
		parsedInt, error := strconv.Atoi(stringParts[1])
		if error != nil {
			fmt.Println(error)
		}
		if direction == "L" {
			newDialPosition = getNewDialPositionPartOne(dialPosition, stringParts[0], parsedInt)
			if newDialPosition == 0 {
				timesAtZero++
			}
			dialPosition = newDialPosition
		}
		if direction == "R" {
			newDialPosition = getNewDialPositionPartOne(dialPosition, stringParts[0], parsedInt)
			if newDialPosition == 0 {
				timesAtZero++
			}
			dialPosition = newDialPosition
		}
	}
	fmt.Println(timesAtZero)
}
func main() {
	partOne("real_input.txt")
	partTwo("real_input.txt")
}
