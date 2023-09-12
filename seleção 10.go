//Faça um algoritmo que leia a idade de uma pessoa e mostre
//a sua classificação, de acordo com a seguinte tabela:
//
//até 9 anos: mirim
//10 a 13 anos: infantil
//14 a 17 anos: juvenil
//maiores de 18 anos: adulto

package main

import "fmt"

func main() {

	var i int

	fmt.Printf("\n Qual é a idade da pessoa? \n")
	fmt.Scan(&i)

	if i < 9 && i != 0 {
		fmt.Printf("mirim")
	} else if i >= 10 && i <= 13 {
		fmt.Printf("infantil")
	} else if i >= 14 && i <= 17 {
		fmt.Printf("juvenil")
	} else if i > 18 {
		fmt.Printf("adulto")
	} else {
		fmt.Printf(" caractere inválido")

	}

}
