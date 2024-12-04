package cmd

import (
	"errors"
	"fmt"

	"github.com/jaredbancroft/aoc2024/internal/helpers"
	"github.com/spf13/cobra"
)

// day4Cmd represents the day1 command
var day4Cmd = &cobra.Command{
	Use:   "day4",
	Short: "Advent of Code 2024 - Day 4: Ceres Search",
	Long: `
--- Day 4: Ceres Search ---

	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		lines, err := helpers.ReadStringFile("inputs/day4.txt")
		if err != nil {
			return err
		}

		part1 := 0
		part2 := 0

		for j, line := range lines {
			for i, c := range line {
				if c == rune('X') {
					part1 += checkAll(i, j, lines)
				}
				if c == rune('A') {
					part2 += checkAll2(i, j, lines)
				}
			}
		}

		fmt.Println("Part 1 Total:", part1)
		fmt.Println("Part 2 Total:", part2)
		return nil
	},
}

func checkAll(x, y int, values []string) int {
	count := 0

	count += checkUp(x, y, values)
	count += checkDown(x, y, values)
	count += checkLeft(x, y, values)
	count += checkRight(x, y, values)
	count += checkUpLeft(x, y, values)
	count += checkUpRight(x, y, values)
	count += checkDownLeft(x, y, values)
	count += checkDownRight(x, y, values)

	return count
}

func checkAll2(x, y int, values []string) int {
	ul, err := checkUpLeft2(x, y, values)
	if err != nil {
		return 0
	}
	ur, err := checkUpRight2(x, y, values)
	if err != nil {
		return 0
	}
	dl, err := checkDownLeft2(x, y, values)
	if err != nil {
		return 0
	}
	dr, err := checkDownRight2(x, y, values)
	if err != nil {
		return 0
	}

	test := valid{ul, ur, dl, dr}

	valid1 := valid{77, 77, 83, 83}
	valid2 := valid{83, 83, 77, 77}
	valid3 := valid{77, 83, 77, 83}
	valid4 := valid{83, 77, 83, 77}

	if test == valid1 || test == valid2 || test == valid3 || test == valid4 {
		return 1
	}

	return 0
}

type valid struct {
	ul, ur, dl, dr byte
}

func checkUpLeft2(x, y int, values []string) (byte, error) {
	if y-1 < 0 || x-1 < 0 {
		return 0, errors.New("out of bounds")
	}
	if values[y-1][x-1] != 77 && values[y-1][x-1] != 83 {
		return 0, errors.New("not M or S")
	}

	return values[y-1][x-1], nil
}

func checkUpRight2(x, y int, values []string) (byte, error) {
	if y-1 < 0 || x+1 >= len(values[0]) {
		return 0, errors.New("out of bounds")
	}
	if values[y-1][x+1] != 77 && values[y-1][x+1] != 83 {
		return 0, errors.New("not M or S")
	}

	return values[y-1][x+1], nil
}

func checkDownLeft2(x, y int, values []string) (byte, error) {
	if y+1 >= len(values) || x-1 < 0 {
		return 0, errors.New("out of bounds")
	}
	if values[y+1][x-1] != 77 && values[y+1][x-1] != 83 {
		return 0, errors.New("not M or S")
	}

	return values[y+1][x-1], nil
}

func checkDownRight2(x, y int, values []string) (byte, error) {
	if y+1 >= len(values) || x+1 >= len(values[0]) {
		return 0, errors.New("out of bounds")
	}
	if values[y+1][x+1] != 77 && values[y+1][x+1] != 83 {
		return 0, errors.New("not M or S")
	}

	return values[y+1][x+1], nil
}

func checkUp(x, y int, values []string) int {
	for i := 0; i < 3; i++ {
		if y-(i+1) < 0 {
			return 0
		}

		if i == 0 && values[y-(i+1)][x] != 77 {
			return 0
		}

		if i == 1 && values[y-(i+1)][x] != 65 {
			return 0
		}

		if i == 2 && values[y-(i+1)][x] != 83 {
			return 0
		}
	}
	return 1
}

func checkDown(x, y int, values []string) int {
	for i := 0; i < 3; i++ {
		if y+(i+1) >= len(values) {
			return 0
		}

		if i == 0 && values[y+(i+1)][x] != 77 {
			return 0
		}

		if i == 1 && values[y+(i+1)][x] != 65 {
			return 0
		}

		if i == 2 && values[y+(i+1)][x] != 83 {
			return 0
		}
	}
	return 1
}

func checkLeft(x, y int, values []string) int {
	for i := 0; i < 3; i++ {
		if x-(i+1) < 0 {
			return 0
		}

		if i == 0 && values[y][x-(i+1)] != 77 {
			return 0
		}

		if i == 1 && values[y][x-(i+1)] != 65 {
			return 0
		}

		if i == 2 && values[y][x-(i+1)] != 83 {
			return 0
		}
	}
	return 1
}

func checkRight(x, y int, values []string) int {
	for i := 0; i < 3; i++ {
		if x+(i+1) >= len(values[0]) {
			return 0
		}

		if i == 0 && values[y][x+(i+1)] != 77 {
			return 0
		}

		if i == 1 && values[y][x+(i+1)] != 65 {
			return 0
		}

		if i == 2 && values[y][x+(i+1)] != 83 {
			return 0
		}
	}
	return 1
}

func checkUpLeft(x, y int, values []string) int {
	for i := 0; i < 3; i++ {
		if y-(i+1) < 0 || x-(i+1) < 0 {
			return 0
		}

		if i == 0 && values[y-(i+1)][x-(i+1)] != 77 {
			return 0
		}

		if i == 1 && values[y-(i+1)][x-(i+1)] != 65 {
			return 0
		}

		if i == 2 && values[y-(i+1)][x-(i+1)] != 83 {
			return 0
		}
	}
	return 1
}

func checkUpRight(x, y int, values []string) int {
	for i := 0; i < 3; i++ {
		if y-(i+1) < 0 || x+(i+1) >= len(values[0]) {
			return 0
		}

		if i == 0 && values[y-(i+1)][x+(i+1)] != 77 {
			return 0
		}

		if i == 1 && values[y-(i+1)][x+(i+1)] != 65 {
			return 0
		}

		if i == 2 && values[y-(i+1)][x+(i+1)] != 83 {
			return 0
		}
	}
	return 1
}

func checkDownLeft(x, y int, values []string) int {
	for i := 0; i < 3; i++ {
		if y+(i+1) >= len(values) || x-(i+1) < 0 {
			return 0
		}

		if i == 0 && values[y+(i+1)][x-(i+1)] != 77 {
			return 0
		}

		if i == 1 && values[y+(i+1)][x-(i+1)] != 65 {
			return 0
		}

		if i == 2 && values[y+(i+1)][x-(i+1)] != 83 {
			return 0
		}
	}
	return 1
}

func checkDownRight(x, y int, values []string) int {
	for i := 0; i < 3; i++ {
		if y+(i+1) >= len(values) || x+(i+1) >= len(values[0]) {
			return 0
		}

		if i == 0 && values[y+(i+1)][x+(i+1)] != 77 {
			return 0
		}

		if i == 1 && values[y+(i+1)][x+(i+1)] != 65 {
			return 0
		}

		if i == 2 && values[y+(i+1)][x+(i+1)] != 83 {
			return 0
		}
	}
	return 1
}

func init() {
	rootCmd.AddCommand(day4Cmd)
}
