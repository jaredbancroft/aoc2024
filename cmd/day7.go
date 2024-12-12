package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jaredbancroft/aoc2024/internal/helpers"
	"github.com/spf13/cobra"
)

type Entry struct {
	target int
	values []int
}

// day7Cmd represents the day7 command
var day7Cmd = &cobra.Command{
	Use:   "day7",
	Short: "Advent of Code 2024 - Day 7: Bridge Repair",
	Long: `
--- Day 7: Bridge Repair ---

	`,
	RunE: func(cmd *cobra.Command, args []string) error {

		inputs, err := helpers.ReadStringFile("inputs/day7.txt")
		if err != nil {
			return err
		}

		var entries []Entry

		for _, input := range inputs {
			a := strings.Split(input, ":")
			for i := range a {
				a[i] = strings.TrimSpace(a[i])
			}
			b := strings.Split(a[1], " ")
			target, err := strconv.Atoi(a[0])
			if err != nil {
				return err
			}

			values := make([]int, 0, len(b))
			for _, numStr := range b {
				if numStr == "" {
					continue
				}
				v, err := strconv.Atoi(numStr)
				if err != nil {
					return err
				}
				values = append(values, v)
			}

			entries = append(entries, Entry{target: target, values: values})
		}

		part1 := 0
		part2 := 0

		for _, entry := range entries {
			if calculatePart1(entry.values[0], entry.values[1:], entry.target) {
				part1 += entry.target
			}
			if calculatePart2(entry.values[0], entry.values[1:], entry.target) {
				part2 += entry.target
			}
		}

		fmt.Println("Part 1 Total:", part1)
		fmt.Println("Part 2 Total:", part2)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(day7Cmd)
}

func calculatePart1(current int, remaining []int, target int) bool {
	if len(remaining) == 0 {
		return current == target
	}

	if calculatePart1(current+remaining[0], remaining[1:], target) {
		return true
	}

	return calculatePart1(current*remaining[0], remaining[1:], target)
}

func calculatePart2(current int, remaining []int, target int) bool {
	if len(remaining) == 0 {
		return current == target
	}

	next := remaining[0]
	rest := remaining[1:]

	if calculatePart2(current+next, rest, target) {
		return true
	}

	if calculatePart2(current*next, rest, target) {
		return true
	}

	concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", current, next))
	return calculatePart2(concatenated, rest, target)
}
