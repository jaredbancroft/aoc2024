package cmd

import (
	"fmt"
	"slices"
	"strings"

	"github.com/jaredbancroft/aoc2024/internal/helpers"
	"github.com/spf13/cobra"
)

// day6Cmd represents the day1 command
var day6Cmd = &cobra.Command{
	Use:   "day6",
	Short: "Advent of Code 2024 - Day 6: Guard Gallivant",
	Long: `
--- Day 6: Guard Gallivant ---

	`,
	RunE: func(cmd *cobra.Command, args []string) error {

		part1 := 0
		part2 := 0

		inputs, err := helpers.ReadStringFile("inputs/day6.txt")
		if err != nil {
			return err
		}

		labMap := make([][]string, len(inputs))
		visitedSpaces := make(map[coord]int)

		var startX int
		var startY int
		var startD direction

		for i, line := range inputs {
			values := strings.Split(line, "")
			if slices.Contains(values, "^") {
				startX = slices.Index(values, "^")
				startY = i
				startD = up
			}
			labMap[i] = append(labMap[i], values...)
		}

		path := make(map[coord]direction)

		g := guard{
			coord{startX, startY},
			startD,
			&labMap,
			&visitedSpaces,
			&path,
		}

		cond := clear

		for cond != oob {
			cond = g.patrol()
		}
		part1 = len(visitedSpaces)

		for y, col := range labMap {
			for x := range col {
				if labMap[y][x] == "." || labMap[y][x] == "^" {
					g.setPosition(coord{startX, startY})
					g.setDirection(startD)
					path = make(map[coord]direction)
					visitedSpaces = make(map[coord]int)
					labMap[y][x] = "#"

					cond := clear
					for cond != oob {
						cond = g.patrol()
						if cond == loop {
							part2++
							break
						}
					}
					labMap[y][x] = "."
				}
			}
		}

		// loc.p.x = startx
		// loc.p.y = starty
		// loc.d = startd
		// loc.m = labMap
		// loc.c = true
		// loc.v = make(map[pair]int)
		// loc.q = make([]pair, 0)
		// loc.ld = false

		// for loc.c {
		// 	loc.move()
		// }

		// part1 := len(loc.v)

		// loc2 := guard{}
		// loc2.p.x = startx
		// loc2.p.y = starty
		// loc2.d = startd
		// loc2.m = labMap
		// loc2.c = true
		// loc2.v = make(map[pair]int)
		// loc2.q = make([]pair, 0)
		// loc2.ld = true
		// loc2.lc = 0

		// for j, values := range loc2.m {
		// 	for i := range values {

		// 		if i == startx && j == starty {
		// 			continue
		// 		}
		// 		if loc2.m[j][i] == "#" {
		// 			continue
		// 		}
		// 		loc2.m[j][i] = "#"
		// 		for loc2.c {
		// 			loc2.move()
		// 		}
		// 		loc2.m[j][i] = "."
		// 		loc2.c = true
		// 		loc2.q = nil
		// 		loc2.p.x = startx
		// 		loc2.p.y = starty
		// 	}
		// }

		// part2 := loc.lc

		fmt.Println("Part 1 Total:", part1)
		fmt.Println("Part 2 Total:", part2)
		return nil
	},
}

func printMap(m *[][]string) {
	for _, row := range *m {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}
}

type condition int

const (
	oob condition = iota
	clear
	obstructed
	loop
)

var conditionName = map[condition]string{
	oob:        "Out of Bounds",
	clear:      "Clear",
	obstructed: "Obstructed",
	loop:       "Loop",
}

func (c condition) String() string {
	return conditionName[c]
}

type direction int

const (
	up direction = iota
	right
	down
	left
)

var directionName = map[direction]string{
	up:    "up",
	right: "right",
	down:  "down",
	left:  "left",
}

func (d direction) String() string {
	return directionName[d]
}

func changeDirection(d direction) direction {
	switch d {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	default:
		panic(fmt.Errorf("unknown direction: %s", d))
	}
}

type coord struct {
	x, y int
}

type guard struct {
	pos     coord
	dir     direction
	_map    *[][]string
	visited *map[coord]int
	path    *map[coord]direction
}

func (g *guard) setPosition(c coord) {
	g.pos = c
}

