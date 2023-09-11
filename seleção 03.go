//Faça um algoritmo que leia
//um número inteiro e mostre se
//ele é par ou ímpar.

package main

import "fmt"

func main() {

	var n int
	var p float64

	fmt.Print("escreva um numero\n ")
	fmt.Scan(&n)

	p = float64(n % 2)

	if p == 0 {
		fmt.Printf(" %d é número par ", n)
	} else {
		fmt.Printf(" %d é número impar ", n)
	}
}
