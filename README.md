# Algorithmie B1
Exercices d'algorithmie de première année de bachelor, répartis en plusieurs séances.  
Un énoncé par séance sous la forme .md  

## Objectif
Introduction du concept de Push Swap.  

Ces séances nous font créer un programme qui trie une liste de *n* arguments.  
Pour cela, il répartit les valeurs sous la forme de deux colonnes, qu'il réorganise à travers plusieurs opérations, qu'on pourrait apparenter à un Rubik's Cube.  

## Les opérations autorisées
|    Opération    |              Effet           |  Exemple (sommet de la pile à gauche)   | Résultat |
|----------------|------------------------------|------------|----------|
|     `sa`     | Échange les deux premiers éléments de *a* | a=[1 2 3 4 5]  |   a=[2 1 3 4 5]   |
|     `sb`    | Échange les deux premiers éléments de *b* | b=[1 2 3 4 5] |   b=[2 1 3 4 5]  |
|     `pa`    | Prend le premier de *b* et le pose sur *a* | a=[1 2 3] b=[4 5] |   a=[4 1 2 3] b=[5] |
|    `pb` | Prend le premier de *a* et le pose sur *b* | a=[1 2 3] b=[4 5] | a=[2 3] b=[1 4 5] |
|   `ra` | Décale *a* vers le haut : le premier devient le dernier | a=[1 2 3] | a=[2 3 1] |
| `rb`   | Décale *b* vers le haut : le premier devient le dernier | b=[1 2 3] | b=[2 3 1] |
| `rra` | Décale *a* vers le bas : le dernier devient le premier | a=[1 2 3] | a=[3 1 2] |
| `rrb`| Décale *b* vers le bas : le dernier devient le premier | b=[1 2 3] | b=[3 1 2] |


## Mesurer l'optimisation
Pour comparer la méthode de tri par selection avec la méthode de tri par bloc, tapez :
```bash
go run . comparer
```
