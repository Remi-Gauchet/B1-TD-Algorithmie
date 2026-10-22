# Séance 2 — Corrigé des exercices

**Algorithmie · Bachelor 1 · Ycode A2627_4659 · Go**

Pour chaque exercice : le programme complet, une explication pas à pas, et les erreurs les plus fréquentes.

Chaque programme est le contenu **complet** du fichier `main.go` de votre dossier `seance2`. Remplacez tout le fichier, puis lancez :

```
go run main.go
```

Rappel des règles : le **sommet** d'une pile est la case 0. On ne peut toucher qu'au sommet.

---

# Exercice 1 — Le simulateur de départ

## Le programme

```go
package main

import "fmt"

type Piles struct {
	a          []int
	b          []int
	operations []string
}

func (p *Piles) ra() {
	if len(p.a) < 2 {
		return
	}
	p.a = append(p.a[1:], p.a[0])
	p.operations = append(p.operations, "ra")
}

func (p *Piles) sa() {
	if len(p.a) < 2 {
		return
	}
	p.a[0], p.a[1] = p.a[1], p.a[0]
	p.operations = append(p.operations, "sa")
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

func (p *Piles) afficher() {
	fmt.Println("a :", p.a, " b :", p.b)
}

func main() {
	p := Piles{a: []int{5, 2, 9, 1, 7}}
	p.afficher()
	p.pb()
	p.afficher()
	p.ra()
	p.afficher()
	p.pa()
	p.afficher()
	fmt.Println("opérations :", p.operations)
}
```

## Résultat

```
a : [5 2 9 1 7]  b : []
a : [2 9 1 7]  b : [5]
a : [9 1 7 2]  b : [5]
a : [5 9 1 7 2]  b : []
opérations : [pb ra pa]
```

## Pas à pas

**La structure `Piles` regroupe trois choses** : la pile a, la pile b, et la liste des opérations déjà faites. Comme une fiche à trois cases.

**`p := Piles{a: []int{5, 2, 9, 1, 7}}`** remplit seulement la case `a`. Les cases `b` et `operations` restent vides.

**`func (p *Piles) ra()`** est une fonction attachée à la structure. On l'appelle avec un point : `p.ra()`. On appelle ce genre de fonction une **méthode**.

**L'étoile dans `*Piles` est indispensable.** Elle veut dire : « travaille sur la vraie structure, pas sur une copie ». Sans elle, chaque opération modifierait une copie qui disparaît aussitôt, et les piles ne bougeraient jamais.

## Chaque opération, ligne par ligne

**`pb`** — prendre le sommet de a et le poser sur b.

```go
p.b = append([]int{p.a[0]}, p.b...)
p.a = p.a[1:]
```

La première ligne fabrique une nouvelle pile b : le sommet de a, suivi de tout l'ancien b. Les trois points `p.b...` veulent dire « tous les éléments de `p.b`, un par un ». La deuxième ligne retire le sommet de a : `p.a[1:]` garde tout à partir de la case 1.

**`ra`** — le sommet de a passe tout en dessous.

```go
p.a = append(p.a[1:], p.a[0])
```

On prend tout sauf le sommet, et on ajoute l'ancien sommet à la fin. `[2 9 1 7 5]` après un `ra` sur `[5 2 9 1 7]`.

**`pa`** fait comme `pb`, dans l'autre sens. **`sa`** échange les cases 0 et 1, comme les échanges de la séance 1.

**Le test de longueur au début de chaque opération** protège le programme. Faire `pb` sur une pile a vide, ou `sa` avec un seul élément, est impossible : l'opération ne fait rien, au lieu de planter.

**L'enregistrement se fait dans l'opération elle-même**, juste après l'action. Impossible d'oublier de noter une opération, ou d'en noter une qui n'a pas eu lieu : une opération refusée par le test de longueur s'arrête avant d'être enregistrée.

## Erreurs courantes

