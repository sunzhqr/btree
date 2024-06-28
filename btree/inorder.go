package main

import (
	"fmt"
)

type BTreeNode struct {
	Value int
	Left  *BTreeNode
	Right *BTreeNode
}

type BTree struct {
	Root *BTreeNode
}

// NewBTree creates a new BTree.
func NewBTree() *BTree {
	return &BTree{}
}

// ReplaceOrInsert adds a new value to the tree or replaces an existing one.
func (b *BTree) ReplaceOrInsert(value int) {
	newNode := &BTreeNode{Value: value}
	if b.Root == nil {
		b.Root = newNode
		return
	}

	currentNode := b.Root
	for {
		if value < currentNode.Value {
			if currentNode.Left == nil {
				currentNode.Left = newNode
				return
			}
			currentNode = currentNode.Left
		} else if value > currentNode.Value {
			if currentNode.Right == nil {
				currentNode.Right = newNode
				return
			}
			currentNode = currentNode.Right
		} else {
			currentNode.Value = value
			return
		}
	}
}

// InOrderTraversal traverses the tree in in-order and applies fn to each node.
func (b *BTree) InOrderTraversal(fn func(n *BTreeNode)) {
	var recurse func(node *BTreeNode)
	recurse = func(node *BTreeNode) {
		if node == nil {
			return
		}
		recurse(node.Left)
		fn(node)
		recurse(node.Right)
	}
	recurse(b.Root)
}

func main() {
	tree := NewBTree()
	tree.ReplaceOrInsert(50)
	tree.ReplaceOrInsert(30)
	tree.ReplaceOrInsert(70)
	tree.ReplaceOrInsert(20)
	tree.ReplaceOrInsert(40)
	tree.ReplaceOrInsert(60)
	tree.ReplaceOrInsert(80)

	tree.InOrderTraversal(func(n *BTreeNode) {
		fmt.Print(n.Value, " ") // This will print: 20 30 40 50 60 70 80
	})
}
