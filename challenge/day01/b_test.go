package day01

import (
	"testing"

	"github.com/moh682/advent-of-code/challenge"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	t.Run("b", func(t *testing.T) {
		input := challenge.FromLiteral(`two1nine
		eightwothree
		abcone2threexyz
		xtwone3four
		4nineeightseven2
		zoneight234
		7pqrstsixteen`)

		result := partB(input)

		require.Equal(t, 281, result)
	})
}