**L'étoile oubliée : `func (p Piles) ra()`.** Le programme tourne sans erreur, mais les quatre lignes affichent la même pile, et l'historique reste vide.

**`p.b = append(p.b, p.a[0])` dans `pb`.** L'élément est ajouté à la **fin** de b, c'est-à-dire tout en dessous, au lieu d'arriver au sommet.

**Les lignes de `pb` inversées.** Si on retire d'abord le sommet de a, `p.a[0]` désigne déjà l'élément suivant : c'est le mauvais nombre qui part sur b.

---

# Exercice 2 — Les quatre opérations manquantes

## Le programme

```go
package main

import "fmt"

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

func (p *Piles) afficher() {
	fmt.Println("a :", p.a, " b :", p.b)
}

func main() {
	p := Piles{a: []int{1, 2, 3}}
	p.ra()
	p.afficher()
	p.rra()
	p.afficher()

	q := Piles{a: []int{1, 2, 3}}
	q.ra()
	q.ra()
	q.ra()
	q.afficher()
	fmt.Println("opérations :", q.operations)
}
```

## Résultat

```
a : [2 3 1]  b : []
a : [1 2 3]  b : []
a : [1 2 3]  b : []
opérations : [ra ra ra]
```

## Pas à pas

**`rra`** — le dernier élément de a revient au sommet.

```go
dernier := p.a[len(p.a)-1]
p.a = append([]int{dernier}, p.a[:len(p.a)-1]...)
```

`p.a[len(p.a)-1]` est le dernier élément. `p.a[:len(p.a)-1]` est tout sauf le dernier. On fabrique une nouvelle pile : le dernier, suivi de tout le reste.

**`sb`, `rb` et `rrb`** sont des copies exactes de `sa`, `ra` et `rra`, avec `p.b` à la place de `p.a`, et le bon nom dans l'historique.

**Les deux tests de l'énoncé :**

| Test | Ce qu'on vérifie |
|---|---|
| `ra` puis `rra` sur `[1 2 3]` | `ra` donne `[2 3 1]`, `rra` ramène `[1 2 3]` : les deux opérations s'annulent |
| `ra` trois fois sur `[1 2 3]` | Chaque élément passe une fois en dessous : on retrouve `[1 2 3]` |

Ces tests sont précieux : si `rra` était écrit à l'envers, `ra` puis `rra` ne ramènerait pas la pile de départ.

## Erreurs courantes

**`rra` qui fait la même chose que `ra`.** Relisez le sens : `ra` envoie le sommet en dessous, `rra` remonte le dessous au sommet.

**`rb` qui enregistre `"ra"`.** La copie a été faite sans changer le nom dans l'historique.

**`sb` qui modifie `p.a`.** Un `a` oublié pendant la copie : vérifiez chaque ligne.

---

# Exercice 3 — Savoir si la pile est triée

## Le programme

```go
package main

import "fmt"

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

func (p *Piles) afficher() {
	fmt.Println("a :", p.a, " b :", p.b)
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

func main() {
	p1 := Piles{a: []int{1, 2, 3}}
	p2 := Piles{a: []int{3, 1, 2}}
	p3 := Piles{a: []int{}}

	fmt.Println(p1.aEstTriee())
	fmt.Println(p2.aEstTriee())
	fmt.Println(p3.aEstTriee())
}
```

## Résultat

```
true
false
true
```

## Pas à pas

**On compare chaque élément avec son voisin du dessous.** Dès qu'un élément est plus grand que le suivant, la pile n'est pas triée : on renvoie `false` immédiatement.

**Si la boucle arrive au bout sans rien trouver**, tous les voisins sont dans le bon ordre : on renvoie `true`.

**La boucle s'arrête à `len(p.a)-1`**, pour la même raison que dans le tri à bulles : on lit `p.a[i+1]`.

