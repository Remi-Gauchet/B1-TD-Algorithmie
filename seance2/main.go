package main

import (
	"fmt"
	"math/rand"
)

// structure
type Piles struct {
	a          []int
	b          []int
	operations []string
}

// opérations
func (p *Piles) ra() {
	if len(p.a) < 2 {
		return
	}
	p.a = append(p.a[1:], p.a[0])
	p.operations = append(p.operations, "ra")
}

func (p *Piles) rb() {
	if len(p.b) < 2 {
		return
	}
	p.b = append(p.b[1:], p.b[0])
	p.operations = append(p.operations, "rb")
}

func (p *Piles) sa() {
	if len(p.a) < 2 {
		return
	}
	p.a[0], p.a[1] = p.a[1], p.a[0]
	p.operations = append(p.operations, "sa")
}

func (p *Piles) sb() {
	if len(p.b) < 2 {
		return
	}
	p.b[0], p.b[1] = p.b[1], p.b[0]
	p.operations = append(p.operations, "sb")
}

func (p *Piles) pb() {
	if len(p.a) == 0 {
		return
	}
	p.b = append([]int{p.a[0]}, p.b...)
	p.a = p.a[1:]
	p.operations = append(p.operations, "pb")
}

func (p *Piles) pa() {
	if len(p.b) == 0 {
		return
	}
	p.a = append([]int{p.b[0]}, p.a...)
	p.b = p.b[1:]
	p.operations = append(p.operations, "pa")
}

func (p *Piles) rra() {
	if len(p.a) < 2 {
		return
	}
	p.a = append([]int{p.a[len(p.a)-1]}, p.a[:len(p.a)-1]...)
	p.operations = append(p.operations, "rra")
}

func (p *Piles) rrb() {
	if len(p.b) < 2 {
		return
	}
	p.b = append([]int{p.b[len(p.b)-1]}, p.b[:len(p.b)-1]...)
	p.operations = append(p.operations, "rrb")
}

// affiche les piles a et b
func (p *Piles) afficher() {
	fmt.Println("a :", p.a, " b :", p.b)
}

// Si la pile est triée, resoudre dit quelles opérations ont eu lieu.
// Sinon, il dit qu'elle n'est pas triée
func (p *Piles) resoudre() {
	if p.aEstTriee() == true {
		fmt.Println("opérations :", p.operations)
	} else {
		fmt.Println("La pile n'est pas triée.")
	}
}

// lancement du programme
func main() {
	var taille int = 10
	p := Piles{a: rand.Perm(taille)}
	for len(p.a) != 0 {
		for p.a[0] != p.a[p.positionDuMinimum()] {
			if p.positionDuMinimum() < (len(p.a) / 2) {
				p.ra()
			} else {
				p.rra()
			}
		}
		p.pb()
	}
	for len(p.b) != 0 {
		p.pa()
	}
	p.afficher()
	p.resoudre()
	fmt.Println("Nombre d'opérations :", len(p.operations))
}
