package tree

/*
* Arquivo que contem todas as funções utilitárias do pacote tree
* Entre eles temos :
*
* IsEmpty - verifica se à árvore está vázia
* Height - verifica à altura da árvore
* Size - Contabiliza a quantidade de nôs presentes na árvore
* NumberOfSheets - contabiliza a quantidade de folhas presentes na árvore
* Contains - verifica se o valor existe na árvore
* Search - procura pelo valor na árvore
*
* */

func (tree *Tree) IsEmpty() bool {
	return tree.root == nil
}

func (tree *Tree) Height() int64 {
	return calculateHeight(tree.root)
}

func calculateHeight(node *Node) int64 {
	if node == nil {
		return 0
	}

	leftHeight := calculateHeight(node.left)
	rightHeight := calculateHeight(node.rigth)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}

func (tree *Tree) Size() int64 {
	return calculateSize(tree.root)
}

func calculateSize(node *Node) int64 {
	if node == nil {
		return 0
	}

	leftSize := calculateSize(node.left)
	rightSize := calculateSize(node.rigth)

	return leftSize + rightSize + 1
}

func (tree *Tree) NumberOfSheets() int64 {
	return calculateNumberOfSheets(tree.root)
}

func calculateNumberOfSheets(node *Node) int64 {
	if node == nil {
		return 0
	}

	leftSide := calculateNumberOfSheets(node.left)
	rightSide := calculateNumberOfSheets(node.rigth)

	if node.left == nil && node.rigth == nil {
		return 1
	}

	return leftSide + rightSide
}

func (tree *Tree) Contains(value int64) bool {
	if tree.root.Content == value {
		return true
	}

	if tree.root == nil {
		return false
	}

	currentNode := tree.root
	for currentNode != nil {

		if value < currentNode.Content {

			if currentNode.left == nil {
				break
			}

			if currentNode.left.Content == value {
				return true
			}

			currentNode = currentNode.left
			continue
		}

		if value > currentNode.Content {

			if currentNode.rigth == nil {
				break
			}

			if currentNode.rigth.Content == value {
				return true
			}

			currentNode = currentNode.rigth
			continue
		}
	}

	return false
}

func (tree *Tree) Search(value int64) *Node {
	if tree.root.Content == value {
		return tree.root
	}

	if tree.root == nil {
		return nil
	}

	currentNode := tree.root
	for currentNode != nil {

		if value < currentNode.Content {

			if currentNode.left == nil {
				break
			}

			if currentNode.left.Content == value {
				return currentNode.left
			}

			currentNode = currentNode.left
			continue
		}

		if value > currentNode.Content {

			if currentNode.rigth == nil {
				break
			}

			if currentNode.rigth.Content == value {
				return currentNode.rigth
			}

			currentNode = currentNode.rigth
			continue
		}
	}

	return nil
}
