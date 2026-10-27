# Algorithmie B1
Exercices d'algorithmie de première année de bachelor, répartis en plusieurs séances.  
Un énoncé par séance sous la forme .md  

## Objectif
Introduction du concept de Push Swap.  

Ces séances nous font créer un programme qui trie une liste de *n* arguments.  
Pour cela, il répartit les valeurs sous la forme de deux colonnes, qu'il réorganise à la manière d'un rubik's Cube.  

## Test
Pour tester le programme en prenant pour exemple la pile suivante : [4, 1, 5, 0, 3, 2]  
Tapez :
```bash
go run . essai
```
Votre programme doit afficher  :
```bash
a : []  b : [3 1 0 2 4 5]
opérations : [ra pb ra pb rb pb pb rb pb rb pb rb]
nombre d'opérations : 12
```
