//Faça um algoritmo que leia a idade de uma pessoa em anos e mostre a idade em dias.

package main

import "fmt"

func main() {

	var age int

	fmt.Println("Quantos anos vc tem? ")
	fmt.Scan(&age)
	fmt.Println("Vc já viveu mais de ", age*365, "dias ")

}
