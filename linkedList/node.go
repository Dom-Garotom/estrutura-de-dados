package linkedlist

type Node struct {
	data     string
	nextNode *Node
}

/*
NewNode é uma função responsável por criar um novo nô

	Retorna um novo nô
*/
func NewNode(value string) *Node {
	return &Node{
		data: value,
	}
}
