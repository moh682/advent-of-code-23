package part1

import (
	"strconv"
	"strings"
)

func Run(lines string) int {

	sum := 0
	for _, line := range strings.Split(lines, "\n") {

		numbers := make([]string, 0)
		for _, char := range strings.Split(line, "") {
			_, err := strconv.Atoi(char)
			if err != nil {
				continue
			}
			numbers = append(numbers, char)
		}

		n1 := numbers[0]
		n2 := numbers[len(numbers)-1]

		localNumStr := n1 + n2
		localNum, err := strconv.Atoi(localNumStr)
		if err != nil {
			panic(err)
		}

		sum += localNum
	}

	return sum
}
