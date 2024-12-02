package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc2024",
	Short: "Advent of Code 2024",
	Long: `
Advent of Code 2024

The Chief Historian is always present for the big Christmas sleigh launch, but nobody has seen him in months!
Last anyone heard, he was visiting locations that are historically significant to the North Pole; a group of 
Senior Historians has asked you to accompany them as they check the places they think he was most likely to visit.

As each location is checked, they will mark it on their list with a star. They figure the Chief Historian must be 
in one of the first fifty places they'll look, so in order to save Christmas, you need to help them get fifty stars 
on their list before Santa takes off on December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; 
the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.DisableSuggestions = true
}