**Une pile vide est triée.** La boucle ne fait aucun tour, et la fonction renvoie `true`. C'est logique : il n'y a aucun élément mal placé.

## Extension — la fonction `resoudre`

```go
package main

import "fmt"

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

func (p *Piles) afficher() {
	fmt.Println("a :", p.a, " b :", p.b)
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

// resoudre affiche les opérations si la pile a est triée et b vide,
// et un message d'erreur sinon.
func (p *Piles) resoudre() {
	if p.aEstTriee() && len(p.b) == 0 {
		fmt.Println("opérations :", p.operations)
	} else {
		fmt.Println("Erreur : la pile n'est pas triée.")
	}
}

func main() {
	p := Piles{a: []int{3, 1, 2}}
	p.ra()
	p.resoudre()

	q := Piles{a: []int{3, 1, 2}}
	q.sa()
	q.resoudre()
}
```

```
opérations : [ra]
Erreur : la pile n'est pas triée.
```

La fonction vérifie deux choses : a est triée, **et** b est vide. Une pile a triée alors que des nombres sont encore sur b n'est pas une solution.

## Erreurs courantes

**`p.a[i] >= p.a[i+1]`.** Avec deux nombres égaux côte à côte, la pile est déclarée non triée, ce qui est faux.

**`return true` à l'intérieur de la boucle.** La fonction s'arrête après avoir regardé une seule paire.

---

# Exercice 4 — Vérifier le tableau des trois nombres

## Le programme

```go
package main

import "fmt"

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

func (p *Piles) afficher() {
	fmt.Println("a :", p.a, " b :", p.b)
}

func main() {
	fmt.Println("Cas 1 3 2 :")
	p := Piles{a: []int{1, 3, 2}}
	p.afficher()
	p.sa()
	p.afficher()
	p.ra()
	p.afficher()

	fmt.Println("Cas 2 3 1 :")
	q := Piles{a: []int{2, 3, 1}}
	q.afficher()
	q.rra()
	q.afficher()

	fmt.Println("Cas 3 2 1 :")
	r := Piles{a: []int{3, 2, 1}}
	r.afficher()
	r.sa()
	r.afficher()
	r.rra()
	r.afficher()
}
```

## Résultat

```
Cas 1 3 2 :
a : [1 3 2]  b : []
a : [3 1 2]  b : []
a : [1 2 3]  b : []
Cas 2 3 1 :
a : [2 3 1]  b : []
a : [1 2 3]  b : []
Cas 3 2 1 :
a : [3 2 1]  b : []
a : [2 3 1]  b : []
a : [1 2 3]  b : []
```

## Pas à pas

Les trois déroulés aboutissent à `[1 2 3]`. Regardez le cas `1 3 2` : on ne peut pas échanger directement le 3 et le 2, qui ne sont pas au sommet. On fait d'abord `sa` pour obtenir `3 1 2`, puis `ra` envoie le 3 en dessous.

C'est la contrainte du jeu : on ne peut agir qu'au sommet, alors on prépare la pile pour que le bon mouvement devienne possible.

---

# Exercice 5 — Trier trois nombres

## Le programme

```go
package main

import "fmt"

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

func (p *Piles) afficher() {
	fmt.Println("a :", p.a, " b :", p.b)
}

// trierTrois trie une pile a de trois éléments en deux opérations au plus.
func (p *Piles) trierTrois() {
	if len(p.a) != 3 {
		return
	}

	haut := p.a[0]
	milieu := p.a[1]
	bas := p.a[2]

	if haut < milieu && milieu < bas {
		// 1 2 3 : déjà trié
	} else if haut > milieu && milieu < bas && haut < bas {
		// 2 1 3
		p.sa()
	} else if haut > milieu && milieu < bas && haut > bas {
		// 3 1 2
		p.ra()
	} else if haut < milieu && milieu > bas && haut > bas {
		// 2 3 1
		p.rra()
	} else if haut < milieu && milieu > bas && haut < bas {
		// 1 3 2
		p.sa()
		p.ra()
	} else {
		// 3 2 1
		p.sa()
		p.rra()
	}
}

func main() {
	cas := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
		{42, 7, 19},
	}

	for _, depart := range cas {
		p := Piles{a: []int{depart[0], depart[1], depart[2]}}
		p.trierTrois()
		fmt.Println(depart, "->", p.a, "en", len(p.operations), "opération(s) :", p.operations)
	}
}
```

