//Faça um algoritmo que leia um número inteiro positivo
//e mostre todos os seus divisores.

package main

import "fmt"

func main() {

	var n int

	fmt.Printf("Escreva um número lhe será retornado seus divisores\n ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, "\n")
		}

	}

}
