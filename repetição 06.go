//Faça um algoritmo que imprima a tabuada de multiplicação
// de 1 a 10 para um número fornecido pelo usuário

package main

import "fmt"

func main() {

	var n int

	fmt.Print("\nEscolha qualquer número e será mostrada sua tabuada \n")
	fmt.Scan(&n)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d \n", i*n)
	}
}
