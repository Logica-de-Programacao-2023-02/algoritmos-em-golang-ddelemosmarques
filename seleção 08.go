//Faça um algoritmo que leia um número inteiro
//entre 1 e 7 e mostre o dia da semana correspondente
//(1 = domingo, 2 = segunda-feira, etc.).

package main

import "fmt"

func main() {

	dia := 0
	fmt.Print("\n\nEscolha o dia da semana com o respectivo número\n")
	fmt.Print("\t ----------------------------\n")
	fmt.Print("\t 1 - Domingo\n")
	fmt.Print("\t 2 - Segunda\n")
	fmt.Print("\t 3 - Terça\n")
	fmt.Print("\t 4 - Quarta\n")
	fmt.Print("\t 5 - Quinta\n")
	fmt.Print("\t 6 - Sexta\n")
	fmt.Print("\t 7 - Sábado\n")
	fmt.Print("\t ----------------------------\n")

	fmt.Scan(&dia)

	switch dia {
	case 1:
		fmt.Print("O dia escolhido foi  Domingo\n\n")
	case 2:
		fmt.Print("O dia escolhido foi   Segunda\n\n")
	case 3:
		fmt.Print("O dia escolhido foi  Terça\n\n")
	case 4:
		fmt.Print("O dia escolhido foi  Quarta\n\n")
	case 5:
		fmt.Print("O dia escolhido foi  Quinta\n\n")
	case 6:
		fmt.Print("O dia escolhido foi Sexta\n\n")
	case 7:
		fmt.Print("O dia escolhido foi Sábado\n\n")
	default:
		fmt.Print(" o número escolhido não é válido")
	}
}
