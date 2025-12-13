package two

import "github.com/ml-sutton/2025-advent-of-code/day-6/models"

func Consumer(ingestBus chan models.Problem, egressBus chan int) {
	var total = 0
	for problem := range ingestBus {
		var localTotal = 0
		if problem.Operator == "*" {
			for index, number := range problem.Numbers {
				if index == 0 {
					localTotal = number
				} else {
					localTotal *= number
				}
			}
		} else {
			for _, number := range problem.Numbers {
				localTotal += number
			}
		}
		total += localTotal
	}
	egressBus <- total
	close(egressBus)
}
