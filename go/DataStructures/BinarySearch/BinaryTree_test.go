package BinaryTree

import (
	"bytes"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	os.Stdout = old

	return buf.String()
}

func TestTraversals(t *testing.T) {
	bst := NewBST()
	bst.root = NewNode(5)
	bst.root.left = NewNode(3)
	bst.root.right = NewNode(7)
	bst.root.left.left = NewNode(2)
	bst.root.left.right = NewNode(4)
	bst.root.right.left = NewNode(6)
	bst.root.right.right = NewNode(8)

	postOrderExpected := "2\n4\n3\n6\n8\n7\n5\n"
	inOrderExpected := "2\n3\n4\n5\n6\n7\n8\n"
	preOrderExpected := "5\n3\n2\n4\n7\n6\n8\n"

	postOrderOutput := captureOutput(func() { PostOrderTraversal(bst.root) })
	assert.Equal(t, postOrderExpected, postOrderOutput, "PostOrder traversal should match")

	inOrderOutput := captureOutput(func() { InOrderTraversal(bst.root) })
	assert.Equal(t, inOrderExpected, inOrderOutput, "InOrder traversal should match")

	preOrderOutput := captureOutput(func() { PreOrderTraversal(bst.root) })
	assert.Equal(t, preOrderExpected, preOrderOutput, "PreOrder traversal should match")
}

func TestSearch(t *testing.T) {
	bst := NewBST()
	bst.root = NewNode(5)
	bst.root.left = NewNode(3)
	bst.root.right = NewNode(7)
	bst.root.left.left = NewNode(2)
	bst.root.left.right = NewNode(4)
	bst.root.right.left = NewNode(6)
	bst.root.right.right = NewNode(8)

	targetNode := Search(bst.root, 4)
	assert.NotNil(t, targetNode, "Search should return a valid node")
	assert.Equal(t, 4, targetNode.data, "Search should find node with value 4")

	notFoundNode := Search(bst.root, 10)
	assert.NotNil(t, notFoundNode, "Search should return a non-nil node")
	assert.Equal(t, &Node{}, notFoundNode, "Search should return an empty node when target is not found")
}
