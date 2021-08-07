package main

import (
	"errors"
	"fmt"
)

type RBT struct {
	root *Node
}

type Node struct {
	key         int
	val         string
	left, right *Node
	color       bool //red=true black=false
}

func main() {
	rbt := &RBT{nil}
	rbt.put(6, "string6")
	rbt.put(4, "string4")
	rbt.put(2, "string2")
	rbt.put(5, "string5")
	rbt.put(9, "string9")
	rbt.put(8, "string8")
	rbt.put(11, "string11")
	rbt.put(12, "string12")
	rbt.root.right.right.printLR()
	// rbt.get(2)
}

func (n *Node) isRed() bool {
	if n == nil {
		return false
	}
	return n.color
}

func (n *Node) rotateLeft() (*Node, error) {
	if !n.right.isRed() {
		return nil, errors.New("link not red")
	}
	right := n.right
	n.right = right.left
	right.left = n
	right.color = n.color
	n.color = true
	return right, nil
}

func (n *Node) rotateRight() (*Node, error) {
	if !n.left.isRed() {
		return nil, errors.New("link not red")
	}
	left := n.left
	n.left = left.right
	left.right = n
	left.color = n.color
	n.color = true
	return left, nil
}

func (n *Node) colorFlip() error {
	if !n.left.isRed() || !n.right.isRed() {
		return errors.New("link not red")
	}
	n.color = true
	n.left.color = false
	n.right.color = false
	return nil
}

func (rbt *RBT) put(key int, val string) {
	node := &Node{key, val, nil, nil, true}
	if rbt.root == nil {
		rbt.root = node
		return
	}
	resNode := rbt.root.put(key, val)
	rbt.root = resNode
}

func (n *Node) put(key int, val string) *Node {
	if n == nil {
		return &Node{key, val, nil, nil, true}
	}
	if key < n.key {
		n.left = n.left.put(key, val)
	} else {
		n.right = n.right.put(key, val)
	}
	if n.right.isRed() && !n.left.isRed() {
		n, _ = n.rotateLeft()
	}

	if n.left.isRed() && n.left.left.isRed() {
		n, _ = n.rotateRight()
	}
	if n.right != nil {
		if n.left.isRed() && n.right.isRed() {
			n.colorFlip()
		}
	}
	return n
}

func (rbt *RBT) get(key int) (string, error) {
	node, err := rbt.search(key)
	if err != nil {
		return "", err
	}
	return node.val, nil
}

func (rbt RBT) search(key int) (*Node, error) {
	for {
		if rbt.root == nil {
			return nil, errors.New("not found")
		}
		if rbt.root.key == key {
			return rbt.root, nil
		}
		if rbt.root.key < key {
			rbt.root = rbt.root.right
			continue
		}
		if rbt.root.key > key {
			rbt.root = rbt.root.left
			continue
		}
	}
}

func (n *Node) printLR() {
	fmt.Println("root: ", n)
	fmt.Println("left: ", n.left)
	fmt.Println("right: ", n.right)
}
