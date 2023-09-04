//Faça um algoritmo que leia
//o número de dias trabalhados
//por um funcionário e o valor da sua diária
// e calcule o seu salário.

package main

import "fmt"

func main() {

	var dt int64
	var diaria int64

	fmt.Println("Quantos dias foram trabalhados? ")
	fmt.Scan(&dt)
	fmt.Println("Qual é o valor da diária deste funcionário? ")
	fmt.Scan(&diaria)

	fmt.Print(" O salário que o colabroarador receberá é:\n ", dt*diaria)
}
