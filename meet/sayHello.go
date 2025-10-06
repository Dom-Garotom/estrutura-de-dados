package meet

import "fmt"

// Aqui criamos uma dentro do modulo meet para que possamos usar essa funções em outro modulos diferentes.
// Funções exportadas começam com letras maiúculas 
// funções não exportadas começam com letras minúsculas

func SayHello () {
	var userName string = getName()

	fmt.Printf("\nOlá %s! \nÉ um prazer te-lo aqui comigo durante esse aprendizado de golang.", userName)
}