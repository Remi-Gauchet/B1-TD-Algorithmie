# Séance 2 — Push-Swap : trier avec des contraintes

**Algorithmie · Bachelor 1 · TDP 4 h · Ycode A2627_4659 · Go**

Ce document se suffit à lui-même. Tout se fait sur machine. Chaque exercice indique ce que votre programme doit afficher.

| Temps | Contenu |
|---|---|
| 0 h 00 | Partie 1 — Les deux piles et les opérations |
| 0 h 50 | Partie 2 — Construire le simulateur |
| 1 h 40 | Pause |
| 1 h 55 | Partie 3 — Trier trois nombres |
| 2 h 35 | Partie 4 — Une méthode générale |
| 3 h 30 | Partie 5 — Mesurer et améliorer |

---

# Partie 1 — Les règles du jeu

## Le problème

Vous avez deux piles, nommées **a** et **b**. Au départ, la pile a contient des nombres dans le désordre, la pile b est vide.

Objectif : la pile a triée dans l'ordre croissant, la pile b vide, en utilisant **le moins d'opérations possible**.

Vous n'avez pas accès à un élément au milieu d'une pile. Seul le sommet est accessible. C'est toute la difficulté.

## Les opérations autorisées

| Opération | Effet |
|---|---|
| `sa` | Échange les deux premiers éléments de a |
| `sb` | Échange les deux premiers éléments de b |
| `pa` | Prend le premier de b et le pose sur a |
| `pb` | Prend le premier de a et le pose sur b |
| `ra` | Décale a vers le haut : le premier devient le dernier |
| `rb` | Même chose sur b |
| `rra` | Décale a vers le bas : le dernier devient le premier |
| `rrb` | Même chose sur b |

Rien d'autre. Pas d'accès à une position choisie, pas d'échange entre deux éléments éloignés.

## Ce qui change par rapport à la séance 1

En séance 1, vous pouviez échanger n'importe quelles cases. Ici non. Un algorithme efficace sur une tranche libre peut devenir mauvais avec ces règles.

C'est le point central du module : un algorithme n'est pas bon dans l'absolu, il est bon **pour un jeu d'opérations donné**.

## Ce que le programme doit produire

Attention à ce détail, il surprend tout le monde : le résultat attendu n'est pas la pile triée. C'est **la liste des opérations** qui permet de la trier. C'est elle qu'on affiche, et c'est son nombre de lignes qu'on cherche à réduire.

---

# Partie 2 — Construire le simulateur

## Représenter les piles

On regroupe les deux piles et l'historique des opérations dans une seule structure. Le sommet de chaque pile est la case 0.

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

Trois choses à repérer.

`func (p *Piles) ra()` signifie que `ra` est une opération **de** la structure. On l'appelle avec `p.ra()`. L'étoile indique qu'elle modifie la structure et non une copie — sans elle, vos opérations n'auraient aucun effet.

Le test de longueur au début de chaque opération évite de planter sur une pile vide. Une opération impossible ne fait rien.

L'opération s'enregistre dans `operations` au moment même où elle est exécutée. C'est ce qui garantit que l'historique est exact.

---

## Exercice 1

Recopiez le programme et lancez-le.

**Votre programme doit afficher** :

```
a : [5 2 9 1 7]  b : []
a : [2 9 1 7]  b : [5]
a : [9 1 7 2]  b : [5]
a : [5 9 1 7 2]  b : []
opérations : [pb ra pa]
```

Si une ligne diffère, relisez l'opération concernée avant de continuer. Tout le reste de la séance repose sur ce simulateur.

---

## Exercice 2

Ajoutez les quatre opérations manquantes : `rra`, `sb`, `rb`, `rrb`.

Pour `rra` : le dernier élément revient devant. Le dernier élément d'une tranche s'écrit `p.a[len(p.a)-1]`, et tout sauf le dernier s'écrit `p.a[:len(p.a)-1]`.

**Test à écrire vous-même** : partez de `a : [1 2 3]`, appliquez `ra` puis `rra`. Vous devez retrouver `[1 2 3]`. Les deux opérations s'annulent.

**Extension** : partez de `a : [1 2 3]` et appliquez `ra` trois fois. Vous devez retrouver `[1 2 3]`.

---

## Exercice 3

Écrivez une fonction qui dit si la pile a est triée :

```go
func (p *Piles) aEstTriee() bool {
	// renvoie true si chaque élément est inférieur ou égal au suivant
}
```

**Tests à écrire vous-même** : `[1 2 3]` donne `true`, `[3 1 2]` donne `false`, une pile vide donne `true`.

**Extension** : écrivez une fonction `resoudre()` qui affiche la liste des opérations seulement si la pile est triée à la fin, et un message d'erreur sinon.

---

# Partie 3 — Trier trois nombres

## Pourquoi commencer si petit

Trois nombres, c'est six arrangements possibles. On peut tous les énumérer et trouver la meilleure solution pour chacun. C'est le seul cas où « optimal » se démontre à la main.

Et cette solution servira de brique pour les cas plus grands.

## Les six cas

Les nombres sont 1, 2 et 3. Le sommet est à gauche.

| État de départ | Solution | Nombre d'opérations |
|---|---|---|
| `1 2 3` | rien | 0 |
| `1 3 2` | `sa` puis `ra` | 2 |
| `2 1 3` | `sa` | 1 |
| `2 3 1` | `rra` | 1 |
| `3 1 2` | `ra` | 1 |
| `3 2 1` | `sa` puis `rra` | 2 |

Aucun cas ne demande plus de deux opérations. Et aucun n'utilise la pile b : pour trois éléments, elle est inutile.

