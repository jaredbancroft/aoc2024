package cmd

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jaredbancroft/aoc2024/internal/helpers"
	"github.com/spf13/cobra"
)

// day5Cmd represents the day1 command
var day5Cmd = &cobra.Command{
	Use:   "day5",
	Short: "Advent of Code 2024 - Day 5: Print Queue",
	Long: `
--- Day 5: Print Queue ---

	`,
	RunE: func(cmd *cobra.Command, args []string) error {

		inputs, err := helpers.ReadGroupStringFile("inputs/day5.txt")
		if err != nil {
			return err
		}

		rules := map[string][]string{}
		updates := [][]string{}

		for i, input := range inputs {
			if i == 0 {
				for _, rule := range input {
					values := strings.Split(rule, "|")
					rules[values[0]] = append(rules[values[0]], values[1])
				}
			}

			if i == 1 {
				for _, update := range input {
					updates = append(updates, strings.Split(update, ","))
				}
			}
		}

		part1 := 0
		part2 := 0

		for _, update := range updates {
			valid := true
			for i := 0; i < len(update); i++ {
				number := update[i]
				for _, rule := range rules[number] {
					previous_nums := update[:i]
					index := slices.Index(previous_nums, rule)
					if index != -1 {
						update = slices.Delete(update, index, index+1)
						update = slices.Insert(update, i, rule)
						valid = false
						i = 0
						break
					}
				}
			}
			val, _ := strconv.Atoi(update[len(update)/2])
			if valid {
				part1 += val
			} else {
				part2 += val
			}
		}

		fmt.Println("Part 1 Total:", part1)
		fmt.Println("Part 2 Total:", part2)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(day5Cmd)
}
