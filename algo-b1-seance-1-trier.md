# Séance 1 — Trier : deux méthodes, deux coûts

**Algorithmie · Bachelor 1 · TDP 4 h · Ycode A2627_4659 · Go**

Ce document se suffit à lui-même. Lisez chaque partie, tapez le code, faites les exercices dans l'ordre. Chaque exercice indique exactement ce que votre programme doit afficher : vous savez seul si vous avez réussi.

| Temps | Contenu |
|---|---|
| 0 h 00 | Partie 0 — Mise en route |
| 0 h 20 | Partie 1 — Le tri par sélection |
| 1 h 20 | Partie 2 — Observer le déroulé |
| 1 h 50 | Pause |
| 2 h 05 | Partie 3 — Le tri à bulles |
| 3 h 00 | Partie 4 — Compter et comparer |
| 3 h 40 | Synthèse |

---

# Partie 0 — Mise en route

## Créer et lancer un programme

Créez un dossier `seance1`, et dedans un fichier `main.go` :

```go
package main

import "fmt"

func main() {
	fmt.Println("Prêt")
}
```

Dans le terminal, placez-vous dans le dossier et lancez :

```
go run main.go
```

**Vous devez voir** : `Prêt`

Si la commande échoue, vous n'êtes pas dans le bon dossier. Utilisez `cd` pour vous y rendre.

## Les trois outils dont vous aurez besoin

**Une tranche** est une suite de valeurs numérotées à partir de zéro.

```go
nombres := []int{5, 2, 9, 1, 7}
fmt.Println(nombres[0])    // 5
fmt.Println(len(nombres))  // 5
```

`nombres[0]` est la première case, `nombres[4]` la dernière. Demander `nombres[5]` fait planter le programme : il n'y a que cinq cases, numérotées de 0 à 4.

**Une boucle** répète une instruction en faisant varier un compteur.

```go
for i := 0; i < len(nombres); i++ {
	fmt.Println(i, nombres[i])
}
```

`i := 0` est la valeur de départ, `i < len(nombres)` la condition pour continuer, `i++` l'avancement à chaque tour.

**Un échange** se fait en une ligne en Go.

```go
nombres[0], nombres[1] = nombres[1], nombres[0]
```

Les deux affectations ont lieu en même temps. Vous n'avez pas besoin de variable intermédiaire.

---

## Exercice 1

Écrivez un programme qui affiche chaque nombre de la tranche `{5, 2, 9, 1, 7}` précédé de sa position.

**Votre programme doit afficher** :

```
0 5
1 2
2 9
3 1
4 7
```

**Extension** : affichez-les dans l'ordre inverse, de la position 4 à la position 0.

---

## Exercice 2

Écrivez un programme qui trouve le plus petit nombre de la tranche `{5, 2, 9, 1, 7}` et affiche sa valeur **et** sa position.

M�thode : créez une variable `indexMin` valant `0` avant la boucle. Parcourez la tranche ; à chaque fois que vous trouvez un nombre plus petit que `nombres[indexMin]`, remplacez `indexMin` par la position courante.

**Votre programme doit afficher** :

```
minimum 1 en position 3
```

**Extension** : faites la même chose pour le maximum.

**Bonus** : affichez le minimum, le maximum, et leur écart.

---

# Partie 1 — Le tri par sélection

## Le principe

Cherche le plus petit nombre. Place-le en première position. Recommence avec le reste, en ignorant ce qui est déjà placé.

C'est exactement l'exercice 2, répété : d'abord sur toute la tranche, puis sur tout sauf la première case, puis sur tout sauf les deux premières, et ainsi de suite.

## Le déroulé sur `5 2 9 1 7`

| Tour | Zone encore à trier | Minimum | Après l'échange |
|---|---|---|---|
| 1 | `5 2 9 1 7` | 1 | `1 2 9 5 7` |
| 2 | `2 9 5 7` | 2 | `1 2 9 5 7` |
| 3 | `9 5 7` | 5 | `1 2 5 9 7` |
| 4 | `9 7` | 7 | `1 2 5 7 9` |

Au tour 2, le minimum est déjà à sa place : on échange le nombre avec lui-même, ce qui ne change rien. L'algorithme ne pouvait pas le savoir à l'avance.

## Le code

```go
package main

import "fmt"

func triSelection(nombres []int) {
	for i := 0; i < len(nombres); i++ {
		indexMin := i
		for j := i + 1; j < len(nombres); j++ {
			if nombres[j] < nombres[indexMin] {
				indexMin = j
			}
		}
		nombres[i], nombres[indexMin] = nombres[indexMin], nombres[i]
	}
}

func main() {
	nombres := []int{5, 2, 9, 1, 7}
	triSelection(nombres)
	fmt.Println(nombres)
}
```