## Résultat

```
[1 2 3] -> [1 2 3] en 0 opération(s) : []
[1 3 2] -> [1 2 3] en 2 opération(s) : [sa ra]
[2 1 3] -> [1 2 3] en 1 opération(s) : [sa]
[2 3 1] -> [1 2 3] en 1 opération(s) : [rra]
[3 1 2] -> [1 2 3] en 1 opération(s) : [ra]
[3 2 1] -> [1 2 3] en 2 opération(s) : [sa rra]
[42 7 19] -> [7 19 42] en 1 opération(s) : [ra]
```

## Pas à pas

**On donne un nom aux trois cases** : `haut`, `milieu` et `bas`. Le code se lit plus facilement que `p.a[0]`, `p.a[1]` et `p.a[2]`.

**Chaque cas se reconnaît par des comparaisons.** Pour `2 1 3`, par exemple : le haut est plus grand que le milieu, le milieu est plus petit que le bas, et le haut est plus petit que le bas.

| Cas | haut > milieu ? | milieu > bas ? | haut > bas ? | Opérations |
|---|---|---|---|---|
| `1 2 3` | non | non | — | aucune |
| `2 1 3` | oui | non | non | `sa` |
| `3 1 2` | oui | non | oui | `ra` |
| `2 3 1` | non | oui | oui | `rra` |
| `1 3 2` | non | oui | non | `sa` puis `ra` |
| `3 2 1` | oui | oui | — | `sa` puis `rra` |

**Le dernier cas est traité par le `else`.** Après avoir écarté les cinq autres, il ne reste que `3 2 1`.

**Extension — des nombres quelconques.** Le code ne compare que des positions relatives : « plus grand que », « plus petit que ». Il ne cherche jamais la valeur 1, 2 ou 3. Il fonctionne donc avec `{42, 7, 19}`, reconnu comme le cas `3 1 2`.

**La fonction suppose trois nombres différents**, comme dans le vrai Push-Swap, qui interdit les doublons.

## Erreurs courantes

**Comparer à des valeurs fixes : `if p.a[0] == 2`.** Le code marche avec 1, 2 et 3, mais pas avec `{42, 7, 19}`.

**Lire `p.a[0]` après une opération.** Si vous comparez les cases au fur et à mesure, elles ont déjà changé. On les range dans `haut`, `milieu` et `bas` **avant** toute opération.

**Deux conditions qui se recouvrent.** Si un cas correspond à deux branches, seule la première s'exécute : testez les six arrangements pour vérifier que chacun tombe dans la bonne.

---

# Exercice 6 — La position du minimum

## Le programme

```go
package main

import "fmt"

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

func (p *Piles) afficher() {
	fmt.Println("a :", p.a, " b :", p.b)
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

func main() {
	p := Piles{a: []int{5, 2, 9, 1, 7}}
	fmt.Println(p.positionDuMinimum())

	q := Piles{a: []int{1, 2, 3}}
	fmt.Println(q.positionDuMinimum())
}
```

## Résultat

```
3
0
```

## Pas à pas

C'est l'exercice 2 de la séance 1, appliqué à la pile a. On retient une **position**, qui commence à 0, et on la remplace dès qu'on trouve plus petit.

La boucle peut commencer à 1 : la case 0 est déjà la candidate de départ, inutile de la comparer à elle-même.

---

# Exercice 7 — La méthode générale

## Le programme

