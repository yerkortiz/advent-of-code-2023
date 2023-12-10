package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCubeConundrum(t *testing.T) {
	games := []string{
		"Game 1: 15 blue;",
		"Game 2: 1 blue, 3 red; 5 blue, 5 green",
		"Game 3: 4 green, 1 blue; 6 blue, 5 green, 1 red; 11 green, 10 blue",
		"Game 4: 12 blue, 12 green, 3 red; 15 blue, 1 green, 10 red; 8 blue, 3 red, 2 green; 14 red, 8 blue",
		"Game 5: 7 blue, 8 red, 5 green; 15 blue, 16 red, 14 green; 3 blue, 14 red, 10 green",
	}

	outputs := []int{
		0, 2, 3, 0, 0,
	}

	for i, game := range games {
		assert.Equal(t, outputs[i], CubeConundrum(game))
	}
}

func TestCubeConundrumPart2(t *testing.T) {
	games := []string{
		"Game 1: 2 blue, 2 red, 2 green;",
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	}

	outputs := []int{
		8, 48,
	}

	for i, game := range games {
		assert.Equal(t, outputs[i], CubeConundrumPart2(game))
	}
}
