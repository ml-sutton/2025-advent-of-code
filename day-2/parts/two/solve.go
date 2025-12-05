package two

import (
	"fmt"
	"strings"

	"github.com/ml-sutton/2025-advent-of-code/day-2/utils"
)

func Solve(ids []string) int64 {
	var sum int64 = 0
	for _, id := range ids {
		var (
			leftSide  string
			rightSide string
		)
		var parts []string = strings.Split(id, "-")
		leftSide, rightSide = parts[0], strings.Split(parts[1], ",")[0]
		var (
			leftSideInt  int64
			rightSideInt int64
		)
		leftSideInt = utils.ParseInt(leftSide)
		rightSideInt = utils.ParseInt(rightSide)

		for id := leftSideInt; id <= rightSideInt; id++ {
			var strId string = fmt.Sprint(id)
			var length int = len(strId)
			var doubled string = fmt.Sprint(strId, strId)
			var slice = doubled[1 : 2*length-1]
			if strings.Contains(slice, strId) {
				sum += id
			}
		}
	}

	return sum

}