```go
package main

import "fmt"

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

func (p *Piles) afficher() {
	fmt.Println("a :", p.a, " b :", p.b)
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

// trierSimple pousse le minimum de a sur b, tour après tour,
// en le faisant remonter uniquement avec ra. Puis il remonte tout sur a.
func (p *Piles) trierSimple() {
	for !p.aEstTriee() {
		position := p.positionDuMinimum()
		for i := 0; i < position; i++ {
			p.ra()
		}
		p.pb()
	}

	for len(p.b) > 0 {
		p.pa()
	}
}

func main() {
	p := Piles{a: []int{5, 2, 9, 1, 7}}
	p.trierSimple()

	p.afficher()
	fmt.Println("opérations :", p.operations)
	fmt.Println("nombre d'opérations :", len(p.operations))
}
```

## Résultat

```
a : [1 2 5 7 9]  b : []
opérations : [ra ra ra pb ra ra pb ra ra pb ra pb pa pa pa pa]
nombre d'opérations : 16
```

## Pas à pas

La méthode répète trois gestes tant que a n'est pas triée : trouver le minimum, le faire remonter au sommet avec des `ra`, le pousser sur b.

| Étape | Pile a avant | Minimum | Opérations | a après | b après |
|---|---|---|---|---|---|
| 1 | `5 2 9 1 7` | 1, position 3 | `ra ra ra pb` | `7 5 2 9` | `1` |
| 2 | `7 5 2 9` | 2, position 2 | `ra ra pb` | `9 7 5` | `2 1` |
| 3 | `9 7 5` | 5, position 2 | `ra ra pb` | `9 7` | `5 2 1` |
| 4 | `9 7` | 7, position 1 | `ra pb` | `9` | `7 5 2 1` |
| 5 | `9` | — | a est triée, on s'arrête | `9` | `7 5 2 1` |
| 6 | | | `pa pa pa pa` | `1 2 5 7 9` | vide |

**Pour faire remonter un élément en position `k`, il faut `k` rotations `ra`.** D'où la boucle `for i := 0; i < position; i++`.

**La boucle s'arrête dès que a est triée**, sans attendre qu'elle soit vide. Ici, il ne reste que le 9 : une pile d'un seul élément est toujours triée. Ce qui reste dans a est forcément plus grand que tout ce qui est sur b, puisqu'on a poussé les plus petits d'abord.

**b est rangée du plus grand au sommet au plus petit au fond** : `7 5 2 1`. En remontant b avec des `pa`, chaque nombre se pose sur un plus grand que lui. a se reconstruit triée.

**16 opérations pour 5 nombres.** C'est beaucoup : la moitié, 8 sur 16, sont des rotations. C'est ce coût qu'on va réduire.

## Bonus — dix nombres au hasard

Vous trouverez en moyenne une quarantaine d'opérations, parfois moins de 20, parfois plus de 50. Le total change à chaque tirage, parce que le nombre de rotations dépend de l'endroit où se trouvent les minimums.

## Erreurs courantes

**Le programme ne s'arrête plus.** La boucle `for !p.aEstTriee()` tourne sans fin si aucune opération ne fait progresser la pile, par exemple si `pb` n'est jamais appelé. Arrêtez avec `Ctrl + C`, puis vérifiez que chaque tour pousse bien un élément sur b.

**Il reste des nombres sur b.** La boucle finale des `pa` a été oubliée.

**`for i := 0; i <= position; i++`.** Une rotation de trop : le minimum dépasse le sommet et repart en dessous.

---

# Exercice 8 — Choisir le sens de rotation

## Le programme

