package helpers

import (
	"os"
	"strconv"
	"strings"
)

// ReadIntFile will read a file of integers separated by a newline
// and return a slice of integers
func ReadIntFile(fname string) ([]int, error) {
	b, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	nums := make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}

		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

// ReadStringFile will read a file of strings separated by a newline
// and return a slice of strings
func ReadStringFile(fname string) ([]string, error) {
	b, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")

	return lines, nil
}

// ReadGroupStringFile will read a file of strings separated by a newline
// and return a slice of strings
func ReadGroupStringFile(fname string) (map[int][]string, error) {
	g := make(map[int][]string)
	b, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	groupNum := 0
	groupVals := []string{}
	for _, line := range lines {
		if line == "" {
			g[groupNum] = groupVals
			groupVals = []string{}
			groupNum++
			continue
		}
		groupVals = append(groupVals, line)
	}
	g[groupNum] = groupVals

	return g, nil
}
