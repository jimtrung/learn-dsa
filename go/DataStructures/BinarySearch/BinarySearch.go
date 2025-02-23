package BinaryTree

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(value int) *Node {
    return &Node{
        data: value,
        left: nil,
        right: nil,
    }
}

type BinarySearchTree struct {
    root *Node
}

func NewBST() *BinarySearchTree {
    return &BinarySearchTree{
        root: nil,
    }
}

func PostOrderTraversal(root *Node) {
    if root == nil {
        return
    }

    PostOrderTraversal(root.left)
    PostOrderTraversal(root.right)
    fmt.Println(root.data)
}

func InOrderTraversal(root *Node) {
    if root == nil {
        return
    }

    InOrderTraversal(root.left)
    fmt.Println(root.data)
    InOrderTraversal(root.right)
}

func PreOrderTraversal(root *Node) {
    if root == nil {
        return
    }

    fmt.Println(root.data)
    PreOrderTraversal(root.left)
    PreOrderTraversal(root.right)
}

func Search(root *Node, target int) *Node {
    if root == nil {
        return &Node{}
    } else if root.data == target {
        return root
    } else if root.data > target {
        return Search(root.left, target)
    } else {
        return Search(root.right, target)
    }
}
