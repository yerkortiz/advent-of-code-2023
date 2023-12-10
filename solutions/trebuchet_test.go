package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrebuchet(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "abc123",
			expected: 13,
		},
		{
			input:    "treb7uchet",
			expected: 77,
		},
	}

	for _, tc := range testCases {
		actual := Trebuchet(tc.input)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestTrebuchetPart2(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "abc123",
			expected: 13,
		},
		{
			input:    "treb7uchet",
			expected: 77,
		},
		{
			input:    "onecaca9",
			expected: 19,
		},
		{
			input:    "fivetwohlxdql43kfzzbhtncg",
			expected: 53,
		},
		{
			input:    "7one83sixthreecllxjnphb2",
			expected: 72,
		},
		{
			input:    "nine7one83sixthreecllxjnphb29",
			expected: 99,
		},
		{
			input:    "sixsixsix",
			expected: 66,
		},
	}

	for _, tc := range testCases {
		actual := TrebuchetPart2(tc.input)
		assert.Equal(t, tc.expected, actual)
	}
}