```go
package main

import (
	"fmt"
	"math/rand"
)

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

func (p *Piles) afficher() {
	fmt.Println("a :", p.a, " b :", p.b)
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

// trierSimple pousse le minimum de a sur b, tour après tour,
// en le faisant remonter uniquement avec ra. Puis il remonte tout sur a.
func (p *Piles) trierSimple() {
	for !p.aEstTriee() {
		position := p.positionDuMinimum()
		for i := 0; i < position; i++ {
			p.ra()
		}
		p.pb()
	}

	for len(p.b) > 0 {
		p.pa()
	}
}

// trierOptimise fait la même chose que trierSimple, mais choisit
// le sens de rotation le plus court pour amener le minimum au sommet.
func (p *Piles) trierOptimise() {
	for !p.aEstTriee() {
		position := p.positionDuMinimum()

		if position < len(p.a)/2 {
			for i := 0; i < position; i++ {
				p.ra()
			}
		} else {
			for i := 0; i < len(p.a)-position; i++ {
				p.rra()
			}
		}

		p.pb()
	}

	for len(p.b) > 0 {
		p.pa()
	}
}

// pileAleatoire renvoie les nombres de 0 à n-1 dans un ordre au hasard.
func pileAleatoire(n int) []int {
	return rand.Perm(n)
}

// copier renvoie une copie indépendante de la tranche.
func copier(nombres []int) []int {
	copie := make([]int, len(nombres))
	copy(copie, nombres)
	return copie
}

func main() {
	p := Piles{a: []int{5, 2, 9, 1, 7}}
	p.trierOptimise()
	p.afficher()
	fmt.Println("opérations :", p.operations)
	fmt.Println("nombre d'opérations :", len(p.operations))
	fmt.Println()

	tailles := []int{5, 10, 20, 50}
	for _, n := range tailles {
		depart := pileAleatoire(n)

		simple := Piles{a: copier(depart)}
		simple.trierSimple()

		optimise := Piles{a: copier(depart)}
		optimise.trierOptimise()

		gain := 100 * (len(simple.operations) - len(optimise.operations)) / len(simple.operations)
		fmt.Printf("taille %-3d  ra seule : %-5d  ra + rra : %-5d  gain : %d %%\n",
			n, len(simple.operations), len(optimise.operations), gain)
	}
}
```

## Résultat

```
a : [1 2 5 7 9]  b : []
opérations : [rra rra pb rra rra pb rra pb rra pb pa pa pa pa]
nombre d'opérations : 14

taille 5    ra seule : 14     ra + rra : 13     gain : 7 %
taille 10   ra seule : 48     ra + rra : 31     gain : 35 %
taille 20   ra seule : 129    ra + rra : 92     gain : 28 %
taille 50   ra seule : 670    ra + rra : 405    gain : 39 %
```

Le tableau change à chaque lancement, puisque les piles sont tirées au hasard.

## Pas à pas

**Deux chemins pour amener la position `k` au sommet d'une pile de `n` éléments :**
- `k` rotations `ra`, qui font monter par le haut ;
- `n − k` rotations `rra`, qui font remonter par le bas.

**La règle : si `k` est dans la première moitié, `ra` ; sinon, `rra`.** On choisit le chemin le plus court.

Sur `[5 2 9 1 7]`, le 1 est en position 3, sur une pile de 5. Avec `ra`, il faut 3 rotations ; avec `rra`, seulement 5 − 3 = 2. D'où les deux `rra` du début.

**On compare les deux méthodes sur les mêmes piles.** `copier` donne à chacune sa propre copie de la pile de départ, comme en séance 1.

**`rand.Perm(n)`** renvoie les nombres de 0 à n − 1 dans un ordre au hasard. Chaque nombre apparaît une seule fois : il n'y a pas de doublon, comme dans le vrai Push-Swap.

**Le gain en pour cent** se calcule ainsi : (ancien total − nouveau total) × 100 ÷ ancien total.

## Bonus — le gain sur une pile de 50

Environ **40 %** d'opérations en moins. Le résultat est identique, seul le coût change : c'est une optimisation.

Le gain est plus faible sur les petites piles : avec 5 éléments, il y a peu de rotations à économiser.

