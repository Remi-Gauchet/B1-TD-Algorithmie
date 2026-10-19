package main

import "fmt"

func triSelection(nombres []int) (int, int) {
	comparaisons := 0
	echanges := 0
	for i := 0; i < len(nombres); i++ {
		indexMin := 0
		if indexMin != i {
			for j := i + 1; j < len(nombres); j++ {
				if nombres[j] < nombres[indexMin] {
					indexMin = j
				}
				comparaisons++
			}
			echanges++
			nombres[i], nombres[indexMin] = nombres[indexMin], nombres[i]
			fmt.Println(nombres, "nombre selectionné", nombres[i])
		}
	}
	return comparaisons, echanges
}

func main() {
	nombres := []int{1, 2, 3, 4, 5}
	comparaisons, echanges := triSelection(nombres)
	fmt.Println("comparaisons :", comparaisons)
	fmt.Println("échanges :", echanges)
}
