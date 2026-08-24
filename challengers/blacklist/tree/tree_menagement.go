package tree

import "fmt"

/*
AddNode adiciona um IP na arvore.
Se o IP ja existir, incrementa o contador de tentativas do no.
*/
func (tree *Tree) AddNode(value int32) {
	newNode := createNodeTree(value)

	if tree.root == nil {
		tree.root = newNode
		return
	}

	currentNode := tree.root
	for {

		if currentNode.IP == value {
			currentNode.attempts++
			return
		}

		if value < currentNode.IP {
			if currentNode.left == nil {
				currentNode.left = newNode
				return
			}

			currentNode = currentNode.left
			continue
		}

		if value > currentNode.IP {

			if currentNode.rigth == nil {
				currentNode.rigth = newNode
				return
			}

			currentNode = currentNode.rigth
			continue
		}
	}
}

/*
Remove remove um IP da arvore, quando ele existir.
*/
func (tree *Tree) Remove(ip int32) {
	tree.root = removeNode(tree.root, ip)
}

func removeNode(node *Node, ip int32) *Node {
	if node == nil {
		return nil
	}

	if ip < node.IP {
		node.left = removeNode(node.left, ip)
		return node
	}

	if ip > node.IP {
		node.rigth = removeNode(node.rigth, ip)
		return node
	}

	if ip == node.IP {
		if node.left == nil {
			return node.rigth
		}

		if node.rigth == nil {
			return node.left
		}

		successor := findMin(node.rigth)
		node.IP = successor.IP
		node.attempts = successor.attempts
		node.rigth = removeNode(node.rigth, successor.IP)

		return node
	}

	return node
}

func findMin(node *Node) *Node {
	currentNode := node

	for currentNode.left != nil {
		currentNode = currentNode.left
	}

	return currentNode
}

/*
ReportInOrder imprime os IPs bloqueados em ordem crescente.
*/
func (tree *Tree) ReportInOrder() {
	fmt.Println("RELATORIO DE IPS BLOQUEADOS")

	if tree.root == nil {
		fmt.Println("Nenhum IP bloqueado.")
		return
	}

	fmt.Printf("%-15s %-10s\n", "IP", "TENTATIVAS")
	printInOrder(tree.root)
}

/*
printInOrder percorre a arvore em ordem: esquerda, no atual e direita.
*/
func printInOrder(node *Node) {
	if node == nil {
		return
	}

	printInOrder(node.left)
	fmt.Printf("%-15d %-10d\n", node.IP, node.attempts)
	printInOrder(node.rigth)
}

/*
Search procura um IP na arvore usando a regra da BST.
*/
func (tree *Tree) Search(value int32) *Node {
	if tree.root == nil {
		return nil
	}

	currentNode := tree.root
	for currentNode != nil {
		if currentNode.IP == value {
			return currentNode
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

	return nil
}
