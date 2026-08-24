package main

import (
	"fmt"

	"github.com/ferreira-gn/estrutura-de-dados/challengers/blacklist/tree"
)

func main() {
	blacklist := tree.New()

	for _, ip := range []int64{192168001001, 192168001002, 192168001010} {
		blacklist.AddNode(ip)
	}

	ip := int64(192168001002)
	if blacklist.Contains(ip) {
		fmt.Printf("IP %d bloqueado\n", ip)
		return
	}

	fmt.Printf("IP %d liberado\n", ip)
}
