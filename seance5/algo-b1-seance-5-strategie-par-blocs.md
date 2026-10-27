# Séance 5 — La stratégie par blocs

**Algorithmie · Bachelor 1 · TDP 4 h · Ycode A2627_4659 · Go**

Ce document se suffit à lui-même. Lisez chaque partie, tapez le code, faites les exercices dans l'ordre. Chaque exercice indique ce que votre programme doit afficher.

| Temps | Contenu |
|---|---|
| 0 h 00 | Partie 0 — Le problème de la méthode actuelle |
| 0 h 15 | Partie 1 — Phase 1 : pousser vers b par blocs |
| 1 h 15 | Partie 2 — Phase 2 : remonter les plus grands d'abord |
| 1 h 55 | Pause |
| 2 h 10 | Partie 3 — Comparer les deux stratégies |
| 3 h 15 | Partie 4 — Intégrer au programme |
| 3 h 40 | Synthèse |

Vous partez du programme de la séance 4. Copiez votre dossier `seance4` en `seance5`, et travaillez dans la copie.

---

# Partie 0 — Le problème de la méthode actuelle

## Ce qu'elle fait

Votre méthode générale répète trois gestes : chercher le minimum de a, le faire remonter au sommet, le pousser sur b. La pile b ne sert qu'à **stocker** : on y empile, puis on remonte tout à la fin.

## Pourquoi elle coûte cher

À chaque tour, le minimum peut se trouver n'importe où dans a. Il faut parfois de nombreuses rotations pour l'amener au sommet, et on recommence n fois. En séance 3, vous avez mesuré le résultat : environ 1 450 mouvements pour 100 nombres, 32 000 pour 500.

## L'idée de la séance

Au lieu de chercher **le** plus petit, on accepte d'envoyer sur b **n'importe quel nombre d'un groupe de petits rangs**, dès qu'il passe au sommet. On ne cherche plus : on trie grossièrement en passant.

Puis on fait travailler la pile b : on y range les nombres pour qu'ils soient faciles à retrouver au moment de les remonter.

La stratégie se fait en deux phases :

| Phase | Ce qu'on fait |
|---|---|
| 1 | Vider a dans b, bloc de rangs par bloc de rangs |
| 2 | Remonter b dans a, du plus grand rang au plus petit |

Cette stratégie travaille sur les **rangs** calculés en séance 4 : les nombres vont de 0 à n − 1.

---

# Partie 1 — Phase 1 : pousser vers b par blocs

## La règle

On garde un compteur `pousses` : le nombre d'éléments déjà envoyés sur b. On choisit aussi une **taille de bloc**, par exemple 15.

Tant que a n'est pas vide, on regarde le rang au sommet de a :

| Si le rang au sommet est… | On fait | Pourquoi |
|---|---|---|
| inférieur ou égal à `pousses` | `pb` puis `rb` | C'est un des plus petits : on l'envoie **au fond** de b |
| inférieur ou égal à `pousses + tailleBloc` | `pb` | Il est dans le bloc en cours : on le pose **au sommet** de b |
| plus grand | `ra` | Trop grand pour l'instant : on le fait passer, il reviendra |

Après chaque `pb`, `pousses` augmente de 1. Le bloc **avance** donc au fur et à mesure : plus on pousse de nombres, plus les rangs acceptés deviennent grands.

## Ce que ça produit dans b

Les nombres partent sur b à peu près dans l'ordre de leur rang. Les plus petits de chaque moment vont au fond, les autres au sommet. Les derniers poussés, qui ont les plus grands rangs, se retrouvent donc près d'une **extrémité** de b : au sommet ou au fond, jamais perdus au milieu.

C'est ce qui rendra la phase 2 rapide.

## Pourquoi la boucle s'arrête toujours

Si `pousses` éléments ont été envoyés, au moins un des rangs de 0 à `pousses` est encore dans a. Ce rang est accepté dès qu'il passe au sommet. En un tour complet de a au plus, un élément est poussé : la pile a finit toujours par se vider.

---

## Exercice 1 — À la main

Déroulez la phase 1 **sur papier**, pour la pile de rangs `4 1 5 0 3 2`, avec une taille de bloc de **2**. Le sommet est à gauche.

Recopiez et complétez le tableau, une ligne par décision :

| `pousses` | Rang au sommet | Décision | Pile a après | Pile b après |
|---|---|---|---|---|
| 0 | 4 | 4 > 0 + 2 : `ra` | `1 5 0 3 2 4` | vide |
| 0 | 1 | | | |
| … | | | | |

Rappel : `rb` sur une pile b d'un seul élément ne fait rien.

