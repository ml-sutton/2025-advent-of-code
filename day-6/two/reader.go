package two

import (
	"strconv"
	"strings"

	"github.com/ml-sutton/2025-advent-of-code/day-6/models"
)

func Reader(ingestBus chan models.Problem, lines []string) {
	var tempMatrix [][]string = [][]string{}
	for _, line := range lines {
		var parts = strings.Split(line, " ")
		tempMatrix = append(tempMatrix, parts)
	}
	for i, row := range tempMatrix {
		var filtered []string = []string{}
		for _, column := range row {
			if column != "" {
				filtered = append(filtered, column)
			}
		}
		tempMatrix[i] = filtered
	}
	for i := 0; i <= len(tempMatrix[0])-1; i++ {
		var problem models.Problem
		for j := 0; j <= len(tempMatrix)-1; j++ {
			itemAsInt, err := strconv.Atoi(tempMatrix[j][i])
			if err != nil {
				problem.Operator = tempMatrix[j][i]
			} else {
				problem.Numbers = append(problem.Numbers, itemAsInt)
			}
		}

		ingestBus <- problem
	}
	close(ingestBus)
}
