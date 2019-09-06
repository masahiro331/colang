package main

import (
	"fmt"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	if t == nil {
		return ""
	}
	values := []string{}
	if t.left != nil {
		values = append(values, t.left.String())
	}
	values = append(values, fmt.Sprintf("%d", t.value))
	if t.right != nil {
		values = append(values, t.right.String())
	}
	return strings.Join(values, " ")
}

func main() {
	t := tree{}
	add(&t, 1)
	add(&t, 2)
	add(&t, 3)
	fmt.Println(t.String())
}
