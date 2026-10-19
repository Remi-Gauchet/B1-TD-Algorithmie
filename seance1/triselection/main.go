package main

import (
	"fmt"
)

func afficherEtat(nombres []int, tour int) {
	fmt.Print("tour ", tour, " : [")
	for k, valeur := range nombres {
		if k == tour {
			fmt.Print("| ")
		}
		fmt.Print(valeur)
		if k != len(nombres)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")
}

func triSelection(nombres []int) (int, int) {
	comparaisons := 0
	echanges := 0
	for i := 0; i < len(nombres); i++ {
		indexMin := i

		for j := i + 1; j < len(nombres); j++ {
			if nombres[j] < nombres[indexMin] {
				indexMin = j
			}
			comparaisons++
		}
		echanges++
		nombres[i], nombres[indexMin] = nombres[indexMin], nombres[i]
		afficherEtat(nombres, i+1)
	}
	return comparaisons, echanges
}

func main() {
	nombres := []int{5, 2, 9, 1, 7}
	comparaisons, echanges := triSelection(nombres)
	fmt.Println("comparaisons :", comparaisons)
	fmt.Println("échanges :", echanges)
}
