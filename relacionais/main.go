package main

import "fmt"

func main() {

	//operadores relacionais
	n1 := 12
	n2 := 12

	//comparação igualdade
	isEqual := n1 == n2 // é n1 igual a n2? - true or false
	fmt.Println(isEqual)

	//comparação diferença
	isDiff := n1 != n2
	fmt.Println(isDiff)

	//maior
	isGreater := n1 > n2
	fmt.Println(isGreater)

	//menor
	isLess := n1 < n2
	fmt.Println(isLess)

	//maior ou igual

	isGreaterOrEqual := n1 >= n2
	fmt.Println(isGreaterOrEqual)

	//menor ou igual
	isLessOrEqual := n1 <= n2
	fmt.Println(isLessOrEqual)

	s1 := "go"
	s2 := "lang"

	str := s1 < s2
	fmt.Println(str)

	//nil
	var k *int8
	var c *int8
	c = nil
	k = nil
	fmt.Println(k == c)

}
