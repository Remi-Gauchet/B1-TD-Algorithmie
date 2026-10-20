package main

import (
	"fmt"
	"math/rand"
)

func tranche(n int) []int {
	resultat := make([]int, n)
	for i := 0; i < n; i++ {
		resultat[i] = rand.Intn(100)
	}
	return resultat
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
	}
	return comparaisons, echanges
}

func triBulles(nombres []int) (int, int) {
	var compteurEchange int = 0
	var comparaison int = 0
	for {
		echangeFait := false
		for i := 0; i < len(nombres)-1; i++ {
			if nombres[i] > nombres[i+1] {
				nombres[i], nombres[i+1] = nombres[i+1], nombres[i]
				echangeFait = true
				compteurEchange++
			}
			comparaison++
		}
		if !echangeFait {
			break
		}
	}
	return comparaison, compteurEchange
}

func main() {
	tailles := []int{5, 10, 20, 50}

	for _, n := range tailles {
		original := tranche(n)
		copie1 := make([]int, len(original))
		copy(copie1, original)
		comparSelect, echangSelect := triSelection(copie1)

		copie2 := make([]int, len(original))
		copy(copie2, original)
		comparBulle, echangBulle := triBulles(copie2)
		fmt.Println("n=", n, "selection:", comparSelect, "comp /", echangSelect, "ech   bulles:", comparBulle, "comp /", echangBulle, "ech")
	}

}
