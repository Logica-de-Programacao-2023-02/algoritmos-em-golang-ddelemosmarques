//Faça um algoritmo que leia
//o peso de uma pessoa em quilos
//e converta para libras.

package main

import "fmt"

func main() {

	var p, l float64

	fmt.Print(" qual é o seu peso?\n ")
	fmt.Scan(&p)

	l = p * 2.20462

	fmt.Printf("seu peso em libra é %.2f", l)

}
