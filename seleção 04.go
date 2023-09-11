//Faça um algoritmo que leia a
// altura e o sexo de uma pessoa e mostre
//se ela está abaixo, dentro ou acima do peso ideal,
//utilizando as fórmulas do exercício 3 da lista de algoritmos sequenciais.

package main

import (
	"fmt"
)

func main() {

	var p, a, imc float64
	var s string

	fmt.Printf("Qual é o seu peso? ")
	fmt.Scan(&p)
	fmt.Printf("Qual é a sua altura?")
	fmt.Scan(&a)
	fmt.Printf("Qual é seu sexo? ")
	fmt.Scan(&s)

	imc = p / (a * a)

	if imc < 18.5 {
		fmt.Printf("O cálculo IMC indica que vc esta com Baixo peso")
	} else if imc > 18.5 && imc < 24.9 {
		fmt.Printf("O cálculo IMC indica que vc esta com Peso normal")
	} else if imc > 25 && imc < 29.9 {
		fmt.Printf("O cálculo IMC indica que vc testa com Sobrepeso")
	} else if imc > 30 && imc < 34.9 {
		fmt.Printf("O cálculo IMC indica que vc testa com Obesidade grau I")
	} else if imc > 35 && imc < 39.9 {
		fmt.Printf("O cálculo IMC indica que vc testa com Obesidade grau II")
	} else {
		fmt.Printf("O cálculo IMC indica que vc testa com Obesidade grau III")
	}

}
