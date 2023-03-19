package main

import "fmt"

func main() {

	//aritmeticos

	n1 := 10
	n2 := 2
	//op soma
	soma := n1 + n2
	fmt.Println(soma)

	//op sub
	subtracao := n1 - n2
	fmt.Println(subtracao)

	//op multiplacao
	mult := n1 * n2
	fmt.Println(mult)

	//op divisao
	div := n1 / n2
	fmt.Println(div)

	//op modulo
	mod := n1 % n2
	fmt.Println(mod)

	//op incremento
	n1++ //op unario
	fmt.Println(n1)

	//op decremento
	n1-- //op unario
	fmt.Println(n1)

	//atribuicoes
	n1 += 1
	n1 -= 1
	n1 *= 1
	n1 /= 1
	n1 %= 2

}
