# Séance 4 — Push-Swap devient un vrai programme

**Algorithmie · Bachelor 1 · TDP 4 h · Ycode A2627_4659 · Go**

Ce document se suffit à lui-même. Lisez chaque partie, tapez le code, faites les exercices dans l'ordre. Chaque exercice indique ce que votre programme doit afficher.

| Temps | Contenu |
|---|---|
| 0 h 00 | Partie 0 — Où en est-on ? |
| 0 h 15 | Partie 1 — Lire les nombres en arguments |
| 1 h 00 | Partie 2 — Refuser les entrées invalides |
| 1 h 40 | Pause |
| 1 h 55 | Partie 3 — Afficher le résultat comme un vrai Push-Swap |
| 2 h 25 | Partie 4 — Les petits cas : 2, 3 et 5 nombres |
| 3 h 15 | Partie 5 — Les rangs |
| 3 h 40 | Synthèse |

---

# Partie 0 — Où en est-on ?

## Ce que vous avez déjà

En séance 2, vous avez construit un **simulateur** : deux piles, les huit opérations, et une méthode qui trie la pile a en amenant le minimum au sommet, avec `ra` ou `rra` selon le chemin le plus court.

En séance 3, vous avez vu pourquoi cette méthode coûte cher : elle cherche le minimum encore et encore, et son nombre de mouvements grandit à peu près comme n². Vous avez aussi découvert deux idées : les **rangs** et les **blocs**.

## Ce qui manque encore

Votre programme n'est pas encore un vrai Push-Swap. Trois défauts :

1. **Les nombres sont écrits en dur dans le code.** Un vrai Push-Swap les reçoit au lancement : `go run main.go 5 2 9 1 7`.
2. **Il ne refuse rien.** Que se passe-t-il avec `go run main.go 5 x 1` ? Ou avec deux fois le même nombre ?
3. **Il gaspille des mouvements sur les petites piles.** Trier `3 1 2` demande 1 mouvement, pas 3.

C'est l'objet de cette séance. La stratégie par blocs viendra en séance 5.

## Créer le dossier

Créez un dossier `seance4`, et dedans un fichier `main.go`. Vous le lancerez avec :

```
go run main.go 5 2 9 1 7
```

Tout ce qui suit `main.go` est transmis à votre programme.

---

# Partie 1 — Lire les nombres en arguments

## Les arguments du programme

Les mots tapés après la commande arrivent dans `os.Args`, une tranche de textes :

| Index | Contenu avec `go run main.go 5 2 9` |
|---|---|
| `os.Args[0]` | le chemin du programme lui-même |
| `os.Args[1]` | `"5"` |
| `os.Args[2]` | `"2"` |
| `os.Args[3]` | `"9"` |

Le premier élément est toujours le programme : les vrais arguments commencent à l'index 1. On les récupère avec `os.Args[1:]`.

**Attention : ce sont des textes**, pas des nombres. `"5"` n'est pas `5`. On ne peut pas les comparer ni les trier tels quels.

## Transformer un texte en nombre

```go
nombre, err := strconv.Atoi("5")
```

`strconv.Atoi` renvoie **deux** valeurs : le nombre, et une erreur. Si le texte n'est pas un nombre entier, `err` n'est pas `nil`.

```go
nombre, err := strconv.Atoi("x")
if err != nil {
	// "x" n'est pas un nombre
}
```

## La fonction `executer`

Pour garder `main` très court, on met toute la logique dans une fonction `executer`, qui **renvoie une erreur** au lieu d'arrêter le programme :

```go
func executer(args []string) error {
	// tout le travail ici
	return nil
}

func main() {
	err := executer(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
		os.Exit(1)
	}
}
```

`main` ne fait plus que trois choses : appeler `executer`, afficher l'erreur s'il y en a une, et arrêter le programme avec le code 1.

`os.Stderr` est la **sortie d'erreur** : le canal réservé aux messages d'erreur. Ils restent visibles à l'écran, mais ne se mélangent pas avec les résultats.

---

## Exercice 1

Écrivez un programme qui affiche le nombre d'arguments reçus, puis la liste des arguments.

**Avec `go run main.go 5 2 9 1 7`, votre programme doit afficher** :

```
5 argument(s) : [5 2 9 1 7]
```

**Vérification** : sans argument, le programme affiche `0 argument(s) : []` sans planter.

**Extension** : lancez `go run main.go "5 2" 9`. Combien d'arguments ? Pourquoi ?

