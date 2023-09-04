//Faça um algoritmo que leia
//o preço de um produto
//e mostre o seu valor com desconto de 10%

package main

import "fmt"

func main() {

	var p, d float64

	fmt.Print(" Valor do produto:\n ")
	fmt.Scan(&p)

	d = p * 0.9
	fmt.Printf(" Valor com desconto de 10 por cento: \n %.2f ", d)

}
