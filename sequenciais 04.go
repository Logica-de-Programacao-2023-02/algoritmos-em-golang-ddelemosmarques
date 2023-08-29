//Faça um algoritmo que leia três números reais
//e mostre a média ponderada entre eles, com pesos 2, 3 e 5, respectivamente.

package main

import "fmt"

func main() {

	var n1, n2, n3 int

	var p1 = 2
	var p2 = 3
	var p3 = 5

	fmt.Println("Escreava 3 números")
	fmt.Scan(&n1, &n2, &n3)
	fmt.Println("A média ponderada destes numeros é ")
	fmt.Println((n1*p1 + n2*p2 + n3*p3) / (p1 + p2 + p3))

}
