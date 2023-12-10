package solutions

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slog"
)

const (
	maxBlue  = 14
	maxRed   = 12
	maxGreen = 13
)

func CubeConundrum(s string) int {
	gameID := 0
	temp := 1
	for i := 5; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			gameID *= temp
			gameID += int(s[i] - '0')
			temp *= 10
		} else {
			break
		}
	}

	sets := strings.Split(s[7:], ";")
	colours := []string{"blue", "red", "green"}

	var blue, red, green int
	for _, set := range sets {
		for _, colour := range colours {
			index := strings.Index(set, colour)
			if index != -1 {
				index -= 2
				temp = 1
				for set[index] >= '0' && set[index] <= '9' {
					switch colour {
					case "blue":
						blue += int(set[index]-'0') * temp
					case "red":
						red += int(set[index]-'0') * temp
					case "green":
						green += int(set[index]-'0') * temp
					}
					index--
					temp *= 10
				}
			}
		}

		if blue > maxBlue || red > maxRed || green > maxGreen {
			return 0
		}
		blue, red, green = 0, 0, 0
	}

	return gameID
}

func CubeConundrumPart2(s string) int {
	sets := strings.Split(s[7:], ";")
	colours := []string{"blue", "red", "green"}

	minBlue, minRed, minGreen := 0, 0, 0
	for _, set := range sets {
		for _, colour := range colours {
			index := strings.Index(set, colour)
			if index != -1 {
				index -= 2
				var blue, red, green int
				temp := 1
				for set[index] >= '0' && set[index] <= '9' {
					switch colour {
					case "blue":
						blue += int(set[index]-'0') * temp
					case "red":
						red += int(set[index]-'0') * temp
					case "green":
						green += int(set[index]-'0') * temp
					}
					index--
					temp *= 10
				}

				switch colour {
				case "blue":
					if minBlue < blue {
						minBlue = blue
					}
				case "red":
					if minRed < red {
						minRed = red
					}
				case "green":
					if minGreen < green {
						minGreen = green
					}
				}
			}
		}
	}

	return minBlue * minGreen * minRed
}

func SolveCubeConundrum() {
	inputFile, err := os.OpenFile("./data/cube_conundrum.in", os.O_RDONLY, os.ModePerm)
	if err != nil {
		slog.Error("cube_conundrum: open file error", "error", err)
		return
	}
	defer inputFile.Close()

	output := 0
	sc := bufio.NewScanner(inputFile)
	for sc.Scan() {
		input := sc.Text()
		output += CubeConundrumPart2(input)
	}
	if err := sc.Err(); err != nil {
		slog.Error("cube_conundrum: scan file error", "error", err)
		return
	}

	outputFile, err := os.Create("./data/cube_conundrum2.out")
	if err != nil {
		slog.Error("cube_conundrum: create file error", "error", err)
		return
	}

	defer outputFile.Close()

	_, err = outputFile.WriteString(strconv.Itoa(output) + "\n")
	if err != nil {
		slog.Error("cube_conundrum: write file error", "error", err)
	}
}
