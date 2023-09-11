//Faça um algoritmo que leia o nome,
//a idade e o peso de uma pessoa
//e imprima esses valores na tela.

package main

import "fmt"

func main() {
	var nome string
	var idade int
	var peso float64
	fmt.Println("qual é o seu nome? ")
	fmt.Scan(&nome)
	fmt.Println("qual a sua idade? ")
	fmt.Scan(&idade)
	fmt.Println("Quantos quilos você tem? ")
	fmt.Scan(&peso)
	fmt.Println("Confira seus dados ")
	fmt.Println("seu nome é", nome)
	fmt.Println("vc tem", idade, "anos")
	fmt.Println("vc tem", peso, "Kgs")

}
