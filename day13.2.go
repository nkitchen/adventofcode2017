package main

import (
	"container/heap"
	"flag"
	"fmt"
	"log"
	"strconv"
)

func bitCount(x int) int {
	n := 0
	for x != 0 {
		n++
		x &= x - 1
	}
	return n
}

type point struct {
	x, y int
}

type state struct {
	point
	stepsSoFar int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type stateQueue []*state

func (q stateQueue) Len() int {
	return len(q)
}

func (q stateQueue) Less(i, j int) bool {
	s := q[i]
	t := q[j]
	sd := s.stepsSoFar
	td := t.stepsSoFar
	return sd < td
}

func (q stateQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *stateQueue) Push(x interface{}) {
	y := x.(*state)
	*q = append(*q, y)
}

func (q *stateQueue) Pop() interface{} {
	n := q.Len() - 1
	y := (*q)[n]
	*q = (*q)[:n]
	return y
}

var seed int

func isWall(x, y int) bool {
	s := x*x + 3*x + 2*x*y + y + y*y + seed
	return bitCount(s)%2 == 1
}

func main() {
	flag.Parse()

	var err error
	seed, err = strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	s0 := state{point{1, 1}, 0}

	qSlice := []*state{&s0}
	queue := stateQueue(qSlice)
	visited := map[point]bool{}

	for len(queue) > 0 {
		s := heap.Pop(&queue).(*state)
		if s.stepsSoFar > 50 {
			break
		}

		visited[point{s.x, s.y}] = true

		for xn := s.x - 1; xn <= s.x+1; xn++ {
			for yn := s.y - 1; yn <= s.y+1; yn++ {
				if xn < 0 || yn < 0 {
					continue
				}

				if abs(xn-s.x)+abs(yn-s.y) != 1 {
					continue
				}

				if visited[point{xn, yn}] {
					continue
				}

				if isWall(xn, yn) {
					continue
				}

				sn := state{point{xn, yn}, 1 + s.stepsSoFar}
				heap.Push(&queue, &sn)
			}
		}
	}

	fmt.Println(len(visited))
}
