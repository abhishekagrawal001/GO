package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func preorderTraversal(root *Node) {
	if root != nil {
		fmt.Printf("%d->", root.data)
		preorderTraversal(root.left)
		preorderTraversal(root.right)
	}
}

func inorderTraversal(root *Node) {
	if root != nil {
		inorderTraversal(root.left)
		fmt.Printf("%d->", root.data)
		inorderTraversal(root.right)
	}
}

func postorderTraversal(root *Node) {
	if root != nil {
		postorderTraversal(root.left)
		postorderTraversal(root.right)
		fmt.Printf("%d->", root.data)
	}
}

func main() {
	root := &Node{10, nil, nil}
	root.left = &Node{20, nil, nil}
	root.right = &Node{30, nil, nil}
	root.left.left = &Node{40, nil, nil}
	root.left.right = &Node{50, nil, nil}
	root.right.right = &Node{60, nil, nil}
	root.left.left.left = &Node{70, nil, nil}

	fmt.Printf("Pre Order Traversal of the given tree is: ")
	preorderTraversal(root)
	fmt.Println()

	fmt.Printf("In Order Traversal of the given tree is: ")
	inorderTraversal(root)
	fmt.Println()

	fmt.Printf("Post Order Traversal of the given tree is: ")
	postorderTraversal(root)
}
