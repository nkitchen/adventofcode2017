package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

func dprint(obj... interface{}) {
	if false {
		fmt.Println(obj...)
	}
}

// A node of an indexable binary tree
type node struct {
	value       int
	left, right *node
	// Number of nodes in subtree
	size int
}

func (t *node) String() string {
	if t == nil {
		return "_"
	}

	return fmt.Sprintf("([%v] %v %v %v)", t.size, t.value, t.left, t.right)
}

func insert(t *node, x int) *node {
	if t == nil {
		return &node{x, nil, nil, 1}
	}

	if x == t.value {
		return t
	}

	if x < t.value {
		t.left = insert(t.left, x)
	} else {
		t.right = insert(t.right, x)
	}
	t.size++
	return t
}

func insertRange(t *node, min, max int) *node {
	m := (min + max) / 2
	t = insert(t, m)
	if min < m {
		t = insertRange(t, min, m-1)
	}
	if m < max {
		t = insertRange(t, m+1, max)
	}
	return t
}

func deleteAt(t *node, i int) *node {
	dprint("deleteAt", t, i)
	nl := 0
	if t.left != nil {
		nl = t.left.size
	}

    if i == nl {
		// It's this node.
		if t.right == nil {
			// Size is already correct
			return t.left
		} else {
			t.right = deleteFirst(t.right, func(value int) {
				t.value = value
			})
		}
	} else if i < nl {
		t.left = deleteAt(t.left, i)
	} else if i >= nl + 1 {
		t.right = deleteAt(t.right, i - (nl + 1))
	}
	t.size--
	return t
}

func deleteFirst(t *node, f func(int)) *node {
	if t.left == nil {
		f(t.value)
		return t.right
	} else {
		t.left = deleteFirst(t.left, f)
		t.size--
		return t
	}
}

func main() {
	flag.Parse()

	n, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	t := insertRange(nil, 0, n-1)

    for t.size > 1 {
		i := 0
		for i < t.size {
			if t.size == 1 {
				break
			}
			j := (i + t.size / 2) % t.size
			t = deleteAt(t, j)
			if j > i {
				i++
			}
			dprint(t)
		}
	}

	fmt.Println(t.value)
}
