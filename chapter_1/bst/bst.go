package main

import (
	"fmt"
)

//Node describes a node in the binary search tree
type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Value  int
}

func main() {
	nums := [...]int{1, 2, 3, 4, 5, 6}

	var root *Node
	for _, num := range nums {
		if root == nil {
			root = &Node{
				Value: num,
			}
		} else {
			insert(num, root)
		}
	}

	print(root)
}

func print(curr *Node) {
	if curr == nil {
		return
	}

	if curr.Left != nil {
		print(curr.Left)
	}

	fmt.Printf("%v ", curr.Value)

	if curr.Right != nil {
		print(curr.Right)
	}
}

func insert(value int, curr *Node) {
	if curr.Value == value {
		return
	} else if curr.Value < value {
		if curr.Left == nil {
			curr.Left = &Node{
				Value:  value,
				Parent: curr,
			}
		} else {
			insert(value, curr.Left)
		}
	} else {
		if curr.Right == nil {
			curr.Right = &Node{
				Value:  value,
				Parent: curr,
			}
		} else {
			insert(value, curr.Right)
		}
	}
}
