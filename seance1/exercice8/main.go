package main

import (
	"fmt"
)

func afficherEtat(nombres []int, tour int) {
	fmt.Println("tour", tour, nombres)
}

func triBulles(nombres []int) {
	var compteurEchange int = 0
	var compteurPassage int = 0
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
		compteurPassage++
		afficherEtat(nombres, compteurPassage)
	}
	fmt.Println("Passage :", compteurPassage)
	fmt.Println("Echange :", compteurEchange)
	fmt.Println("Comparaison :", comparaison)
}

func main() {
	nombres := []int{4, 1, 3, 2}
	triBulles(nombres)
	fmt.Println(nombres)
}
