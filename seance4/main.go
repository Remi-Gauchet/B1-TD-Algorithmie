package main

import (
	"fmt"
	"os"
)

// executer renvoie une erreur au lieu d'arrêter le programme
func executer(args []string) error {
	nombres, err := lireNombres(args)
	if err != nil {
		return err
	}
	err = verifierDoublons(nombres)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(len(args), "nombre(s) :", nombres)
	return nil
}

// main lance le programme
func main() {
	err := executer(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
		os.Exit(1)
	}
}
