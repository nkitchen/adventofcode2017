package main

import (
	"container/heap"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
)

const (
	W = 4
	H = 4
)

type state struct {
	x, y int
	path string
}

func (s *state) dist() int {
	return len(s.path) + (W - s.x) + (H - s.y)
}

type stateQueue []*state

func (q stateQueue) Len() int {
	return len(q)
}

func (q stateQueue) Less(i, j int) bool {
	s := q[i]
	t := q[j]
	return s.dist() < t.dist()
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

var step = "UDLR"
var dir = []struct{ dx, dy int }{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

func main() {
	flag.Parse()

	passcode := flag.Arg(0)

	s0 := state{0, 0, ""}

	queue := stateQueue([]*state{&s0})

    maxLen := -1
	for len(queue) > 0 {
		s := heap.Pop(&queue).(*state)

		if s.x == W-1 && s.y == H-1 {
			maxLen = len(s.path)
			continue
		}

		h := md5.New()
		io.WriteString(h, passcode)
		io.WriteString(h, s.path)
		sum := fmt.Sprintf("%x", h.Sum(nil))

		for i := 0; i < len(step); i++ {
			switch sum[i] {
			case 'b':
			case 'c':
			case 'd':
			case 'e':
			case 'f':
				// Door open
				break
			default:
				// Door closed
				continue
			}

			xn := s.x + dir[i].dx
			yn := s.y + dir[i].dy
			if xn < 0 || xn >= W {
				continue
			}
			if yn < 0 || yn >= H {
				continue
			}

			sn := state{xn, yn, s.path + string(step[i])}
			heap.Push(&queue, &sn)
		}
	}

	fmt.Println(maxLen)
}
