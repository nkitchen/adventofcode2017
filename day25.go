package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const checkLen = 20
var expected [checkLen]int

const (
	cpy = iota
	inc
	dec
	jnz
	out
)

type Arg interface {
	Value() int
	Name() string
}

type LiteralArg int

func (a LiteralArg) Value() int {
	return int(a)
}

func (a LiteralArg) Name() string {
	panic(fmt.Sprintf("Misused literal: %v", int(a)))
}

var regs map[string]int

type RegisterArg string

func (a RegisterArg) Value() int {
	return regs[string(a)]
}

func (a RegisterArg) Name() string {
	return string(a)
}

func newArg(s string) Arg {
	x, err := strconv.Atoi(s)
	if err == nil {
		return LiteralArg(x)
	}
	return RegisterArg(s)
}

type Cmd struct {
	op  int
	x, y Arg
}

func main() {
	for i := 0; i < checkLen; i++ {
		expected[i] = i % 2
	}

	cmds := []Cmd{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var x, y string

		n, _ := fmt.Sscanf(line, "cpy %s %s", &x, &y)
		if n == 2 {
			cmds = append(cmds, Cmd{cpy, newArg(x), newArg(y)})
			continue
		}

		n, _ = fmt.Sscanf(line, "inc %s", &x)
		if n == 1 {
			cmds = append(cmds, Cmd{inc, RegisterArg(x), LiteralArg(0)})
			continue
		}

		n, _ = fmt.Sscanf(line, "dec %s", &x)
		if n == 1 {
			cmds = append(cmds, Cmd{dec, RegisterArg(x), LiteralArg(0)})
			continue
		}

		n, _ = fmt.Sscanf(line, "jnz %s %s", &x, &y)
		if n == 2 {
			cmds = append(cmds, Cmd{jnz, newArg(x), newArg(y)})
			continue
		}

		n, _ = fmt.Sscanf(line, "out %s", &x)
		if n == 1 {
			cmds = append(cmds, Cmd{out, newArg(x), LiteralArg(0)})
			continue
		}

		log.Fatal("Unknown command: %s", line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for a0 := 1; a0 < 1e6; a0++ {
		output := run(cmds, a0)
		//fmt.Println(a0, output)
		if output == expected {
			fmt.Println(a0)
			break
		}
	}
}

func run(cmds []Cmd, a0 int) [checkLen]int {
	regs = make(map[string]int)
	regs["a"] = a0

	output := make([]int, 0, checkLen)
	ip := 0
	for 0 <= ip && ip <= len(cmds) {
		cmd := cmds[ip]

		jumped := false
		switch cmd.op {
		case out:
			output = append(output, cmd.x.Value())
			if len(output) == checkLen {
				a := [checkLen]int{}
				copy(a[:checkLen], output)
				return a
			}
		case cpy:
			regs[cmd.y.Name()] = cmd.x.Value()
		case jnz:
			if cmd.x.Value() != 0 {
				ip += cmd.y.Value()
				jumped = true
			}
		case inc:
			regs[cmd.x.Name()]++
		case dec:
			regs[cmd.x.Name()]--
		default:
			log.Fatal("Unknown cmd: %v", cmd)
		}

		if !jumped {
			ip++
		}
	}

	rv := [checkLen]int{}
	return rv
}