---

## Exercice 4

Vérifiez trois lignes du tableau avec votre simulateur, en affichant l'état après chaque opération.

**Vérification** : chaque déroulé aboutit à `a : [1 2 3]`.

---

## Exercice 5

Écrivez une fonction `trierTrois()` qui trie une pile a de trois éléments, quel que soit l'arrangement de départ.

M�thode : comparez `p.a[0]`, `p.a[1]` et `p.a[2]` pour identifier dans quel cas vous êtes, puis appliquez les opérations du tableau.

**Vérification** : testez votre fonction sur les six arrangements. Les six doivent aboutir à `[1 2 3]`, et le nombre d'opérations doit correspondre à la dernière colonne du tableau.

**Extension** : faites fonctionner votre fonction avec trois nombres quelconques, par exemple `{42, 7, 19}`, et pas seulement avec 1, 2 et 3.

---

# Partie 4 — Une méthode générale

## Le principe

Tant que la pile a n'est pas triée : cherche le plus petit élément de a, fais tourner a avec `ra` jusqu'à l'amener au sommet, puis pousse-le sur b avec `pb`. Quand a est vide, remonte tout b sur a avec `pa`.

Pourquoi ça marche : on pousse les éléments sur b du plus petit au plus grand. Donc b finit rangée du plus grand en haut au plus petit en bas. En la remontant, a se remplit dans le bon ordre.

Cette méthode fonctionne toujours. Elle est aussi très coûteuse : chaque rotation compte, et on en fait beaucoup. C'est volontaire — on part d'une solution qui marche, on mesure son coût, puis on l'améliore.

---

## Exercice 6

Écrivez une fonction `positionDuMinimum()` qui renvoie la position du plus petit élément de la pile a.

**Vérification** : sur `[5 2 9 1 7]`, elle renvoie `3`. Sur `[1 2 3]`, elle renvoie `0`.

---

## Exercice 7

Écrivez la méthode générale décrite plus haut, en utilisant votre fonction précédente.

**Vérification** : sur `{5, 2, 9, 1, 7}`, la pile a vaut `[1 2 5 7 9]` à la fin, la pile b est vide, et votre programme affiche la liste complète des opérations.

**Extension** : affichez le nombre total d'opérations utilisées.

**Bonus** : testez sur une pile de dix nombres aléatoires. Notez le total.

---

# Partie 5 — Mesurer et améliorer

## La première optimisation

Pour amener un élément au sommet, vous utilisez `ra`. Mais si cet élément se trouve dans la seconde moitié de la pile, `rra` l'atteint plus vite.

Exemple sur une pile de dix éléments : amener celui qui est en position 8 demande huit `ra`, ou seulement deux `rra`.

La règle : si la position est inférieure à la moitié de la taille, utilisez `ra`. Sinon, `rra`.

Cette amélioration ne change pas le résultat, seulement le nombre d'opérations. C'est la première vraie décision d'optimisation du module.

---

## Exercice 8

Modifiez votre méthode générale pour qu'elle choisisse entre `ra` et `rra` selon la position du minimum.

**Vérification** : le résultat reste correct, et le nombre total d'opérations diminue.

**Extension** : remplissez ce tableau en lançant les deux versions sur les mêmes piles aléatoires.

| Taille | Version `ra` seule | Version `ra` + `rra` |
|---|---|---|
| 5 | | |
| 10 | | |
| 20 | | |
| 50 | | |

**Bonus** : de combien de pour cent avez-vous réduit le nombre d'opérations sur une pile de 50 ?

---

## Exercice 9 — Synthèse

Écrivez un programme qui génère une pile aléatoire de taille donnée, applique votre meilleure méthode, vérifie que le résultat est trié, et affiche le nombre d'opérations.

**Vérification** : lancé dix fois de suite, votre programme annonce toujours une pile triée, avec un nombre d'opérations qui varie selon le tirage.

---

# Synthèse

Les opérations disponibles font partie du problème, pas du décor.
Un petit cas se résout de façon optimale par énumération ; un grand cas demande une stratégie.
On part d'une méthode qui fonctionne, on la mesure, puis on l'améliore — jamais l'inverse.
Le livrable n'est pas la pile triée, c'est la suite d'opérations qui y mène.

Le cours de complexité qui suit donne les outils pour comparer ces méthodes autrement qu'en comptant à la main.

---

# Erreurs fréquentes

| Erreur | Ce que vous voyez | Correction |
|---|---|---|
| Oublier l'étoile dans `func (p *Piles)` | Les opérations n'ont aucun effet | Ajouter `*` : la fonction doit modifier la structure |
| Confondre `ra` et `rra` | La pile tourne dans le mauvais sens | `ra` : le premier va à la fin. `rra` : le dernier revient devant |
| Confondre `pa` et `pb` | Les éléments partent du mauvais côté | `pb` prend dans a et pose sur b |
| Opération sur une pile vide | `index out of range` | Tester la longueur au début de chaque opération |
| Pousser par le mauvais bout | La pile se comporte comme une file d'attente | Un élément poussé arrive **au sommet**, donc en position 0 |
| Oublier d'enregistrer l'opération | L'historique est incomplet | Enregistrer dans la même fonction que l'action |
| Chercher à trier la pile b | Travail inutile | b ne sert que de stockage temporaire |
| Boucle `tant que a n'est pas triée` sans progression | Le programme ne s'arrête plus | `Ctrl + C`, puis vérifier qu'une opération est bien effectuée à chaque tour |
| Afficher la pile au lieu des opérations | Le livrable est faux | Le résultat attendu est la liste des opérations |
