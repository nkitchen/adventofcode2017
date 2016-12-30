package main

import (
	"astar"
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type pos struct {
	x, y int
}

type state struct {
	pos
	// Bit string of visited locations
	visited int
}

type node struct {
	state
	distTo int
	estFrom int
}

func (v node) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "node{(%d, %d) [", v.x, v.y)
	for loc := 0; loc < numLocs; loc++ {
		if v.Visited(loc) {
			fmt.Fprintf(buf, "%d", loc)
		}
	}
	fmt.Fprintf(buf, "] d=%d e=%d}", v.distTo, v.estFrom)
	return buf.String()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func newNode(p pos, visited int, distTo int) node {
	v := node{state{p, visited}, distTo, 0}

	if v.IsGoal() {
		return v
	}

	// Estimated distance remaining to travel:
	// Find bounding box of unvisited locations.
	xMin := len(ductMap[0])
	xMax := -1
	yMin := len(ductMap)
	yMax := -1
	toVisit := 0

	for loc := 1; loc < numLocs; loc++ {
		if v.Visited(loc) {
			continue
		}
		toVisit++
		p := locPos[loc]
		xMin = min(xMin, p.x)
		xMax = max(xMax, p.x)
		yMin = min(yMin, p.y)
		yMax = max(yMax, p.y)
	}

	e := 0
	p0 := locPos[0]
	if toVisit > 0 {
		// One edge length in each dimension,
		// plus the distance to the nearest edge
		e += xMax - xMin + yMax - yMin
		dxMin := abs(v.x - xMin)
		dxMax := abs(xMax - v.x)
		e += min(dxMin, dxMax)
		dyMin := abs(v.y - yMin)
		dyMax := abs(yMax - v.y)
		e += min(dyMin, dyMax)

		// ...plus distance from the bounding box to location 0,
		// if it's outside the box.
		if p0.x < xMin || p0.x > xMax {
			e += min(xMin - p0.x, p0.x - xMax)
		}
		if p0.y < yMin || p0.y > yMax {
			e += min(yMin - p0.y, p0.y - yMax)
		}
	} else {
		e += abs(v.x - p0.x) + abs(v.y - p0.y)
	}

	v.estFrom = e

	return v
}

func (v node) Visited(loc int) bool {
	return (v.visited >> uint(loc)) & 1 == 1
}

func (v node) IsGoal() bool {
	return v.visited == (1 << uint(numLocs)) - 1 &&
		v.pos == locPos[0]
}

func (v node) DistanceTo() int {
	return v.distTo
}

func (v node) EstimateFrom() int {
	return v.estFrom
}

func (v node) NextNodes() []astar.Node {
	r := []astar.Node{}
	for _, p := range ([]pos{
		{v.x - 1, v.y},
		{v.x + 1, v.y},
		{v.x, v.y - 1},
		{v.x, v.y + 1},
	}) {
		c := ductMap[p.y][p.x:p.x+1]
		loc, err := strconv.Atoi(c)
		if err == nil {
			// Digit
			vn := newNode(p, v.visited | (1 << uint(loc)),
			              1 + v.distTo)
			r = append(r, vn)
		} else if c[0] == '.' {
			vn := newNode(p, v.visited, 1 + v.distTo)
			r = append(r, vn)
		}
	}
	return r
}

type history map[state]bool

func (h history) Visited(v astar.Node) bool {
	s := v.(node).state
	return h[s]
}

func (h history) SetVisited(v astar.Node) {
	s := v.(node).state
	h[s] = true
}

var ductMap = []string{}

// Coordinates of numbered locations
var locPos = map[int]pos{}

var numLocs = 0

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		ductMap = append(ductMap, line)
		y := len(ductMap) - 1
		for x := range line {
			loc, err := strconv.Atoi(line[x:x+1])
			if err == nil {
				locPos[loc] = pos{x, y}
				if loc >= numLocs {
					numLocs = loc + 1
				}
			}
		}
	}

	v0 := newNode(locPos[0], 1, 0)

	hist := history(map[state]bool{})
	vf := astar.Search(v0, hist)
	fmt.Println(vf.DistanceTo())
}
