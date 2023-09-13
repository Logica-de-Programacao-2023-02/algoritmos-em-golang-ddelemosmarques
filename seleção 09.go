//Faça um algoritmo que leia três números reais
// e mostre-os em ordem crescente.

package main

import "fmt"

func main() {

	var x, y, z int

	fmt.Print("\nDigite 3 números diferentes\n\n")
	fmt.Scan(&x, &y, &z)

	if x < y && y < z {
		fmt.Printf("%d,%d,%d", x, y, z)
	} else if x < z && z < y {
		fmt.Printf("%d,%d,%d", x, z, y)
	} else if y < x && x < z {
		fmt.Printf("%d,%d,%d", y, x, z)
	} else if y < z && z < x {
		fmt.Printf("%d,%d,%d", y, z, x)
	} else if z < x && x < y {
		fmt.Printf("%d,%d,%d", z, x, y)
	} else if z < y && y < x {
		fmt.Printf("%d,%d,%d", z, y, x)
	} else {
		fmt.Printf("Esccreva 3 numeros diferentes")
	}
}
