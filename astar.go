package astar

import "container/heap"
import "fmt"

func dprintln(obj ...interface{}) {
	if false {
		fmt.Println(obj...)
	}
}

type Node interface {
	IsGoal() bool
	DistanceTo() int
	EstimateFrom() int
	NextNodes() []Node
}

type nodeQueue []Node

func (q nodeQueue) Len() int {
	return len(q)
}

func (q nodeQueue) Less(i, j int) bool {
	u := q[i]
	v := q[j]
	return u.DistanceTo() + u.EstimateFrom() <
	       v.DistanceTo() + v.EstimateFrom()
}

func (q nodeQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *nodeQueue) Push(x interface{}) {
	y := x.(Node)
	*q = append(*q, y)
}

func (q *nodeQueue) Pop() interface{} {
	n := q.Len() - 1
	y := (*q)[n]
	*q = (*q)[:n]
	return y
}

type History interface {
	Visited(Node) bool
	SetVisited(Node)
}

type nullHist byte

func (nullHist) Visited(Node) bool { return false }
func (nullHist) SetVisited(Node) {}

// Performs A* search from an initial Node.
// Returns a goal Node.
func Search(v0 Node, hist History) Node {
	if hist == nil {
		hist = nullHist(0)
	}

	queue := nodeQueue([]Node{v0})

	for len(queue) > 0 {
		v := heap.Pop(&queue).(Node)
		dprintln("popped", v)
		if v.IsGoal() {
			return v
		}

		hist.SetVisited(v)

		for _, vn := range v.NextNodes() {
			if hist.Visited(vn) {
				continue
			}

			heap.Push(&queue, vn)
			dprintln("pushed", vn)
		}
	}

	return nil
}
