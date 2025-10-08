package linkedList

/*
Nô de dados que aponta para o próximo elemento de um lista encadeada.
*/
type Node struct {
	data     string
	nextNode *Node
}

/*
Função responsável por criar um novo nô
	Retorna um novo nô
*/
func NewNode(value string) *Node {
	return &Node{
		data: value,
	}
}
