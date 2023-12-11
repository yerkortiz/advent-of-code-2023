package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGearRatios(t *testing.T) {
	grid := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	expected := 4361
	actual := GearRatios(grid)
	assert.Equal(t, expected, actual)
}