Trois points à comprendre avant de continuer.

La boucle sur `i` est le numéro du tour. Elle désigne la position que l'on va remplir.

La boucle sur `j` cherche le minimum. Elle part de `i + 1` et non de `0`, parce que tout ce qui se trouve avant la position `i` est déjà trié et définitif.

`indexMin` retient **où** se trouve le plus petit, pas sa valeur. C'est la position dont on a besoin pour faire l'échange.

---

## Exercice 3

Recopiez le programme et lancez-le.

**Votre programme doit afficher** :

```
[1 2 5 7 9]
```

**Extension** : testez avec `{1, 2, 3, 4, 5}` puis avec `{9, 7, 5, 2, 1}`. Le résultat est correct dans les deux cas.

---

## Exercice 4

Modifiez `triSelection` pour qu'elle affiche l'état de la tranche après chaque tour.

Placez un `fmt.Println(nombres)` à la fin de la boucle sur `i`, pas à l'intérieur de celle sur `j`.

**Votre programme doit afficher** :

```
[1 2 9 5 7]
[1 2 9 5 7]
[1 2 5 9 7]
[1 2 5 7 9]
[1 2 5 7 9]
```

Comparez ces lignes avec la colonne de droite du tableau de déroulé.

**Extension** : affichez aussi quel nombre a été sélectionné à chaque tour.

---

## Exercice 5

Faites compter à votre fonction le nombre de comparaisons et le nombre d'échanges.

Une comparaison a lieu à chaque passage dans la boucle sur `j`. Un échange a lieu une fois par tour de la boucle sur `i`.

Faites renvoyer les deux totaux par la fonction :

```go
func triSelection(nombres []int) (int, int) {
	comparaisons := 0
	echanges := 0
	// ... votre code
	return comparaisons, echanges
}
```

et récupérez-les dans `main` :

```go
comparaisons, echanges := triSelection(nombres)
fmt.Println("comparaisons :", comparaisons)
fmt.Println("échanges :", echanges)
```

**Sur `{5, 2, 9, 1, 7}`, vous devez obtenir** : 10 comparaisons et 5 échanges.

**Extension** : relancez avec une tranche déjà triée. Les chiffres changent-ils ?

**Bonus** : n'effectuez l'échange que si `indexMin` est différent de `i`. Combien d'échanges économisez-vous sur `{1, 2, 3, 4, 5}` ?

---

# Partie 2 — Observer le déroulé

## Pourquoi cette partie

Un algorithme qui donne le bon résultat peut quand même être mal compris. Afficher les étapes intermédiaires est la première technique de débogage, et vous l'utiliserez toute l'année.

---

## Exercice 6

Écrivez une fonction `afficherEtat` qui prend la tranche et le numéro du tour, et affiche une ligne lisible :

```
tour 1 : [1 2 9 5 7]
```

Appelez-la depuis `triSelection`.

**Votre programme doit afficher cinq lignes**, numérotées de 1 à 5, dont la dernière est `tour 5 : [1 2 5 7 9]`.

**Extension** : marquez d'un symbole la partie déjà triée, par exemple `tour 3 : [1 2 5 | 9 7]`.

---

# Partie 3 — Le tri à bulles

## Le principe

Parcours la tranche du début à la fin. À chaque étape, compare deux cases voisines et échange-les si elles sont dans le mauvais ordre. Recommence le parcours entier tant qu'au moins un échange a eu lieu.

Le nom vient de l'image : les grandes valeurs remontent vers la fin comme des bulles.

## La différence avec le tri par sélection

Le tri par sélection **cherche** le minimum, puis fait un seul échange par tour.
Le tri à bulles ne cherche rien : il compare uniquement des voisins, et échange beaucoup.

Les deux donnent le bon résultat. Ils n'ont pas le même coût — c'est ce que vous allez mesurer.

## Le squelette

```go
func triBulles(nombres []int) {
	for {
		echangeFait := false
		for i := 0; i < len(nombres)-1; i++ {
			if nombres[i] > nombres[i+1] {
				nombres[i], nombres[i+1] = nombres[i+1], nombres[i]
				echangeFait = true
			}
		}
		if !echangeFait {
			break
		}
	}
}
```

`for { }` sans condition répète indéfiniment. C'est le `break` qui arrête la boucle, quand un passage complet n'a produit aucun échange.

Remarquez `len(nombres)-1` : on compare la case `i` avec la case `i+1`, donc on doit s'arrêter une case avant la fin. Sans le `-1`, le programme plante.

---

## Exercice 7

Recopiez `triBulles` et faites-la fonctionner sur `{5, 2, 9, 1, 7}`.

