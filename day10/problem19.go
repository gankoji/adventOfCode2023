package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type direction uint

const (
	NORTH direction = iota
	EAST
	SOUTH
	WEST
)

func (d direction) String() string {
	switch d {
	case NORTH:
		return "N"
	case SOUTH:
		return "S"
	case WEST:
		return "W"
	case EAST:
		return "E"
	default:
		panic("forgot to handle a direction type")
	}
}

func (d direction) reverse() direction {
	return (d + 2) % 4
}

type coord struct {
	x int
	y int
}

func (c coord) add(c2 coord) coord {
	return coord{x: c.x + c2.x, y: c.y + c2.y}
}

type grid map[coord]byte

type scenario struct {
	width   int
	height  int
	g       grid
	pos     coord
	heading direction
	steps   int
}

func (s *scenario) advance() {
	for _, v := range vectors {
		if v.dir == s.heading.reverse() {
			continue
		}
		if !slices.Contains(connections[s.g[s.pos]], v.dir) {
			continue
		}

		c := s.pos.add(v.coord)
		tile := s.g[c]
		if slices.Contains(connections[tile], v.dir.reverse()) {
			s.heading = v.dir
			s.pos = c
			s.steps++
			return
		}
	}
	panic("could not find a path to advance, something is broken")
}

func (s scenario) String() string {
	var sb strings.Builder
	for y := s.height - 1; y >= 0; y-- {
		for x := 0; x < s.width; x++ {
			tile := s.g[coord{x, y}]
			if x == s.pos.x && y == s.pos.y {
				switch s.heading {
				case NORTH:
					sb.WriteRune('⮙')
				case SOUTH:
					sb.WriteRune('⮛')
				case EAST:
					sb.WriteRune('⮚')
				case WEST:
					sb.WriteRune('⮘')
				}
			} else if tile > 0 {
				sb.WriteByte(tile)
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	sb.WriteString(fmt.Sprintf("Pos: %v\tHeading: %v\tSteps: %d\n", s.pos, s.heading, s.steps))
	return sb.String()
}

type vector struct {
	coord
	dir direction
}

var vectors = []vector{
	{coord: coord{1, 0}, dir: EAST},
	{coord: coord{-1, 0}, dir: WEST},
	{coord: coord{0, -1}, dir: SOUTH},
	{coord: coord{0, 1}, dir: NORTH},
}

var connections = map[byte][]direction{
	'|': {NORTH, SOUTH},
	'-': {EAST, WEST},
	'L': {NORTH, EAST},
	'J': {NORTH, WEST},
	'7': {SOUTH, WEST},
	'F': {SOUTH, EAST},
	'S': {NORTH, EAST, SOUTH, WEST},
}

func parseInput(lines [][]byte) scenario {
	grid, start := grid{}, coord{}
	for y, line := range lines {
		y = len(lines) - y - 1 // invert for normal north/south
		for x, b := range line {
			c := coord{x, y}
			grid[c] = b
			if b == 'S' {
				start = c
			}
		}
	}

	// pick a valid starting direction
	var chosenVec vector
	for _, v := range vectors {
		tile := grid[start.add(v.coord)]
		if slices.Contains(connections[tile], v.dir.reverse()) {
			chosenVec = v
			break
		}
	}

	return scenario{width: len(lines[0]), height: len(lines), g: grid, pos: start, heading: chosenVec.dir, steps: 0}
}

func part1(s scenario) {
	for {
		s.advance()
		if s.g[s.pos] == 'S' {
			break
		}
	}
	fmt.Printf("Part 1: %d\n", s.steps/2)
}

func part2(s scenario) {
	// trace out the loop
	visited := map[coord]struct{}{s.pos: {}}
	for {
		s.advance()
		visited[s.pos] = struct{}{}
		if s.g[s.pos] == 'S' {
			break
		}
	}

	// Replace 'S' with it's underlying pipe type so the following algorithm
	// can detect walls where the start was.
	var startDirs []direction
	for _, v := range vectors {
		tile := s.g[s.pos.add(v.coord)]
		if slices.Contains(connections[tile], v.dir.reverse()) {
			startDirs = append(startDirs, v.dir)
		}
	}
	if len(startDirs) != 2 {
		panic("uh-oh, somehow the start position doesn't have exactly 2 paths!")
	}
	for pt, c := range connections {
		if pt != 'S' && slices.Contains(c, startDirs[0]) && slices.Contains(c, startDirs[1]) {
			s.g[s.pos] = pt
			break
		}
	}

	insideCount := 0
	for x := 1; x < s.width-1; x++ {
		for y := 1; y < s.height-1; y++ {
			if _, ok := visited[coord{x, y}]; ok {
				// skip anything that overlaps the loop
				continue
			}

			// count how many lines it crosses
			crossings := 0
			for i := x + 1; i < s.width; i++ {
				c := coord{i, y}
				tile := s.g[c]
				if _, ok := visited[c]; ok && slices.Contains([]byte{'|', 'J', 'L'}, tile) {
					crossings++
				}
			}

			// odd crossings means it's inside the loop
			if crossings%2 == 1 {
				insideCount++
			}
		}
	}

	fmt.Printf("Part 2: %d\n", insideCount)
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	must(err, "reading input file")
	scen := parseInput(bytes.Split(data, []byte{'\n'}))
	part1(scen)
	part2(scen)
}

func must(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %e", msg, err)
	}
}
