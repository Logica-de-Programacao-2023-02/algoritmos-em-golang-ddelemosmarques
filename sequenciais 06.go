///Faça um algoritmo que leia o salário de um funcionário
//e calcule o seu novo salário com um aumento de 15%.

package main

import "fmt"

func main() {

	var s1 float64

	fmt.Println("Quanto vc ganha? ")
	fmt.Scan(&s1)
	fmt.Println("Parabens,vc ganhou um aumento de 15% em seu salario.")
	fmt.Println("De R$", s1, "vc passará para R$", s1+(s1*0.15))

}
