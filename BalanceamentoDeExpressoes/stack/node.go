package stack

type Node struct {
	data     string
	nextNode *Node
}

/*
NewNode é uma função responsável por criar um novo elemento

	Retorna um novo nô
*/
func NewNode(value string) *Node {
	return &Node{
		data: value,
	}
}
