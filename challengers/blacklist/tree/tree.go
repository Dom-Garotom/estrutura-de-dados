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
	IP       int32
	attempts int32
}

func (node *Node) Attempts() int32 {
	return node.attempts
}

func createNodeTree(ip int32) *Node {
	return &Node{
		IP:       ip,
		attempts: 1,
	}
}
