package main

import "fmt"

func mergeSort(v []int) []int {
	if len(v) <= 1 {
		return v
	}
	m := len(v) / 2
	a := mergeSort(v[:m])
	b := mergeSort(v[m:])

	r := []int{}
	for len(a) > 0 && len(b) > 0 {
		if a[0] <= b[0] {
			r = append(r, a[0])
			a = a[1:]
		} else {
			r = append(r, b[0])
			b = b[1:]
		}
	}
	r = append(r, a...)
	r = append(r, b...)
	return r
}

func main() {
	v := []int{7, 4, 2, 1, 8, 3, 6, 5}
	fmt.Println(mergeSort(v))
}
