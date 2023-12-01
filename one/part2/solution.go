package part2

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func convertToNumber(char string) int {
	switch char {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return 0
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

func Run(lines string) int {

	sum := 0
	for _, line := range strings.Split(lines, "\n") {

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

		fmt.Println(n1.String())
		fmt.Println(n2.String())

		temp := strconv.Itoa(n1.value) + strconv.Itoa(n2.value)
		localNum, err := strconv.Atoi(temp)
		if err != nil {
			panic(err)
		}
		sum += localNum

		fmt.Println("-----------------------")
		fmt.Println("line:", line)
		fmt.Println(n1.String())
		fmt.Println(n2.String())
		fmt.Println("sum:", localNum)

	}

	return sum

}