**Vérification** : à la fin, a est vide, b vaut `3 1 0 2 4 5`, et vous avez écrit 12 opérations.

---

## Exercice 2 — Coder la phase 1

Écrivez la méthode :

```go
// pousserParBlocs vide a dans b, en ne poussant que les rangs
// du bloc en cours. Les plus petits du bloc partent au fond de b.
func (p *Piles) pousserParBlocs(tailleBloc int) {
	pousses := 0

	for len(p.a) > 0 {
		// à compléter avec la règle de la partie 1
	}
}
```

Pour la tester sans toucher au reste, ajoutez un **mode d'essai** : si le premier argument est `essai`, le programme lance une fonction de test au lieu de trier.

Au début d'`executer`, juste après le cas sans argument :

```go
	if args[0] == "essai" {
		essai()
		return nil
	}
```

Et la fonction d'essai :

```go
// essai déroule la phase 1 sur l'exemple de la partie 1.
func essai() {
	p := Piles{a: []int{4, 1, 5, 0, 3, 2}}
	p.pousserParBlocs(2)

	fmt.Println("a :", p.a, " b :", p.b)
	fmt.Println("opérations :", p.operations)
	fmt.Println("nombre d'opérations :", len(p.operations))
}
```

**Avec `go run main.go essai`, votre programme doit afficher** :

```
a : []  b : [3 1 0 2 4 5]
opérations : [ra pb ra pb rb pb pb rb pb rb pb rb]
nombre d'opérations : 12
```

Comparez avec votre tableau de l'exercice 1 : les opérations doivent être les mêmes.

**Extension** : relancez l'essai avec une taille de bloc de 0, puis de 10. Que devient b, et combien d'opérations faut-il ?

---

# Partie 2 — Phase 2 : remonter les plus grands d'abord

## La règle

Pour reconstruire a dans l'ordre croissant, on doit poser les nombres **du plus grand au plus petit** : le plus grand d'abord, qui finira au fond de a, et le plus petit en dernier, au sommet.

Tant que b n'est pas vide :

1. trouver la position du plus grand rang de b ;
2. l'amener au sommet de b par le chemin le plus court : `rb` ou `rrb` ;
3. le remonter sur a avec `pa`.

C'est exactement ce que vous faites déjà sur a avec `positionDuMinimum` et `amenerAuSommet`, transposé à b et au maximum.

## Pourquoi c'est rapide

Grâce à la phase 1, le plus grand rang restant est presque toujours près d'une extrémité de b. Avec le choix entre `rb` et `rrb`, il ne faut que peu de rotations pour l'atteindre.

Le programme cherche toujours le maximum à chaque tour : le travail **interne** reste important. Mais les **mouvements** Push-Swap, ceux qu'on compte, diminuent fortement. C'est la distinction de la séance 3 entre travail interne et opérations produites.

---

## Exercice 3 — Coder la phase 2

Écrivez trois méthodes :

- `positionDuMaximumB() int` : la position du plus grand élément de b ;
- `amenerAuSommetB(position int)` : comme `amenerAuSommet`, mais sur b, avec `rb` et `rrb` ;
- `remonterLesPlusGrands()` : la boucle de la partie 2.

Modifiez `essai` pour lancer les deux phases et afficher les piles entre les deux :

```go
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
```

**Avec `go run main.go essai`, votre programme doit afficher** :

```
après la phase 1 - a : []  b : [3 1 0 2 4 5]
après la phase 2 - a : [0 1 2 3 4 5]  b : []
opérations : [ra pb ra pb rb pb pb rb pb rb pb rb rrb pa rrb pa pa rrb pa pa pa]
nombre d'opérations : 21
```

**Extension** : avant de lancer le programme, déroulez la phase 2 à la main à partir de b = `3 1 0 2 4 5`. Pour chaque `pa`, notez la position du maximum et la rotation choisie.

---

## Exercice 4 — Assembler la stratégie

Écrivez la méthode qui enchaîne les deux phases :

```go
// trierParBlocs trie a en deux phases : pousser par blocs, puis remonter.
func (p *Piles) trierParBlocs(tailleBloc int) {
	p.pousserParBlocs(tailleBloc)
	p.remonterLesPlusGrands()
}
```

Puis une fonction qui choisit la taille de bloc. Prenez pour l'instant 15 jusqu'à 100 éléments, et 30 au-delà :

```go
// tailleDeBloc choisit la taille des blocs selon le nombre d'éléments.
func tailleDeBloc(n int) int {
	if n <= 100 {
		return 15
	}
	return 30
}
```

