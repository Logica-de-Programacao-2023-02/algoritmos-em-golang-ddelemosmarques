//Faça um algoritmo que leia dois números inteiros
//e mostre o resultado da multiplicação entre eles,
//se ambos forem positivos;
//ou a soma, se pelo menos um deles for negativo

package main

import "fmt"

func main() {

	var x, y float64

	fmt.Print(" escreva dois nùmeros\n")
	fmt.Scan(&x, &y)

	if x > 0 && y > 0 {
		fmt.Printf(" %.2f é a multiplicação entre estes dois numeros poisitivos", x*y)
	} else {
		fmt.Printf(" %.2f é a soma dos dois", x+y)
	}

}
