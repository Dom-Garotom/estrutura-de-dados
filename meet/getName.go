package meet

import "fmt"


func getName () string {
	fmt.Printf("Qual o seu nome meu nobre ?\n")
	var name string
	
	// Passamos o endereço da nossa váriavel 
	// para que a função possa atribuir o valor digitado no terminal.
	fmt.Scan(&name)

	return name
}