---

## Exercice 2

Écrivez une fonction `lireNombres(args []string) ([]int, error)` qui transforme chaque argument en nombre avec `strconv.Atoi`.

Si un argument n'est pas un nombre entier, la fonction renvoie immédiatement une erreur qui le nomme. Pour fabriquer ce message, utilisez :

```go
return nil, fmt.Errorf("%q n'est pas un nombre entier", texte)
```

`%q` affiche le texte entre guillemets : on voit ainsi même un argument vide.

Appelez `lireNombres` depuis `executer`, et affichez les nombres obtenus.

**Avec `go run main.go 5 2 9 1 7`, votre programme doit afficher** :

```
nombres : [5 2 9 1 7]
```

**Avec `go run main.go 5 x 1`, il doit afficher** :

```
Erreur : "x" n'est pas un nombre entier
exit status 1
```

La dernière ligne est ajoutée par `go run` : elle confirme que le programme s'est arrêté avec le code 1.

**Extension** : testez `3.5`, `-4` et `+4`. Lesquels sont acceptés ?

---

# Partie 2 — Refuser les entrées invalides

## Pourquoi c'est important

Un programme qui reçoit des données doit les vérifier **avant** de travailler dessus. Sinon, il produit un résultat faux, ou il plante loin de la cause du problème.

Pour Push-Swap, trois situations sont à prévoir :

| Situation | Exemple | Comportement attendu |
|---|---|---|
| Un argument n'est pas un nombre | `5 x 1` | Message d'erreur, code 1 |
| Un nombre apparaît deux fois | `5 2 5` | Message d'erreur, code 1 |
| Aucun argument | `go run main.go` | Rien n'est affiché, code 0 : il n'y a rien à trier |

## Repérer un doublon avec une map

Une **map** associe une valeur à une clé. Ici, on s'en sert comme d'une liste de nombres déjà rencontrés :

```go
vus := map[int]bool{}
```

`map[int]bool` : des clés qui sont des nombres, des valeurs vrai ou faux. Une clé absente vaut `false`.

```go
for _, nombre := range nombres {
	if vus[nombre] {
		// déjà vu : c'est un doublon
	}
	vus[nombre] = true
}
```

La première fois qu'on lit 5, `vus[5]` vaut `false`, et on le passe à `true`. La deuxième fois, il vaut `true` : c'est un doublon.

**Pourquoi refuser les doublons ?** Avec deux 5, l'ordre entre eux n'a pas de sens, et la notion de rang que vous utiliserez en partie 5 ne fonctionne plus. Le sujet Push-Swap les interdit.

---

## Exercice 3

Écrivez une fonction `verifierDoublons(nombres []int) error` qui renvoie une erreur dès qu'un nombre apparaît deux fois.

Appelez-la dans `executer`, juste après `lireNombres`.

**Avec `go run main.go 5 2 5`, votre programme doit afficher** :

```
Erreur : le nombre 5 apparaît plusieurs fois
exit status 1
```

**Vérification** : `go run main.go 5 2 9 1 7` affiche toujours `nombres : [5 2 9 1 7]`.

**Extension** : lancez `go run main.go 05 5`. Est-ce un doublon ? Pourquoi ?

---

## Exercice 4

Ajoutez au début d'`executer` le cas sans argument : si `args` est vide, la fonction renvoie `nil` immédiatement.

Puis recopiez et complétez ce tableau en lançant chaque commande :

| Commande | Affichage obtenu | Code de sortie |
|---|---|---|
| `go run main.go 5 2 9 1 7` | | |
| `go run main.go` | | |
| `go run main.go 5 x 1` | | |
| `go run main.go 3.5` | | |
| `go run main.go 5 2 5` | | |
| `go run main.go -3 10 0` | | |

Pour lire le code de sortie juste après une commande : `echo $?` sous Mac ou Git Bash, `echo $LASTEXITCODE` sous PowerShell.

**Vérification** : aucun cas ne fait planter le programme, et seuls les cas invalides renvoient le code 1.

---

# Partie 3 — Afficher le résultat comme un vrai Push-Swap

## Ce que le programme doit afficher

Un Push-Swap affiche **uniquement** la liste des opérations, **une par ligne**, et rien d'autre :

```
rra
rra
pb
...
```

Pas de pile, pas de titre, pas de compteur. Cette sortie est destinée à être lue par un autre programme, qui rejoue les opérations pour vérifier le résultat. Tout texte en plus la rendrait illisible pour lui.

