package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jaredbancroft/aoc2024/internal/helpers"
	"github.com/spf13/cobra"
)

// day1Cmd represents the day1 command
var day2Cmd = &cobra.Command{
	Use:   "day2",
	Short: "Advent of Code 2024 - Day 2: Red-Nosed Reports",
	Long: `
--- Day 2: Red-Nosed Reports ---
Fortunately, the first location The Historians want to search isn't a long walk from the Chief Historian's office.

While the Red-Nosed Reindeer nuclear fusion/fission plant appears to contain no sign of the Chief Historian, the 
engineers there run up to you as soon as they see you. Apparently, they still talk about the time Rudolph was saved 
through molecular synthesis from a single electron.

They're quick to add that - since you're already here - they'd really appreciate your help analyzing some unusual 
data from the Red-Nosed reactor. You turn to check if The Historians are waiting for you, but they seem to have 
already divided into groups that are currently searching every corner of the facility. You offer to help with the unusual data.

The unusual data (your puzzle input) consists of many reports, one report per line. Each report is a list of numbers 
called levels that are separated by spaces. For example:

7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
This example data contains six reports each containing five levels.

The engineers are trying to figure out which reports are safe. The Red-Nosed reactor safety systems 
can only tolerate levels that are either gradually increasing or gradually decreasing. So, a report only 
counts as safe if both of the following are true:

The levels are either all increasing or all decreasing.
Any two adjacent levels differ by at least one and at most three.
In the example above, the reports can be found safe or unsafe by checking those rules:

7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.
So, in this example, 2 reports are safe.

Analyze the unusual data from the engineers. How many reports are safe?
	`,
	RunE: func(cmd *cobra.Command, args []string) error {

		reports, err := helpers.ReadStringFile("inputs/day2.txt")
		if err != nil {
			return err
		}

		var safeTotal1 int = 0
		var safeTotal2 int = 0

		for _, levels := range reports {
			stringLevels := strings.Split(levels, " ")
			intLevels := make([]int, len(stringLevels))

			for i, level := range stringLevels {
				intLevels[i], err = strconv.Atoi(level)
				if err != nil {
					return err
				}
			}

			safe, _ := part1(intLevels)
			safeTotal1 += safe
			safeTotal2 = safeTotal2 + part2(intLevels)
		}
		fmt.Println("Safe Reports Part 1: ", safeTotal1)
		fmt.Println("Safe Reports Part 2: ", safeTotal2)
		return nil
	},
}

func compareLevels(a int, b int) int {
	if a-b > 0 && a-b <= 3 {
		return 1
	}

	if a-b < 0 && a-b >= -3 {
		return -1
	}

	return 0
}

func part1(lvl []int) (int, int) {
	trend := 0

	for i := 0; i < len(lvl)-1; i++ {
		if compareLevels(lvl[i], lvl[i+1]) == 0 {
			return 0, i
		}

		trend = trend + compareLevels(lvl[i], lvl[i+1])

		if trend != i+1 && trend != 0-(i+1) {
			return 0, i
		}
	}

	return 1, 0
}

func part2(lvl []int) int {
	safe, idx := part1(lvl)

	if safe != 1 {
		a1 := lvl[:idx]
		b1 := lvl[idx+1:]
		a2 := lvl[:idx+1]
		b2 := lvl[idx+2:]

		retest1 := append([]int{}, a1...)
		retest1 = append(retest1, b1...)

		retest2 := append([]int{}, a2...)
		retest2 = append(retest2, b2...)

		r1, _ := part1(retest1)
		r2, _ := part1(retest2)
		r3, _ := part1(lvl[1:])

		fmt.Println(retest1, retest2, r1, r2, r3)

		if r1 == 1 || r2 == 1 || r3 == 1 {
			return 1
		}

		return 0
	}

	return 1
}

func init() {
	rootCmd.AddCommand(day2Cmd)
}
