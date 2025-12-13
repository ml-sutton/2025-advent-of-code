package utils

import (
	"bufio"
	"strings"
)

func ReadFile(reader *bufio.Reader) []string {
	var lines []string = []string{}
	var (
		line []byte
		err  error
	)
	for {
		line, _, err = reader.ReadLine()
		if err != nil {
			break
		}
		var trimmedLine string = strings.TrimSpace(string(line))
		if trimmedLine == "" {
			continue
		}
		lines = append(lines, trimmedLine)
	}
	return lines
}