## Compter les lignes

Pour connaître le nombre d'opérations, on compte les lignes affichées :

```
go run main.go 5 2 9 1 7 | wc -l
```

sous Mac ou Git Bash, et sous PowerShell :

```
go run main.go 5 2 9 1 7 | Measure-Object -Line
```

Le symbole `|` envoie l'affichage du programme vers la commande qui compte les lignes.

## Brancher le simulateur

Copiez dans votre `main.go` le code de la séance 2 : la structure `Piles`, les huit opérations, `aEstTriee`, `positionDuMinimum` et votre méthode générale. Si vous ne l'avez plus, il est en annexe à la fin de ce document.

Profitez-en pour ranger le choix du sens de rotation dans une fonction à part, qui resservira :

```go
// amenerAuSommet fait remonter l'élément de a en position donnée,
// par le chemin le plus court.
func (p *Piles) amenerAuSommet(position int) {
	if position < len(p.a)/2 {
		for i := 0; i < position; i++ {
			p.ra()
		}
	} else {
		for i := 0; i < len(p.a)-position; i++ {
			p.rra()
		}
	}
}
```

---

## Exercice 5

Dans `executer`, à la place de l'affichage des nombres :

1. créez une pile a à partir des nombres : `p := Piles{a: nombres}` ;
2. triez-la avec votre méthode générale, en utilisant `amenerAuSommet` ;
3. vérifiez que a est triée et b vide, sinon renvoyez une erreur ;
4. affichez chaque opération sur une ligne.

**Avec `go run main.go 5 2 9 1 7`, votre programme doit afficher 14 lignes**, dont les quatre dernières sont `pa`.

**Vérification** : `go run main.go 1 2 3` n'affiche rien, et `go run main.go 5 2 x` affiche toujours son message d'erreur.

**Extension** : comptez les lignes pour `3 1 2`. Combien en faudrait-il au minimum ?

---

# Partie 4 — Les petits cas : 2, 3 et 5 nombres

## Une méthode par taille

Une seule méthode pour toutes les tailles, c'est simple, mais c'est du gaspillage sur les petites piles. Un vrai Push-Swap choisit la méthode selon le nombre d'éléments :

| Nombre d'éléments | Méthode |
|---|---|
| Déjà triée | Aucune opération |
| 2 | `sa` |
| 3 | `trierTrois`, de la séance 2 : 2 opérations au plus |
| 4 ou 5 | `trierCinq`, que vous allez écrire |
| Plus de 5 | La méthode générale, en attendant la séance 5 |

On range ce choix dans une fonction `trier` :

```go
func (p *Piles) trier() {
	if p.aEstTriee() {
		return
	}

	if len(p.a) == 2 {
		p.sa()
	} else if len(p.a) == 3 {
		p.trierTrois()
	} else if len(p.a) <= 5 {
		p.trierCinq()
	} else {
		p.trierOptimise()
	}
}
```

## Trier cinq nombres

L'idée : réduire le problème à un cas qu'on sait déjà résoudre parfaitement.

1. Mettre de côté les deux plus petits nombres sur b, chacun amené au sommet par le chemin le plus court.
2. Il reste trois nombres dans a : on les trie avec `trierTrois`.
3. Remettre les deux nombres de b sur a avec `pa`.

Pourquoi ça marche : b contient le plus petit au fond et le deuxième plus petit au sommet. Le premier `pa` pose le deuxième plus petit sur les trois nombres triés, le second `pa` pose le plus petit au-dessus de lui.

Avec 4 nombres, on ne met de côté qu'**un** nombre. Écrivez donc la boucle « tant que a contient plus de 3 éléments » : elle fonctionne pour 4 comme pour 5.

---

## Exercice 6

Ajoutez `trierTrois` de la séance 2 et la fonction `trier`. Dans `executer`, remplacez l'appel à votre méthode générale par `p.trier()`.

Pour l'instant, écrivez une première version de `trierCinq` qui appelle simplement votre méthode générale. Vous la remplacerez à l'exercice 7.

**Votre programme doit afficher** :

| Commande | Opérations affichées |
|---|---|
| `go run main.go 1 2 3` | rien |
| `go run main.go 2 1` | `sa` |
| `go run main.go 3 1 2` | `ra` |
| `go run main.go 3 2 1` | `sa` puis `rra` |

**Extension** : testez `trierTrois` avec trois nombres quelconques, par exemple `42 7 19`.

---

## Exercice 7

