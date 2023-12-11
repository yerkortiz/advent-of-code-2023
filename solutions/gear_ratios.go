package solutions

import (
	"bufio"
	"os"
	"strconv"

	"golang.org/x/exp/slog"
)

func checkAdjacent(s []string, x int, y int) bool {
	var moves = [][]int{
		{-1, 0},  // up
		{1, 0},   // down
		{0, -1},  // left
		{0, 1},   // right
		{-1, -1}, // up-left
		{-1, 1},  // up-right
		{1, -1},  // down-left
		{1, 1},   // down-right
	}

	for _, move := range moves {
		if x+move[0] < 0 || x+move[0] >= len(s) {
			continue
		}

		if y+move[1] < 0 || y+move[1] >= len(s[0]) {
			continue
		}

		switch s[x+move[0]][y+move[1]] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
			continue
		default:
			return true
		}
	}

	return false
}

func GearRatios(s []string) int {
	sum := 0
	for x, line := range s {
		number := 0
		hasAdjacent := false
		for y, c := range line {
			if c >= '0' && c <= '9' {
				if number == 0 {
					number = int(c - '0')
				} else {
					number = number*10 + int(c-'0')
				}
			} else {
				if hasAdjacent {
					sum += number
				}
				hasAdjacent = false
				number = 0
				continue
			}

			if !hasAdjacent {
				hasAdjacent = checkAdjacent(s, x, y)
			}

			//border case
			if y == len(line)-1 {
				if hasAdjacent {
					sum += number
				}
				hasAdjacent = false
				number = 0
			}
		}
	}
	return sum
}

func SolveGearRatios() {
	inputFile, err := os.OpenFile("./data/gear_ratios.in", os.O_RDONLY, os.ModePerm)
	if err != nil {
		slog.Error("gear_ratios: open file error", "error", err)
		return
	}
	defer inputFile.Close()

	grid := make([]string, 0)
	sc := bufio.NewScanner(inputFile)
	for sc.Scan() {
		input := sc.Text()
		grid = append(grid, input)
	}

	output := GearRatios(grid)
	if err := sc.Err(); err != nil {
		slog.Error("gear_ratios: scan file error", "error", err)
		return
	}

	outputFile, err := os.Create("./data/gear_ratios.out")
	if err != nil {
		slog.Error("gear_ratios: create file error", "error", err)
		return
	}

	defer outputFile.Close()

	_, err = outputFile.WriteString(strconv.Itoa(output) + "\n")
	if err != nil {
		slog.Error("gear_ratios: write file error", "error", err)
	}
}