Dans `trier`, remplacez `p.trierOptimise()` par `p.trierParBlocs(tailleDeBloc(len(p.a)))`.

Vous pouvez maintenant retirer le mode `essai` et sa fonction.

**Avec la pile de 25 nombres de la séance 4, votre programme doit afficher 116 lignes**, au lieu de 165 :

```
go run main.go 18 4 22 9 1 14 25 7 12 3 20 16 5 24 11 8 19 2 15 23 6 13 21 10 17
```

**Vérification** : `go run main.go 5 2 9 1 7` affiche toujours 10 lignes. Les petits cas n'ont pas changé.

---

# Partie 3 — Comparer les deux stratégies

## Une comparaison équitable

La séance 3 l'a montré : une stratégie n'est pas meilleure parce qu'elle paraît plus intelligente. Il faut la **mesurer**, et la comparer à l'ancienne sur **exactement les mêmes piles**.

Pour chaque taille :

1. tirer une pile au hasard avec `rand.Perm(n)`, qui donne les nombres de 0 à n − 1 dans le désordre ;
2. en donner une **copie** à chaque stratégie ;
3. relever le nombre de mouvements de chacune ;
4. recommencer plusieurs fois, et faire la moyenne.

La moyenne est importante : une seule pile peut être particulièrement favorable à une méthode, par hasard.

## Le gain

Comme en séance 3 :

```
gain (%) = (mouvements anciens − mouvements nouveaux) × 100 ÷ mouvements anciens
```

Un gain négatif veut dire que la nouvelle stratégie fait **moins bien**.

---

## Exercice 5 — Le mode `comparer`

Ajoutez un deuxième mode : avec `go run main.go comparer`, le programme compare les deux stratégies au lieu de trier.

Au début d'`executer`, juste après le cas sans argument :

```go
	if args[0] == "comparer" {
		comparer()
		return nil
	}
```

Puis les deux fonctions. Ajoutez `"math/rand"` au bloc `import`.

```go
// copier renvoie une copie indépendante de la tranche.
func copier(nombres []int) []int {
	copie := make([]int, len(nombres))
	copy(copie, nombres)
	return copie
}

// comparer lance les deux stratégies sur les mêmes piles au hasard
// et affiche la moyenne des mouvements pour chaque taille.
func comparer() {
	tailles := []int{10, 20, 50, 100, 500}
	essais := 20

	for _, n := range tailles {
		totalSimple := 0
		totalBlocs := 0

		for e := 0; e < essais; e++ {
			depart := rand.Perm(n)

			// à compléter : lancer trierOptimise sur une copie,
			// trierParBlocs sur une autre copie,
			// et ajouter leurs nombres d'opérations aux totaux
		}

		moyenneSimple := totalSimple / essais
		moyenneBlocs := totalBlocs / essais
		gain := 100 * (moyenneSimple - moyenneBlocs) / moyenneSimple

		fmt.Printf("n=%-4d  séance 2 : %-6d  blocs : %-6d  gain : %d %%\n", n, moyenneSimple, moyenneBlocs, gain)
	}
}
```

**Votre programme doit afficher cinq lignes de ce genre** :

```
n=10    séance 2 : 31      blocs : 35      gain : -12 %
n=20    séance 2 : 91      blocs : 85      gain : 6 %
n=50    séance 2 : 427     blocs : 246     gain : 42 %
n=100   séance 2 : 1458    blocs : 577     gain : 60 %
n=500   séance 2 : 32506   blocs : 5083    gain : 84 %
```

Vos chiffres varient à chaque lancement, puisque les piles sont tirées au hasard. Les ordres de grandeur, eux, doivent être les mêmes.

**Vérification** : à partir de 50 éléments, la stratégie par blocs gagne nettement ; sur 500 éléments, elle utilise plus de 5 fois moins de mouvements.

**Extension** : regardez la ligne `n=10`. Que signifie un gain négatif ? Proposez une explication.

---

## Exercice 6 — Choisir la taille des blocs

La taille de bloc change le résultat. Pour la choisir, on mesure.

Modifiez temporairement `tailleDeBloc` pour qu'elle renvoie toujours la même valeur, par exemple `return 5`. Lancez `go run main.go comparer`, et relevez la colonne « blocs » pour 100 et 500 éléments. Recommencez avec 10, 15, 20, 30 et 40.

| Taille de bloc | 100 éléments | 500 éléments |
|---|---|---|
| 5 | | |
| 10 | | |
| 15 | | |
| 20 | | |
| 30 | | |
| 40 | | |

