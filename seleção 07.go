// Faça um algoritmo que leia o salário de um funcionário
// e calcule o seu novo salário com um aumento de 10% se o salário for menor que R$ 1000,00;
// ou de 5% se o salário for maior ou igual a R$ 1000,00.
package main

import "fmt"

func main() {

	var salario float64

	fmt.Print(" qual é o salário do funcionário? ")
	fmt.Scan(&salario)

	if salario < 1000 {
		fmt.Printf(" o novo salario é R$%.2f", salario*1.10)
	} else {
		fmt.Printf(" o novo salario é R$%.2f", salario*1.05)
	}

}
