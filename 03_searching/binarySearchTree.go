package main

import (
	"errors"
	"fmt"
)

type BST struct {
	root *Node
}

type Node struct {
	key         int
	val         string
	left, right *Node
}

func main() {
	bst := &BST{nil}
	bst.put(6, "string6")
	bst.put(2, "string2")
	bst.put(7, "string7")
	bst.put(4, "string4")
	bst.put(1, "string1")
	bst.root.printLR()
	bst.get(2)
}

func (bst *BST) put(key int, val string) {
	node := &Node{key, val, nil, nil}
	if bst.root == nil {
		bst.root = node
		return
	}
	bst.root.put(key, val)
}

func (n *Node) put(key int, val string) *Node {
	if n == nil {
		return &Node{key, val, nil, nil}
	}
	if key < n.key {
		n.left = n.left.put(key, val)
	} else if key >= n.key {
		n.right = n.right.put(key, val)
	}
	return n
}

func (bst *BST) delete(key int) error {
	if bst.root == nil {
		return errors.New("null tree")
	}
	bst.root = bst.root.delete(key)
	return nil
}

func (n *Node) delete(key int) *Node {
	if key < n.key {
		n.left = n.left.delete(key)
	} else if key > n.key {
		n.right = n.right.delete(key)
	} else {
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		buffer := n
		n := min(buffer.right)
		n.right = deleteMin(buffer.right)
		n.left = buffer.right
	}
	return n
}

func deleteMin(n *Node) *Node {
	if n.left == nil {
		return n.right
	}
	n.left = deleteMin(n.left)
	return n
}
func min(n *Node) *Node {
	if n.left != nil {
		n = min(n.left)
	}
	return n
}

func (bst *BST) get(key int) (string, error) {
	node, err := bst.search(key)
	if err != nil {
		return "", err
	}
	return node.val, nil
}

func (bst BST) search(key int) (*Node, error) {
	for {
		if bst.root == nil {
			return nil, errors.New("not found")
		}
		if bst.root.key == key {
			return bst.root, nil
		}
		if bst.root.key < key {
			bst.root = bst.root.right
			continue
		}
		if bst.root.key > key {
			bst.root = bst.root.left
			continue
		}
	}
}

func (n *Node) printLR() {
	fmt.Println(n)
	fmt.Println(n.left)
	fmt.Println(n.right)
}
