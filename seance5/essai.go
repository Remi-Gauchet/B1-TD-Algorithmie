package main

import (
	"fmt"
)

// essai déroule les deux phases sur l'exemple de la partie 1.
func essai() {
	p := Piles{a: []int{4, 1, 5, 0, 3, 2}}

	p.pousserParBlocs(2)
	fmt.Println("après la phase 1 - a :", p.a, " b :", p.b)

	p.remonterLesPlusGrands()
	fmt.Println("après la phase 2 - a :", p.a, " b :", p.b)

	fmt.Println("opérations :", p.operations)
	fmt.Println("nombre d'opérations :", len(p.operations))
}
