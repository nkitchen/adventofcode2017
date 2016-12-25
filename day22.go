package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const tooBigThreshold = 200

func dprint(obj ...interface{}) {
	if false {
		fmt.Println(obj...)
	}
}

type pos struct {
	x, y int
}

type state struct {
	empty    pos
	goalData pos
}

var tooBig = map[pos]bool{}

type node struct {
	state
	moves int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (v node) goalDist() int {
	return abs(v.goalData.x) + abs(v.goalData.y)
}

func (n node) dist() int {
	emptyToGoal := abs(n.empty.x-n.goalData.x) +
		abs(n.empty.y-n.goalData.y) - 1
	goalToZero := 5 * n.goalData.x
	return emptyToGoal + goalToZero
}

type nodeQueue []node

func (q nodeQueue) Len() int {
	return len(q)
}

func (q nodeQueue) Less(i, j int) bool {
	u := q[i]
	v := q[j]
	return u.dist() < v.dist()
}

func (q nodeQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *nodeQueue) Push(x interface{}) {
	y := x.(node)
	*q = append(*q, y)
}

func (q *nodeQueue) Pop() interface{} {
	n := q.Len() - 1
	y := (*q)[n]
	*q = (*q)[:n]
	return y
}

func main() {
	var v0 node

	xMax := 0
	yMax := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var x, y, size, used int
		n, err := fmt.Sscanf(line, "/dev/grid/node-x%d-y%d %dT %dT",
			&x, &y, &size, &used)
		if n < 4 || err != nil {
			continue
		}

		if used == 0 {
			v0.empty = pos{x, y}
		}
		if used > tooBigThreshold {
			tooBig[pos{x, y}] = true
		}
		if x > xMax {
			xMax = x
		}
		if y > yMax {
			yMax = y
		}
	}

	v0.goalData.x = xMax

	queue := nodeQueue([]node{v0})
	visited := map[state]bool{}

	minDist := v0.goalDist()
	for len(queue) > 0 {
		v := heap.Pop(&queue).(node)
		dprint("popped", v)

		if v.goalData.x == 0 && v.goalData.y == 0 {
			fmt.Println(v.moves)
			break
		}

		visited[v.state] = true

		if v.goalDist() < minDist {
			minDist = v.goalDist()
			dprint(minDist)
		}

		for xn := v.empty.x - 1; xn <= v.empty.x+1; xn++ {
			for yn := v.empty.y - 1; yn <= v.empty.y+1; yn++ {
				if xn < 0 || yn < 0 {
					continue
				}
				if xn > xMax || yn > yMax {
					continue
				}

				// No diagonals
				if abs(xn-v.empty.x)+abs(yn-v.empty.y) != 1 {
					continue
				}

				if tooBig[pos{xn, yn}] {
					continue
				}

				g := v.goalData
				if xn == v.goalData.x && yn == v.goalData.y {
					g = v.empty
				}
				vn := node{
					state: state{
						empty:    pos{xn, yn},
						goalData: g},
					moves: 1 + v.moves}

				if visited[vn.state] {
					continue
				}
				heap.Push(&queue, vn)
				dprint("pushed", vn)
			}
		}
	}
}
