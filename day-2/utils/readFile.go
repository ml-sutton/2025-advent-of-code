package utils

import (
	"bufio"
	"strings"
)

func ReadFile(reader *bufio.Reader) []string {
	var idSlice []string
	var (
		line []byte
		err  error
	)
	for {
		line, _, err = reader.ReadLine()
		if err != nil {
			break
		}
		var trimmedLine string
		trimmedLine = strings.TrimSpace(string(line))
		if trimmedLine == "" {
			continue
		}
		// Split the line on commas
		parts := strings.Split(trimmedLine, ",")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if part != "" {
				idSlice = append(idSlice, part)
			}
		}
	}
	return idSlice
}
