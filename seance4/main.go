package main

import (
	"fmt"
	"os"

	"github.com/Remi-Gauchet/Algorithmie/seance4/outils"
)

// structure
type Piles struct {
	a          []int
	b          []int
	operations []string
}

// exécute les fonctions et renvoie une erreur
func executer(args []string) error {
	//on lit l'argument
	nombres, err := outils.LireNombres(args)
	if err != nil {
		return err
	}

	//on vérifie qu'il y ait bien un argument
	if len(nombres) == 0 {
		return nil
	}

	//on n'accepte pas les doublons
	err = outils.VerifierDoublons(nombres)
	if err != nil {
		return err
	}

	p := Piles{a: nombres}
	// une pile déjà triée est renvoyée immédiatement
	if p.aEstTriee() {
		return nil
	}

	// on commence le tri
	for len(p.a) > 1 {
		p.amenerAuSommet(p.positionDuMinimum())
		p.pb()
	}
	for len(p.b) != 0 {
		p.pa()
	}
	if p.aEstTriee() == false || len(p.b) > 0 {
		return fmt.Errorf("impossible de trier la pile")
	}
	for _, op := range p.operations {
		fmt.Println(op)
	}
	return nil
}

// Lance le programme et l'arrête en renvoyant une erreur si besoin
func main() {
	err := executer(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
		os.Exit(1)
	}
}
