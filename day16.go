package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

var seed int
var seedLen int

// Indices where 0 was appended in an expansion step
var pivots = []int{}

var dataLen int

func main() {
	flag.Parse()

	seedStr := flag.Arg(0)
	seedLen := len(seedStr)

	// Initial state and checksum are little-endian:
	// least significant bit first (the opposite of the usual convention).
	for i := 0; i < len(seedStr); i++ {
		if seedStr[i] == '1' {
			seed |= 1 << uint(i)
		}
	}

	var err error
	dataLen, err = strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	n := seedLen
	for n < dataLen {
		pivots = append(pivots, n)
		n = 2 * n + 1
	}

	//fmt.Printf("D %5d ", dataLen)
	//for i := 0; i < dataLen; i++ {
	//	fmt.Printf("%b", getDataBit(i))
	//}
	//fmt.Println()

	checksumLen := dataLen
	for checksumLen%2 == 0 {
		checksumLen >>= 1
	}

	//for cn := dataLen / 2; cn >= checksumLen; cn /= 2 {
	//  fmt.Printf("C %5d ", cn)
	//	for i := 0; i < cn; i++ {
	//		fmt.Printf("%b", getChecksumBit(i, cn))
	//	}
	//	fmt.Printf("\n")
	//}

	for i := 0; i < checksumLen; i++ {
		b := getChecksumBit(i, checksumLen)
		fmt.Printf("%b", b)
	}
	fmt.Printf("\n")
}

func getChecksumBit(i, checksumLen int) int {
	var a, b int

	if checksumLen * 2 == dataLen {
		a = getDataBit(2 * i)
		b = getDataBit(2 * i + 1)
	} else {
		a = getChecksumBit(2 * i, 2 * checksumLen)
		b = getChecksumBit(2 * i + 1, 2 * checksumLen)
	}

	if a == b {
		return 1
	} else {
		return 0
	}
}

func getDataBit(i int) int {
	b, s := getDataBit2(i)
	return b ^ s
}

func getDataBit2(i int) (bit, invert int) {
	whichPivot := len(pivots) - 1
	invert = 0
	for {
		p := pivots[whichPivot]
		if i == p {
			return 0, invert
		}

		if i > p {
			i = p - (i - p)
			invert ^= 1
		} else if whichPivot == 0 {
			return (seed >> uint(i)) & 1, invert
		} else {
			whichPivot--
		}
	}
}
