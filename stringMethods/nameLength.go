package stringmethods

import (
	"fmt"
	"unicode/utf8"
)

func NameLenght (userName string) {
	fmt.Printf("\n\nGostaria de saber quantas letras tem no seu nome ?\n")

	var nameLenght int = utf8.RuneCountInString(userName)
	var memoryStorage int = len(userName)

	fmt.Printf("Seu nome tem %d letras e tem %d bytes em memôria\n", nameLenght , memoryStorage)
}