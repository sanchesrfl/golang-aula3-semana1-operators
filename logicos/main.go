package main

import (
	"fmt"
)

func main() {

	//operadores logicos
	//podemos utilizar todos os tipos de dados

	s1 := "go"
	s2 := "lang"

	//conjuncao
	// s1 é igual a Go? E S2 é igual a lang?
	//c := (s1 == "go" && s2 == "lang") // da expr é true

	//fmt.Println(c)

	//disjunção || -- OU
	disj := (s1 == "go" || s2 == "lang")
	fmt.Println(disj)

	//negacao !=
	neg := s1 != s2
	fmt.Println(neg)

}
