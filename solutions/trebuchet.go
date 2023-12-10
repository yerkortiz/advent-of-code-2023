package solutions

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slog"
)

// returns the first digit in s multiplied by 10 plust the last digit in s
func Trebuchet(s string) int {
	first, last := 0, 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			if first == 0 {
				first = int(c-'0') * 10
			}
			last = int(c - '0')
		}
	}
	return first + last
}

func TrebuchetPart2(s string) int {
	numbers := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	first, firstIndex := 0, -1
	last, lastIndex := 0, -1

	for i, number := range numbers {
		index := strings.Index(s, number)
		if index == -1 {
			continue
		}

		if firstIndex == -1 {
			firstIndex = index
			lastIndex = index
			first = int(numbers[i%9][0] - '0')
			last = int(numbers[i%9][0] - '0')
		}

		if firstIndex > index {
			firstIndex = index
			first = int(numbers[i%9][0] - '0')
		}

		index = strings.LastIndex(s, number)
		if lastIndex < index {
			lastIndex = index
			last = int(numbers[i%9][0] - '0')
		}
	}

	return first*10 + last
}

func SolveTrebuchet() {
	inputFile, err := os.OpenFile("./data/trebuchet.in", os.O_RDONLY, os.ModePerm)
	if err != nil {
		slog.Error("trebuchet: open file error", "error", err)
		return
	}
	defer inputFile.Close()

	output := 0
	sc := bufio.NewScanner(inputFile)
	for sc.Scan() {
		input := sc.Text()
		output += TrebuchetPart2(input)
	}
	if err := sc.Err(); err != nil {
		slog.Error("trebuchet: scan file error", "error", err)
		return
	}

	outputFile, err := os.Create("./data/trebuchet2.out")
	if err != nil {
		slog.Error("trebuchet: create file error", "error", err)
		return
	}

	defer outputFile.Close()

	_, err = outputFile.WriteString(strconv.Itoa(output) + "\n")
	if err != nil {
		slog.Error("trebuchet: write file error", "error", err)
	}
}
