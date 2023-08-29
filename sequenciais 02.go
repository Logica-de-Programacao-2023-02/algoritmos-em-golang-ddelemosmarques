//Faça um algoritmo que leia um número inteiro e mostre o seu dobro, triplo e quadruplo.

package main

import "fmt"

func main() {

	var n1 int

	fmt.Println("Escreva um número ")
	fmt.Scan(&n1)
	fmt.Println(" o dobro de ", n1, "é", n1*2, " o triplo é ", n1*3, "o quadruplo é ", n1*4)

}
