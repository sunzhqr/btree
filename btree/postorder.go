package main

import (
	"fmt"
)

// BTreeNode represents a node in the binary tree
type BTreeNode struct {
	Value int
	Left  *BTreeNode
	Right *BTreeNode
}

// BTree represents the structure of the binary tree
type BTree struct {
	Root *BTreeNode
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

// PostOrderTraversal initiates a post-order traversal of the tree.
func (b *BTree) PostOrderTraversal(fn func(n *BTreeNode)) {
	postOrderRecursive(b.Root, fn)
}

// postOrderRecursive performs a post-order traversal of the tree.
func postOrderRecursive(node *BTreeNode, fn func(n *BTreeNode)) {
	if node == nil {
		return
	}
	postOrderRecursive(node.Left, fn)
	postOrderRecursive(node.Right, fn)
	fn(node)
}

// main is the entry point for the application.
func main() {
	tree := BTree{}
	tree.ReplaceOrInsert(50)
	tree.ReplaceOrInsert(30)
	tree.ReplaceOrInsert(70)
	tree.ReplaceOrInsert(20)
	tree.ReplaceOrInsert(40)
	tree.ReplaceOrInsert(60)
	tree.ReplaceOrInsert(80)

	// BTree visualization:
	//         50
	//       /    \
	//     30      70
	//    /  \    /  \
	//  20   40  60   80

	fmt.Println("Post-order traversal output:")
	tree.PostOrderTraversal(func(n *BTreeNode) {
		fmt.Print(n.Value, " ") // Should output: 20 40 30 60 80 70 50
	})
}