**Vérification** : pour 100 éléments, la meilleure taille se situe vers 15 ; pour 500 éléments, vers 30. Ce sont les valeurs de `tailleDeBloc`. Remettez-la dans son état de l'exercice 4.

**Extension** : expliquez en une phrase pourquoi un bloc trop petit coûte cher, puis pourquoi un bloc trop grand coûte cher aussi.

---

# Partie 4 — Intégrer au programme

## Le programme de fin de séance

Votre programme choisit maintenant sa méthode ainsi :

| Nombre d'éléments | Méthode |
|---|---|
| Déjà triée | Aucune opération |
| 2 | `sa` |
| 3 | `trierTrois` |
| 4 ou 5 | `trierCinq` |
| Plus de 5 | `trierParBlocs`, avec une taille de bloc de 15 ou 30 |

Il garde les vérifications de la séance 4 : arguments, doublons, résultat final.

---

# Synthèse

## Exercice 7

Vérifiez votre programme complet avec ce tableau :

| Commande | Résultat attendu |
|---|---|
| `go run main.go` | Rien, code 0 |
| `go run main.go 2 1` | `sa` |
| `go run main.go 3 1 2` | `ra` |
| `go run main.go 5 2 9 1 7` | 10 lignes |
| La pile de 25 nombres | 116 lignes |
| `go run main.go 5 2 x` | `Erreur : "x" n'est pas un nombre entier`, code 1 |
| `go run main.go 5 2 5` | `Erreur : le nombre 5 apparaît plusieurs fois`, code 1 |
| `go run main.go comparer` | Cinq lignes de comparaison |

**Vérification** : les huit lignes correspondent.

**Extension** : d'après l'exercice 5, la stratégie par blocs fait moins bien que la méthode générale sur les petites piles. Modifiez `trier` pour utiliser `trierOptimise` en dessous de 20 éléments, et `trierParBlocs` à partir de 20.

Pour mesurer l'effet, `comparer` doit tester votre méthode de choix, et non plus les blocs seuls : remplacez `blocs.trierParBlocs(tailleDeBloc(n))` par `blocs.trier()`. Relancez `comparer` : que devient la ligne `n=10` ?

**Bonus** : la règle « en dessous de 20 » vient d'une mesure. Testez 15 et 25 à la place de 20. Lequel est le meilleur choix ?

---

## Ce qu'il faut retenir

Une bonne stratégie fait travailler les deux piles : b n'est plus une simple réserve, elle est organisée pour faciliter le retour.
On n'a plus besoin du minimum exact : un nombre du bloc en cours suffit, et on le prend dès qu'il passe.
Le travail interne du programme et les mouvements produits sont deux coûts différents.
Une stratégie se compare à une autre sur les mêmes données, avec une moyenne sur plusieurs essais.
Une stratégie meilleure en moyenne peut être moins bonne sur certaines tailles : on choisit la méthode selon la taille.
Les réglages, comme la taille de bloc, se décident en mesurant.

---

# Erreurs fréquentes

| Erreur | Ce que vous voyez | Correction |
|---|---|---|
| Phase 1 sur les nombres au lieu des rangs | Rien n'est poussé, ou le programme ne s'arrête plus | Les piles doivent contenir `rangs(nombres)` |
| `pousses++` oublié après un `pb` | Le bloc n'avance pas, le programme tourne sans fin | Augmenter `pousses` après **chaque** `pb` |
| `pousses++` après un `ra` | Le bloc avance trop vite : le tri reste juste, mais coûte environ deux fois plus | Seuls les `pb` comptent |
| `rb` et `pb` inversés | Le mauvais nombre part au fond de b | D'abord `pb`, puis `rb` |
| `rb` au lieu de `rrb` dans `amenerAuSommetB` | Trop de rotations dans la phase 2 | Même règle que sur a : `rb` pour la première moitié, `rrb` sinon |
| Minimum cherché dans b au lieu du maximum | `Erreur : la pile n'a pas été triée` : a est rangée à l'envers | En phase 2, on remonte le plus **grand** |
| Les deux stratégies sur la même pile | La seconde reçoit une pile déjà triée et ne fait rien | Une copie pour chacune avec `copier` |
| Comparaison sur une seule pile | Les chiffres changent beaucoup d'un lancement à l'autre | Faire la moyenne sur plusieurs essais |
| `100 * (a - b) / a` écrit `(a - b) / a * 100` | Le gain vaut toujours 0 | Multiplier par 100 **avant** de diviser |
| `undefined: rand` | Le programme ne compile pas | Ajouter `"math/rand"` aux imports |
| Mode `essai` laissé dans le programme final | `go run main.go essai` fonctionne encore | Le retirer une fois la stratégie terminée |
