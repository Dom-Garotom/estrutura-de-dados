package main

import (
	"helloWOrld/meet"
	stringmethods "helloWOrld/stringMethods"
)

func main () {
	var userName = meet.SayHello()
	stringmethods.NameLenght(userName)
}