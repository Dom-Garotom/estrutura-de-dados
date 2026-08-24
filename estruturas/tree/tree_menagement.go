package tree

/*
* Arquivo que contem todas as funções responsáveis por gerenciar à árvore binária
* Entre eles temos :
*
* AddNode - Adiciona um novo nó na árvore
* SearchInOrder - Busca por um nó na árvore
*
*
*
* */

func (tree *Tree) AddNode(value int64) {
	newNode := createNodeTree(value)

	if tree.root == nil {
		tree.root = newNode
		return
	}

	currentNode := tree.root
	for currentNode != nil {

		if value < currentNode.Content {

			if currentNode.left != nil {
				currentNode = currentNode.left
				continue
			}

			currentNode.left = newNode
			currentNode = nil
			continue
		}

		if value > currentNode.Content {

			if currentNode.rigth != nil {
				currentNode = currentNode.rigth
				continue
			}

			currentNode.rigth = newNode
			currentNode = nil
			continue
		}
	}
}

// excluir
// rebalancear
// imprimir
//

func (tree *Tree) SearchInOrder(value int64) *Node {
	return inOrder(tree.root, value)
}

func inOrder(node *Node, searchValue int64) *Node {
	if node == nil {
		return nil
	}

	if foundNode := inOrder(node.left, searchValue); foundNode != nil {
		return foundNode
	}

	if node.Content == searchValue {
		return node
	}

	if foundNode := inOrder(node.rigth, searchValue); foundNode != nil {
		return foundNode
	}

	return nil
}
