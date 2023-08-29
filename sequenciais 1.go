// 01 Faça um algoritmo que leia três números inteiros e mostre a soma entre eles

package main

import "fmt"

func main() {
	var n1, n2, n3 int

	fmt.Println("Escreva 3 números ")
	fmt.Scan(&n1)
	fmt.Println("Escreva o segundo ")
	fmt.Scan(&n2)
	fmt.Println("Escreva o terceiro ")
	fmt.Scan(&n3)

	fmt.Println("A soma deles é ", n1+n2+n3)

}
