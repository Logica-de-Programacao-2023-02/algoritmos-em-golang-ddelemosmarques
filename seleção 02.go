//Faça um algoritmo que
//leia três números inteiros
//e mostre o menor deles.

package main

import "fmt"

func main() {

	var x, y, z int

	fmt.Print(" Escreva três numeros diferentes\n")
	fmt.Scan(&x, &y, &z)

	if x < y && x < z {
		fmt.Printf("%d é o menor", x)
	} else if y < x && y < z {
		fmt.Printf("%d é o menor", y)
	} else if z < x && z < y {
		fmt.Printf("%d é o menor", z)
	}
}
