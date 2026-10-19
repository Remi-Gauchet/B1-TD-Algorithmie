package main

import "fmt"

func main() {
	nombres := []int{5, 2, 9, 1, 7}
	indexmin := 0
	valeurmin := nombres[1]
	indexmax := len(nombres) - 1
	valeurmax := nombres[1]
	for i := len(nombres) - 1; i >= 0; i-- {
		if valeurmin > nombres[i] {
			valeurmin = nombres[i]
			indexmin = i
		}
		if valeurmax < nombres[i] {
			valeurmax = nombres[i]
			indexmax = i
		}
	}
	ecart := indexmax - indexmin
	if ecart < 0 {
		ecart = -ecart
	}
	fmt.Println("minimum", valeurmin, "en position", indexmin)
	fmt.Println("maximum", valeurmax, "en position", indexmax)
	fmt.Println("avec un écart de", ecart)
}
