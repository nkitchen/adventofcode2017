package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type dest struct {
	kind  string
	index int
}

func main() {
	botInput := map[int][]int{}
	botDests := map[int][2]dest{}

	queue := []int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		var v int
		var b int
		n, _ := fmt.Sscanf(line, "value %d goes to bot %d", &v, &b)
		if n == 2 {
			botInput[b] = append(botInput[b], v)
			queue = append(queue, b)
			continue
		}

		var loDest dest
		var hiDest dest
		n, _ = fmt.Sscanf(line, "bot %d gives low to %s %d and high to %s %d",
			&b, &loDest.kind, &loDest.index,
			&hiDest.kind, &hiDest.index)
		if n == 5 {
			botDests[b] = [2]dest{loDest, hiDest}
			continue
		}

		log.Fatal("Unexpected input format: %s", line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	outputs := map[int][]int{}
	compared := map[int][][2]int{}

	for len(queue) > 0 {
		b := queue[0]
		queue = queue[1:]

		if len(botInput[b]) < 2 {
			continue
		}

		a := botInput[b][:2]
		botInput[b] = botInput[b][2:]

		if a[0] > a[1] {
			a[0], a[1] = a[1], a[0]
		}
		v := [2]int{a[0], a[1]}
		compared[b] = append(compared[b], v)

		dests, ok := botDests[b]
		if !ok {
			log.Fatal("Unknown bot: %d", b)
		}
		for i := 0; i < 2; i++ {
			d := dests[i]
			if d.kind == "output" {
				outputs[d.index] = append(outputs[d.index], a[i])
			} else {
				if d.kind != "bot" {
					log.Fatal("Unknown destination type: %s", d.kind)
				}

				botInput[d.index] = append(botInput[d.index], a[i])
				queue = append(queue, d.index)
			}
		}
	}

	for b, a := range compared {
		for _, p := range a {
			fmt.Printf("Bot %d compared: %d %d\n", b, p[0], p[1])
		}
	}
}
