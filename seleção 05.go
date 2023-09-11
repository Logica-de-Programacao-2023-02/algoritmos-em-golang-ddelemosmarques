//Faça um algoritmo que leia um número inteiro
//e mostre se ele é múltiplo de 3 e de 5 ao mesmo tempo.

package main

import "fmt"

func main() {

	var n int
	fmt.Print("Escreva um número\n")
	fmt.Scan(&n)

	if n%5 == 0 && n%3 == 0 {
		fmt.Printf(" %d é um numero multipo de 5 e de 3 ao mesmo tempo", n)
	} else {
		fmt.Printf(" %d NÃO é numero multipo de 5 e de 3 ao mesmo tempo", n)
	}
}
