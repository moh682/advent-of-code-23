package day01

import (
	"testing"

	"github.com/moh682/advent-of-code/challenge"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	t.Run("a", func(t *testing.T) {
		input := challenge.FromLiteral(`1abc2
		pqr3stu8vwx
		a1b2c3d4e5f
		treb7uchet`)

		result := partA(input)

		require.Equal(t, 142, result)
	})
}
