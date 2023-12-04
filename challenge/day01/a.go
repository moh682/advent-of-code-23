package day01

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/moh682/advent-of-code/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 1, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(input *challenge.Input) int {

	sum := 0
	for line := range input.Lines {

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
