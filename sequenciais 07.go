//Faça um algoritmo que leia
//um número inteiro e mostre
//o seu sucessor e antecessor.

package main

import "fmt"

func main() {

	var n int
	fmt.Print("Escreva um número \n")
	fmt.Scan(&n)
	fmt.Print("O antecessor do numero ", n, " é ", n-1, " e o sucessor é ", n+1)
}
