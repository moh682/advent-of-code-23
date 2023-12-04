package day01

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/moh682/advent-of-code/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 1, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

type LetterNum struct {
	value int
	label string
	index int
}

func (l *LetterNum) AttachIndex(line string) {

	labelIndex := strings.Index(line, l.label)
	numberIndex := strings.Index(line, strconv.Itoa(l.value))

	if labelIndex == -1 && numberIndex == -1 {
		l.index = -1
	} else if labelIndex == -1 && numberIndex != -1 {
		l.index = numberIndex
	} else if labelIndex != -1 && numberIndex == -1 {
		l.index = labelIndex
	} else {
		if labelIndex < numberIndex {
			l.index = labelIndex
		} else {
			l.index = numberIndex
		}
	}

}

func (l *LetterNum) String() string {
	return fmt.Sprintf("value: %d, label: %s, index: %d", l.value, l.label, l.index)
}

func partB(input *challenge.Input) int {

	sum := 0
	for line := range input.Lines {

		AllNumbers := []*LetterNum{}
		labels := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

		for i, l := range labels {
			s := &LetterNum{value: i + 1, label: l}
			s.AttachIndex(line)
			AllNumbers = append(AllNumbers, s)
		}

		filteredSlice := make([]*LetterNum, 0)
		for _, num := range AllNumbers {
			if num.index != -1 {
				filteredSlice = append(filteredSlice, num)
			}
		}

		sort.Slice(filteredSlice, func(i, j int) bool {
			return filteredSlice[i].index < filteredSlice[j].index
		})

		if len(filteredSlice) == 0 {
			continue
		}

		n1 := filteredSlice[0]
		n2 := filteredSlice[len(filteredSlice)-1]

		temp := strconv.Itoa(n1.value) + strconv.Itoa(n2.value)
		localNum, err := strconv.Atoi(temp)
		if err != nil {
			panic(err)
		}
		sum += localNum
	}

	return sum
}
