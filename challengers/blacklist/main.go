package main

import (
	"fmt"

	"github.com/ferreira-gn/estrutura-de-dados/challengers/blacklist/tree"
)

func main() {
	runChallengeFlow()
}

func runChallengeFlow() {
	blacklist := tree.New()

	blockedIPs := []int32{
		192168001,
		192168002,
		192168010,
		172016001,
		172016002,
		172016050,
		100000001,
		100000002,
		100000010,
		127000001,
	}

	fmt.Println("Iniciando desafio da blacklist com BST")
	fmt.Println()

	insertBlockedIPs(blacklist, blockedIPs)
	searchBlockedIPs(blacklist, []int32{
		192168002,
		172016050,
		888888888,
	})
	printOrderedReport(blacklist)
	removeExpiredPenalty(blacklist, 172016050)
	printOrderedReport(blacklist)
}

func insertBlockedIPs(blacklist *tree.Tree, blockedIPs []int32) {
	fmt.Println("Inserindo IPs na blacklist")

	for _, ip := range blockedIPs {
		blacklist.AddNode(ip)
		fmt.Printf("IP %d inserido\n", ip)
	}

	blacklist.AddNode(192168002)
	blacklist.AddNode(192168002)
	fmt.Println("IP 192168002 recebeu duas tentativas extras")
	fmt.Println()
}

func searchBlockedIPs(blacklist *tree.Tree, ips []int32) {
	fmt.Println("Consultando requisicoes recebidas")

	for _, ip := range ips {
		foundNode := blacklist.Search(ip)

		if foundNode == nil {
			fmt.Printf("IP %d liberado\n", ip)
			continue
		}

		fmt.Printf("IP %d bloqueado | tentativas: %d\n", foundNode.IP, foundNode.Attempts())
	}

	fmt.Println()
}

func printOrderedReport(blacklist *tree.Tree) {
	blacklist.ReportInOrder()
	fmt.Println()
}

func removeExpiredPenalty(blacklist *tree.Tree, ip int32) {
	fmt.Printf("Removendo IP %d apos fim da penalidade\n", ip)
	blacklist.Remove(ip)
	fmt.Println()
}