Écrivez la vraie fonction `trierCinq` avec la méthode de la partie 4.

**Avec `go run main.go 5 2 9 1 7`, votre programme doit afficher 10 lignes** au lieu de 14 :

```
rra
rra
pb
rra
rra
pb
sa
rra
pa
pa
```

**Vérification** : `go run main.go 4 3 2 1` affiche 5 lignes, et `go run main.go 1 2 3 4 5` n'affiche rien.

**Extension** : testez `2 1 3 4 5`. Combien d'opérations votre programme utilise-t-il ? Combien en faudrait-il au minimum ? Notez cette limite : vous y reviendrez.

**Bonus** : il existe 120 façons d'ordonner 5 nombres. Avec cette méthode, aucune ne demande plus de 12 opérations. Testez-en le plus possible.

---

# Partie 5 — Les rangs

## Rappel de la séance 3

Le **rang** d'un nombre est sa place dans l'ordre croissant : 0 pour le plus petit, n − 1 pour le plus grand.

| Nombres | 40 | 10 | 50 | 20 | 30 |
|---|---|---|---|---|---|
| Rangs | 3 | 0 | 4 | 1 | 2 |

## Pourquoi travailler avec les rangs

**Les rangs conservent l'ordre.** Si 40 est plus grand que 20, son rang est plus grand aussi. Trier les rangs donne donc exactement les mêmes opérations que trier les nombres.

**Les rangs sont simples.** Quels que soient les nombres reçus, négatifs ou immenses, les rangs vont toujours de 0 à n − 1. La stratégie par blocs de la séance 5 en a besoin : elle raisonne en « les 15 plus petits », « les 15 suivants »…

## Calculer un rang

Le rang d'un nombre, c'est **le nombre de valeurs plus petites que lui**. 40 a trois valeurs plus petites (10, 20, 30) : son rang est 3.

On compte donc, pour chaque nombre, combien d'autres sont plus petits. Deux boucles imbriquées : c'est un travail en O(n²), mais il n'est fait qu'une fois, et il ne produit aucun mouvement Push-Swap.

---

## Exercice 8

Écrivez une fonction `rangs(nombres []int) []int` qui renvoie la tranche des rangs.

Créez la tranche du résultat avec `make([]int, len(nombres))`, puis remplissez chaque case.

**Vérification** : `rangs([]int{40, 10, 50, 20, 30})` renvoie `[3 0 4 1 2]`.

Dans `executer`, remplacez `p := Piles{a: nombres}` par `p := Piles{a: rangs(nombres)}`.

**Vérification** : `go run main.go 40 10 50 20 30` et `go run main.go 3 0 4 1 2` affichent exactement les mêmes opérations.

**Extension** : que donne `rangs` pour `-5 1000 3` ?

---

# Synthèse

## Exercice 9

Vérifiez votre programme complet avec ce tableau. Lancez chaque commande et comparez.

| Commande | Résultat attendu |
|---|---|
| `go run main.go` | Rien, code 0 |
| `go run main.go 1 2 3` | Rien, code 0 |
| `go run main.go 2 1` | `sa` |
| `go run main.go 3 1 2` | `ra` |
| `go run main.go 5 2 9 1 7` | 10 lignes |
| `go run main.go 40 10 50 20 30` | 6 lignes |
| `go run main.go -3 10 0 7 -1 2` | 17 lignes |
| `go run main.go 5 2 x` | `Erreur : "x" n'est pas un nombre entier`, code 1 |
| `go run main.go 5 2 5` | `Erreur : le nombre 5 apparaît plusieurs fois`, code 1 |

**Vérification** : les neuf lignes correspondent.

**Extension** : avec 6 nombres et plus, votre programme utilise encore la méthode générale. Comptez les lignes pour cette pile de 25 nombres :

```
go run main.go 18 4 22 9 1 14 25 7 12 3 20 16 5 24 11 8 19 2 15 23 6 13 21 10 17
```

Notez ce chiffre : c'est lui que la séance 5 va faire baisser.

---

## Ce qu'il faut retenir

Un programme reçoit ses données de l'extérieur, et doit les vérifier avant de travailler.
Les arguments sont des textes : il faut les convertir, et refuser ce qui n'est pas convertible.
La sortie d'un Push-Swap ne contient que les opérations, une par ligne.
Un petit cas se ramène à un cas déjà résolu : cinq nombres, c'est trois nombres et deux mis de côté.
Les rangs remplacent les nombres sans changer l'ordre, et vont toujours de 0 à n − 1.