## Erreurs courantes

**`for i := 0; i < position; i++` gardé pour les `rra`.** Le nombre de `rra` n'est pas `position`, mais `len(p.a) - position`.

**`len(p.a)` calculé après avoir commencé les rotations.** Ici, ce n'est pas un problème : les rotations ne changent pas la taille de la pile. Mais un `pb` placé dans la boucle de rotation, oui.

**Division entière inattendue.** `len(simple.operations) / len(optimise.operations) * 100` donne 100 ou 0, car Go divise des entiers sans virgule. Multipliez par 100 **avant** de diviser.

---

# Exercice 9 — Synthèse

## Le programme

```go
package main

import (
	"fmt"
	"math/rand"
)

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

func (p *Piles) afficher() {
	fmt.Println("a :", p.a, " b :", p.b)
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

// trierOptimise fait la même chose que trierSimple, mais choisit
// le sens de rotation le plus court pour amener le minimum au sommet.
func (p *Piles) trierOptimise() {
	for !p.aEstTriee() {
		position := p.positionDuMinimum()

		if position < len(p.a)/2 {
			for i := 0; i < position; i++ {
				p.ra()
			}
		} else {
			for i := 0; i < len(p.a)-position; i++ {
				p.rra()
			}
		}

		p.pb()
	}

	for len(p.b) > 0 {
		p.pa()
	}
}

// pileAleatoire renvoie les nombres de 0 à n-1 dans un ordre au hasard.
func pileAleatoire(n int) []int {
	return rand.Perm(n)
}

func main() {
	taille := 20

	p := Piles{a: pileAleatoire(taille)}
	fmt.Println("départ :", p.a)

	p.trierOptimise()

	if p.aEstTriee() && len(p.b) == 0 {
		fmt.Println("pile triée :", p.a)
		fmt.Println("nombre d'opérations :", len(p.operations))
	} else {
		fmt.Println("Erreur : la pile n'est pas triée.")
	}
}
```

## Exemple de résultat

```
départ : [2 11 7 14 4 8 9 15 10 1 17 16 3 13 18 5 19 0 12 6]
pile triée : [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
nombre d'opérations : 93
```

Lancé dix fois, le programme affiche à chaque fois une pile triée, avec un nombre d'opérations différent selon le tirage.

## Pas à pas

**Le programme ne fait pas confiance à l'algorithme : il vérifie.** Après le tri, il contrôle que a est triée et que b est vide avant d'annoncer le résultat. Si un jour vous modifiez `trierOptimise` et introduisez une erreur, le message vous le dira immédiatement.

**On ne garde que la meilleure méthode**, `trierOptimise`. La version avec `ra` seule a servi de point de comparaison.

**Pour changer la taille**, modifiez la variable `taille`.

## Erreurs courantes

**Afficher seulement la pile triée.** Le résultat attendu d'un Push-Swap est la **liste des opérations**, et le chiffre à surveiller est leur nombre.

**Vérifier seulement `aEstTriee`.** Une pile a triée avec des nombres encore sur b n'est pas une solution : il faut aussi `len(p.b) == 0`.

---

# Récapitulatif des points difficiles

**Le sommet d'une pile est la case 0.** Un élément poussé arrive en position 0.

**L'étoile dans `func (p *Piles)`.** Sans elle, les opérations ne modifient qu'une copie.

**Chaque opération vérifie la taille de la pile avant d'agir**, et s'enregistre elle-même dans l'historique.

**`ra` envoie le sommet en dessous, `rra` remonte le dessous au sommet.**

**Trois nombres se trient en deux opérations au plus**, et sans la pile b.

**On part d'une méthode qui marche, on la mesure, puis on l'améliore.** Le choix entre `ra` et `rra` réduit le coût d'environ 40 % sur les grandes piles.

**Le résultat attendu est la liste des opérations**, pas la pile triée.
