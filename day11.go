package main

import (
	"bufio"
	"container/heap"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"regexp"
)

type itemSet uint64

func (s itemSet) forEachItem(f func(item itemSet)) {
	for s != 0 {
		withoutFirstItem := s & (s - 1)
		i := s &^ withoutFirstItem
		f(i)
		s = withoutFirstItem
	}
}

func setOfChip(i uint) itemSet {
	return itemSet(uint64(1) << i)
}

func setOfGen(i uint) itemSet {
	return itemSet(uint64(1) << (i + 32))
}

func (s itemSet) Len() int {
	n := 0
	s.forEachItem(func(itemSet) {
		n++
	})
	return n
}
	
func (s itemSet) safe() bool {
	chips := uint32(s)
	gens := uint32(s >> 32)
	fried := ((chips &^ gens) != 0 && gens != 0)
	return !fried
}

func (s itemSet) String() string {
	bits := uint64(s)
	chips := uint32(bits)
	gens := uint32(bits >> 32)
	return fmt.Sprintf("%b/%b", gens, chips)
}

func encoded(floorContents [4]itemSet, elevatorFloor int) string {
	a := [128]byte{}
	buf := a[:]
	n := 0
	for _, floor := range floorContents {
		n += binary.PutUvarint(buf[n:], uint64(floor))
	}
	n += binary.PutUvarint(buf[n:], uint64(elevatorFloor))
	return string(buf[:n])
}

type state struct {
	floorContents [4]itemSet
	elevatorFloor int
	stepsSoFar    int
	stepsToGo     int
}

func (s *state) estimatedStepsRemaining() int {
	if s.stepsToGo >= 0 {
		return s.stepsToGo
	}

	// Heuristic: a direct ride to the assembler for each pair of items,
	// with no constraints
	n := 0
	top := len(s.floorContents) - 1
	extra := 0
	for i := 0; i < top; i++ {
		floor := s.floorContents[i]
		d := top - i
		m := floor.Len() + extra
		extra = m % 2
		n += d * (m/2) + extra
	}
	s.stepsToGo = n
	return n
}

func (s state) String() string {
	d := s.stepsSoFar
	e := s.estimatedStepsRemaining()
	f := d + e
	return fmt.Sprintf("{%v %v %v %v =%v}", s.floorContents, s.elevatorFloor,
	                   d, e, f)
}

func (s *state) finished() bool {
	for i := 0 ; i < len(s.floorContents) - 1; i++ {
		if s.floorContents[i] != 0 {
			return false
		}
	}
	return true
}

type stateQueue []*state

func (q stateQueue) Len() int {
	return len(q)
}

func (q stateQueue) Less(i, j int) bool {
	s := q[i]
	t := q[j]
	sd := s.stepsSoFar + s.estimatedStepsRemaining()
	td := t.stepsSoFar + t.estimatedStepsRemaining()
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

var elementId = map[string]uint{}

func main() {
	floorIndex := map[string]int{
		"first":  0,
		"second": 1,
		"third":  2,
		"fourth": 3,
	}

	chipRe := regexp.MustCompile(`an? (\S+)-compatible microchip`)
	genRe := regexp.MustCompile(`an? (\S+) generator`)

	floorContents := [4]itemSet{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		var ordinal string
		n, _ := fmt.Sscanf(line, "The %s floor contains", &ordinal)
		if n != 1 {
			continue
		}

		f, ok := floorIndex[ordinal]
		if !ok {
			log.Fatal("Unexpected floor: %s", ordinal)
		}

		for kind, re := range map[string]*regexp.Regexp{
			"chip": chipRe,
			"gen": genRe,
		} {
			for _, m := range re.FindAllStringSubmatch(line, -1) {
				elemName := m[1]
				e, ok := elementId[elemName]
				if !ok {
					e = uint(len(elementId))
					if e >= 32 {
						log.Fatal("Too many elements")
					}
					elementId[elemName] = e
				}
				switch kind {
				case "chip":
					floorContents[f] |= setOfChip(e)
				case "gen":
					floorContents[f] |= setOfGen(e)
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	chips := itemSet(0)
	gens := itemSet(0)
	mask := itemSet(uint64(^uint32(0)))
	for _, floor := range floorContents {
		chips ^= floor & mask
		gens ^= floor >> 32
	}
	if chips != gens {
		log.Fatal("Mismatch between microchips and generators")
	}

	s0 := state{floorContents, 0, 0, -1}

	qSlice := []*state{&s0}
	queue := stateQueue(qSlice)
	visited := map[string]*state{}

	for len(queue) > 0 {
		s := heap.Pop(&queue).(*state)
		enc := encoded(s.floorContents, s.elevatorFloor)
		visited[enc] = s
		//fmt.Println(s)

		if s.finished() {
			fmt.Println(s.stepsSoFar)
			break
		}

		nextFloors := []int{}
		if s.elevatorFloor > 0 {
			nextFloors = append(nextFloors, s.elevatorFloor-1)
		}
		if s.elevatorFloor < len(s.floorContents) - 1 {
			nextFloors = append(nextFloors, s.elevatorFloor+1)
		}

		curItems := s.floorContents[s.elevatorFloor]

		combos := []itemSet{}
		curItems.forEachItem(func(i itemSet) {
			combos = append(combos, i)
			curItems.forEachItem(func(j itemSet) {
				if j > i {
					combos = append(combos, i | j)
				}
			})
		})

		for _, carried := range combos {
			itemsLeft := curItems &^ carried
			if !itemsLeft.safe() {
				continue
			}

			for _, f := range nextFloors {
				nextItems := s.floorContents[f] | carried
				if !nextItems.safe() {
					continue
				}

				floorContents := s.floorContents
				floorContents[s.elevatorFloor] = itemsLeft
				floorContents[f] = nextItems

				enc := encoded(floorContents, f)
				if visited[enc] != nil {
					continue
				}

				sn := state{
					floorContents: floorContents,
					elevatorFloor: f,
					stepsSoFar: 1 + s.stepsSoFar,
					stepsToGo: -1,
				}
				heap.Push(&queue, &sn)
				//fmt.Println("pushed", sn)
			}
		}
	}
}