La séance suivante utilise ces rangs pour trier par blocs, en faisant enfin travailler la pile b.

---

# Erreurs fréquentes

| Erreur | Ce que vous voyez | Correction |
|---|---|---|
| Utiliser `os.Args` au lieu de `os.Args[1:]` | Le chemin du programme est traité comme un nombre | Les arguments commencent à l'index 1 |
| Une seule variable pour `strconv.Atoi` | `assignment mismatch: 1 variable but strconv.Atoi returns 2 values` | `nombre, err := strconv.Atoi(texte)` |
| Ignorer l'erreur avec `_` | `x` devient silencieusement 0 | Tester `err != nil` et renvoyer une erreur |
| Message d'erreur avec `fmt.Println` | L'erreur se mélange aux opérations | `fmt.Fprintln(os.Stderr, …)` |
| `os.Exit(1)` dans `executer` | Le programme s'arrête sans message | Renvoyer l'erreur ; seul `main` arrête le programme |
| Map déclarée sans `{}` | `panic: assignment to entry in nil map` | `vus := map[int]bool{}` |
| `vus[nombre] = true` avant le test | Chaque nombre est signalé comme doublon | Tester d'abord, noter ensuite |
| Affichage de la pile en plus des opérations | Le nombre de lignes est faux | N'afficher que les opérations |
| Boucle `for len(p.a) > 2` dans `trierCinq` | `trierTrois` reçoit 2 éléments et plante | `for len(p.a) > 3` |
| Rang calculé avec `<=` | Chaque rang est décalé de 1 | Compter les valeurs strictement plus petites |
| `rangs` qui modifie la tranche reçue | Les nombres d'origine sont perdus | Écrire dans une nouvelle tranche créée avec `make` |

---

# Annexe — Le code de la séance 2

Si vous ne l'avez plus, copiez ce code dans votre `main.go`, sous la ligne `import`. Il contient la structure, les huit opérations, et les fonctions utiles.

```go
type Piles struct {
	a          []int
	b          []int
	operations []string
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

func (p *Piles) pa() {
	if len(p.b) == 0 {
		return
	}
	p.a = append([]int{p.b[0]}, p.a...)
	p.b = p.b[1:]
	p.operations = append(p.operations, "pa")
}

func (p *Piles) pb() {
	if len(p.a) == 0 {
		return
	}
	p.b = append([]int{p.a[0]}, p.b...)
	p.a = p.a[1:]
	p.operations = append(p.operations, "pb")
}

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

func (p *Piles) rra() {
	if len(p.a) < 2 {
		return
	}
	dernier := p.a[len(p.a)-1]
	p.a = append([]int{dernier}, p.a[:len(p.a)-1]...)
	p.operations = append(p.operations, "rra")
}

func (p *Piles) rrb() {
	if len(p.b) < 2 {
		return
	}
	dernier := p.b[len(p.b)-1]
	p.b = append([]int{dernier}, p.b[:len(p.b)-1]...)
	p.operations = append(p.operations, "rrb")
}

// aEstTriee renvoie true si chaque élément de a est inférieur ou égal au suivant.
func (p *Piles) aEstTriee() bool {
	for i := 0; i < len(p.a)-1; i++ {
		if p.a[i] > p.a[i+1] {
			return false
		}
	}
	return true
}

// positionDuMinimum renvoie la position du plus petit élément de a.
func (p *Piles) positionDuMinimum() int {
	indexMin := 0
	for i := 1; i < len(p.a); i++ {
		if p.a[i] < p.a[indexMin] {
			indexMin = i
		}
	}
	return indexMin
}

// trierTrois trie une pile a de trois éléments en deux opérations au plus.
func (p *Piles) trierTrois() {
	haut := p.a[0]
	milieu := p.a[1]
	bas := p.a[2]

	if haut < milieu && milieu < bas {
		// 1 2 3 : déjà trié
	} else if haut > milieu && milieu < bas && haut < bas {
		p.sa()
	} else if haut > milieu && milieu < bas && haut > bas {
		p.ra()
	} else if haut < milieu && milieu > bas && haut > bas {
		p.rra()
	} else if haut < milieu && milieu > bas && haut < bas {
		p.sa()
		p.ra()
	} else {
		p.sa()
		p.rra()
	}
}
```

Votre bloc `import` doit contenir `"errors"`, `"fmt"`, `"os"` et `"strconv"`.
