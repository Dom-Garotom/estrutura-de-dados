/*
Package tree is a package of binary tree
*/
package tree

type Tree struct {
	root *Node
}

func New() *Tree {
	return &Tree{
		root: nil,
	}
}

type Node struct {
	left     *Node
	rigth    *Node
	Content  int64
	attempts int32
}

func createNodeTree(content int64) *Node {
	return &Node{
		Content: content,
	}
}
