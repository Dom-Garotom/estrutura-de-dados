package tree

/*
IsEmpty informa se a arvore nao possui nenhum no.
*/
func (tree *Tree) IsEmpty() bool {
	return tree.root == nil
}

/*
Height retorna a altura da arvore.
*/
func (tree *Tree) Height() int32 {
	return calculateHeight(tree.root)
}

func calculateHeight(node *Node) int32 {
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

/*
Size retorna a quantidade total de nos presentes na arvore.
*/
func (tree *Tree) Size() int32 {
	return calculateSize(tree.root)
}

func calculateSize(node *Node) int32 {
	if node == nil {
		return 0
	}

	leftSize := calculateSize(node.left)
	rightSize := calculateSize(node.rigth)

	return leftSize + rightSize + 1
}

/*
NumberOfSheets retorna a quantidade de folhas da arvore.
*/
func (tree *Tree) NumberOfSheets() int32 {
	return calculateNumberOfSheets(tree.root)
}

func calculateNumberOfSheets(node *Node) int32 {
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

/*
Contains informa se o IP existe na arvore.
*/
func (tree *Tree) Contains(value int32) bool {
	if tree.root == nil {
		return false
	}

	currentNode := tree.root
	for currentNode != nil {
		if currentNode.IP == value {
			return true
		}

		if value < currentNode.IP {
			currentNode = currentNode.left
			continue
		}

		if value > currentNode.IP {
			currentNode = currentNode.rigth
			continue
		}
	}

	return false
}

/*
SearchInOrder procura um IP usando percurso em ordem.
*/
func (tree *Tree) SearchInOrder(value int32) *Node {
	return inOrder(tree.root, value)
}

func inOrder(node *Node, searchValue int32) *Node {
	if node == nil {
		return nil
	}

	if foundNode := inOrder(node.left, searchValue); foundNode != nil {
		return foundNode
	}

	if node.IP == searchValue {
		return node
	}

	if foundNode := inOrder(node.rigth, searchValue); foundNode != nil {
		return foundNode
	}

	return nil
}