func (g *guard) setDirection(d direction) {
	g.dir = d
}

func (g *guard) setMap(m [][]string) {
	g._map = &m
}

func (g *guard) setVisited(v map[coord]int) {
	g.visited = &v
}

func (g *guard) isObstructed(s string) bool {
	return s == "#"
}

func (g *guard) isOOB(c coord) bool {
	if c.x >= 0 && c.x < len((*g._map)[0]) && c.y >= 0 && c.y < len(*g._map) {
		return false
	}
	return true
}

func (g *guard) isClear(s string) bool {
	return s == "." || s == "^"
}

func (g *guard) patrol() condition {
	nextSpace := g.getNextSpace()

	if g.isOOB(nextSpace) {
		return oob
	}

	if g.isObstructed((*g._map)[nextSpace.y][nextSpace.x]) {
		g.setDirection(changeDirection(g.dir))
		fmt.Println(g.pos, g.dir)
		if (*g.path)[g.pos] == g.dir {
			fmt.Println("test")
			return loop
		}
		return obstructed
	}

	if g.isClear((*g._map)[nextSpace.y][nextSpace.x]) {
		(*g.path)[g.pos] = g.dir
		g.setPosition(nextSpace)
		(*g.visited)[nextSpace]++
		return clear
	}
	return clear
}

func (g *guard) getNextSpace() coord {
	switch g.dir {
	case up:
		return coord{g.pos.x, g.pos.y - 1}
	case right:
		return coord{g.pos.x + 1, g.pos.y}
	case down:
		return coord{g.pos.x, g.pos.y + 1}
	case left:
		return coord{g.pos.x - 1, g.pos.y}
	default:
		panic(fmt.Errorf("unknown direction: %s", g.dir))
	}
}

// type guard struct {
// 	p  pair
// 	d  direction
// 	m  [][]string
// 	c  bool
// 	v  map[pair]int
// 	q  []pair
// 	lc int
// 	ld bool
// }

// func (g *guard) move() {
// 	switch g.d {
// 	case up:
// 		g.p.y = g.p.y - 1
// 		if g.checkBounds() {
// 			if g.checkSpace() {
// 				g.p.y = g.p.y + 1
// 				g.detectLoop(g.p)
// 			} else {
// 				g.v[g.p] += 1
// 			}
// 		} else {
// 			g.c = false
// 		}
// 	case right:
// 		g.p.x = g.p.x + 1
// 		if g.checkBounds() {
// 			if g.checkSpace() {
// 				g.p.x = g.p.x - 1
// 				g.detectLoop(g.p)
// 			} else {
// 				g.v[g.p] += 1
// 			}
// 		} else {
// 			g.c = false
// 		}
// 	case down:
// 		g.p.y = g.p.y + 1
// 		if g.checkBounds() {
// 			if g.checkSpace() {
// 				g.p.y = g.p.y - 1
// 				g.detectLoop(g.p)
// 			} else {
// 				g.v[g.p] += 1
// 			}
// 		} else {
// 			g.c = false
// 		}
// 	case left:
// 		g.p.x = g.p.x - 1
// 		if g.checkBounds() {
// 			if g.checkSpace() {
// 				g.p.x = g.p.x + 1
// 				g.detectLoop(g.p)
// 			} else {
// 				g.v[g.p] += 1
// 			}
// 		} else {
// 			g.c = false
// 		}
// 	}
// }

// func (g *guard) checkBounds() bool {
// 	if g.p.x >= 0 && g.p.x < len(g.m[0]) && g.p.y >= 0 && g.p.y < len(g.m) {
// 		return true
// 	}

// 	return false
// }

// func (g *guard) checkSpace() bool {
// 	if g.m[g.p.y][g.p.x] == "#" {
// 		g.d = changeDirection(g.d)
// 		return true
// 	}

// 	return false
// }

// func (g *guard) detectLoop(n pair) {
// 	if g.ld {
// 		if slices.Contains(g.q, n) {
// 			fmt.Println("loop")
// 			g.lc += 1
// 			g.c = false
// 		}
// 		g.q = append(g.q, n)
// 	}
// }

func init() {
	rootCmd.AddCommand(day6Cmd)
}
