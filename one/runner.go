package one

import (
	"fmt"

	"github.com/moh682/advent-of-code/one/part1"
	"github.com/moh682/advent-of-code/one/part2"
	"github.com/moh682/advent-of-code/services"
)

func Run() {
	fs := services.NewFileService()
	lines, err := fs.ReadFile("one", "input.txt")
	if err != nil {
		panic(err)
	}

	sum1 := part1.Run(lines)
	fmt.Println("------------- day 1 -------------")
	fmt.Println("part 1:", sum1)

	sum2 := part2.Run(lines)
	fmt.Println("part 2:", sum2)
}
