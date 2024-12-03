package cmd

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jaredbancroft/aoc2024/internal/helpers"
	"github.com/spf13/cobra"
)

// day3Cmd represents the day1 command
var day3Cmd = &cobra.Command{
	Use:   "day3",
	Short: "Advent of Code 2024 - Day 3: Mull It Over",
	Long: `
--- Day 3: Mull It Over ---

	`,
	RunE: func(cmd *cobra.Command, args []string) error {

		inputs, err := helpers.ReadStringFile("inputs/day3.txt")
		if err != nil {
			return err
		}

		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don\'t\(\)`)

		part1 := 0
		part2 := 0
		enabled := true

		for _, input := range inputs {
			matches := re.FindAllStringSubmatch(input, -1)

			if len(matches) > 0 {
				for _, match := range matches {
					if match[0] == "do()" {
						enabled = true
						continue
					}

					if match[0] == "don't()" {
						enabled = false
						continue
					}

					X, err := strconv.Atoi(match[1])
					if err != nil {
						return err
					}
					Y, err := strconv.Atoi(match[2])
					if err != nil {
						return err
					}

					part1 += X * Y

					if enabled {
						part2 += X * Y
					}
				}
			} else {
				fmt.Println("No match found")
			}
		}

		fmt.Println("Part 1 Total:", part1)
		fmt.Println("Part 2 Total:", part2)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(day3Cmd)
}