**Votre programme doit afficher** :

```
[1 2 5 7 9]
```

**Extension** : affichez l'état de la tranche à la fin de chaque passage. Combien de passages sont nécessaires ?

---

## Exercice 8

Ajoutez les deux compteurs à `triBulles`, comme vous l'avez fait pour le tri par sélection.

**Sur `{5, 2, 9, 1, 7}`, vous devez obtenir** : 16 comparaisons et 5 échanges.

**Sur `{4, 1, 3, 2}`, vous devez obtenir** : 9 comparaisons et 4 échanges.

**Extension** : lancez les deux tris sur la même tranche et affichez les quatre chiffres côte à côte.

---

# Partie 4 — Compter et comparer

## Pourquoi compter plutôt que chronométrer

Le temps d'exécution dépend de la machine, de ce qui tourne en arrière-plan, du moment de la journée. Le nombre d'opérations ne dépend de rien d'autre que de l'algorithme et de ses données.

C'est pour ça qu'on compare des algorithmes en comptant des opérations, et c'est la base du cours de complexité qui suivra.

## Générer une tranche de test

Pour tester sur des tailles différentes, il vous faut des tranches plus grandes :

```go
import "math/rand"

func tranche(n int) []int {
	resultat := make([]int, n)
	for i := 0; i < n; i++ {
		resultat[i] = rand.Intn(100)
	}
	return resultat
}
```

`make([]int, n)` crée une tranche de `n` cases remplies de zéros. `rand.Intn(100)` tire un nombre entre 0 et 99.

Attention : les deux tris modifient la tranche qu'on leur donne. Pour comparer équitablement, chaque tri doit recevoir sa propre copie :

```go
original := tranche(20)

copie1 := make([]int, len(original))
copy(copie1, original)

copie2 := make([]int, len(original))
copy(copie2, original)
```

---

## Exercice 9

Remplissez ce tableau en lançant vos deux tris sur des tranches aléatoires de tailles croissantes :

| n | Sélection : comparaisons | Bulles : comparaisons | Bulles : échanges |
|---|---|---|---|
| 5 | | | |
| 10 | | | |
| 20 | | | |
| 50 | | | |

**Votre colonne « Sélection » doit valoir exactement** : 10, 45, 190, 1225. Ces valeurs ne dépendent pas du contenu de la tranche — vérifiez-le en relançant plusieurs fois.

**Extension** : les chiffres du tri à bulles changent à chaque exécution. Pourquoi ?

**Bonus** : quand `n` double, par combien le nombre de comparaisons est-il multiplié ? Répondez en une phrase.

---

## Exercice 10 — Synthèse

Écrivez un programme unique qui, pour `n` valant 5, 10, 20 et 50, lance les deux tris sur la même tranche et affiche une ligne par taille :

```
n=5   selection: 10 comp / 5 ech   bulles: 16 comp / 5 ech
```

**Vérification** : votre programme affiche quatre lignes, et les valeurs de la colonne sélection correspondent à celles de l'exercice 9.

---

# Synthèse

Un même problème admet plusieurs algorithmes corrects.
Le bon algorithme n'est pas celui qui donne le bon résultat, c'est celui qui le donne au meilleur coût.
Le coût se mesure en opérations, pas en secondes.
Le tri par sélection fait toujours le même nombre de comparaisons ; le tri à bulles dépend des données.

La séance suivante applique tout ça à un problème où les opérations disponibles sont imposées.

---

# Erreurs fréquentes

| Erreur | Ce que vous voyez | Correction |
|---|---|---|
| `go run main.go` échoue | `no such file` | Vous n'êtes pas dans le bon dossier, utilisez `cd` |
| Variable déclarée et non utilisée | `declared and not used` | Go refuse les variables inutiles, supprimez-la |
| Import manquant ou inutile | `undefined: fmt` ou `imported and not used` | Ajustez le bloc `import` |
| Boucle `j` partant de `0` | Le tri marche mais compte le double | Partir de `i + 1` |
| Retenir la valeur du minimum au lieu de sa position | Impossible de faire l'échange | Stocker l'index |
| `indexMin` déclaré avant la boucle `i` | Résultat faux | Le déclarer à chaque tour |
| Oublier le `-1` dans le tri à bulles | `index out of range` | `i < len(nombres)-1` |
| `Println` placé dans la mauvaise boucle | Trop de lignes affichées | Vérifier l'indentation et la position de l'appel |
| Les deux tris sur la même tranche | Le second reçoit une tranche déjà triée | Faire une copie avec `copy` |
| Compteur déclaré dans la boucle | Total toujours à 0 ou 1 | Le déclarer avant |
