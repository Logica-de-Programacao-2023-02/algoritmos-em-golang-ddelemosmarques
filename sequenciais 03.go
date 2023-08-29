//Faça um algoritmo que leia o peso e a altura de uma pessoa e calcule o seu IMC (Índice de Massa Corporal).

package main

import (
	"fmt"
	"math"
)

func main() {

	var p1, a1 float64
	fmt.Println("Qual é o seu peso? ")
	fmt.Scan(&p1)
	fmt.Println("Qual é a sua altura? ")
	fmt.Scan(&a1)

	fmt.Println("seu imc é ", p1/math.Pow(a1, 2))

}
