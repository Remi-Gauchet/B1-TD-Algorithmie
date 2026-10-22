package main

import (
	"fmt"
	"os"
)

// renvoie une erreur au lieu d'arrêter le programme
func executer(args []string) error {
	nombres, err := lireNombres(args)
	if err != nil {
		return err
	}
	fmt.Println(len(args), "argument(s) :", nombres)
	return nil
}

// lance le programme
func main() {
	err := executer(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
		os.Exit(1)
	}
}
