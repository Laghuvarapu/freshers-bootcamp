package main

import (
	"fmt"
	"strings"
)

type Node struct {
	data  string
	left  *Node
	right *Node
}

func createTree(sli []string) *Node {
	if len(sli) == 0 {
		return nil
	}
	if !check(sli[0]) {
		node := Node{sli[0], nil, nil}
		head := createTree(sli[1:])
		if head == nil {
			return &node
		}
		head.left = &node
		return head
	} else {
		node := Node{sli[0], nil, nil}
		right := createTree(sli[1:])
		node.right = right
		return &node
	}
}
func check(s string) bool {
	if s == "+" || s == "-" {
		return true
	}
	return false
}
func preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.data + " ")
	preOrder(node.left)
	preOrder(node.right)
}
func postOrder(node *Node) {
	if node == nil {
		return
	}
	postOrder(node.left)
	postOrder(node.right)
	fmt.Print(node.data + " ")
}
func main() {
	s := "a - b + c"
	sli := strings.Split(s, " ")
	root := createTree(sli)
	preOrder(root)
	fmt.Println()
	postOrder(root)
	fmt.Println()
}
