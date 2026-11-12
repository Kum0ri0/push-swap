# push-swap

Trier une pile d'entiers avec le moins d'instructions possible, en n'utilisant qu'une seconde pile et un jeu de 11 instructions.

Le projet contient deux programmes écrits en Go, avec uniquement la bibliothèque standard :

- **`push-swap`** calcule et affiche une liste d'instructions qui trie la pile `a` (le plus petit nombre au sommet).
- **`checker`** lit des instructions sur l'entrée standard, les exécute, puis affiche `OK` si `a` est triée et `b` vide, sinon `KO`.

## Compilation

Depuis la racine du projet :

```bash
go build -o push-swap ./cmd/push-swap
go build -o checker ./cmd/checker
```

Sous Windows, ajouter l'extension `.exe` aux noms de sortie (`push-swap.exe`, `checker.exe`) et lancer les commandes dans Git Bash.

## Utilisation

```bash
$ ./push-swap "2 1 3"
sa

$ ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"
OK

$ echo -e "rra\npb\nsa\nrra\npa" | ./checker "3 2 1 0"
OK

$ ./push-swap "0 one 2 3"
Error
```

Les nombres peuvent être donnés dans une seule chaîne (`"3 1 2"`) ou comme arguments séparés (`3 1 2`). Le premier nombre est le sommet de la pile.

## Gestion des erreurs

`Error` est écrit sur la **sortie d'erreur standard** dans les cas suivants :

- un argument qui n'est pas un entier (`"1 a 3"`, `"2.5"`) ;
- un nombre hors des limites d'un entier 32 bits (`2147483648`) ;
- un doublon (`"1 2 1"`, `"0 -0"`) ;
- pour `checker` : une instruction inexistante ou mal formatée (`"SA"`, `"sa "`).

Sans argument, les deux programmes n'affichent rien.

## Organisation du projet

```
cmd/
├── checker/       programme checker : lit et exécute les instructions, affiche OK / KO
└── push-swap/     programme push-swap : trie la pile et affiche les instructions
internal/
├── parse/         lecture et validation des arguments (entiers, limites, doublons)
└── stack/         les piles, les 11 instructions, le journal des coups et les algorithmes de tri
```

Chaque responsabilité est séparée : le parsing ne connaît pas les piles, les piles ne lisent ni n'affichent rien, et les deux programmes ne font qu'assembler ces briques.

### Choix de représentation

Une pile est une slice d'entiers dont le **sommet est le dernier élément**. Ce choix rend `pa` et `pb` très simples (retirer et ajouter en fin de slice). En contrepartie, les nombres sont rangés en ordre inverse à la création de la pile.

Toutes les instructions passent par une seule méthode, `Executer`, qui exécute l'instruction et l'ajoute au **journal des coups**. C'est ce journal que `push-swap` affiche.

## Algorithme

Le choix dépend de la taille de la pile :

| Situation | Méthode |
|---|---|
| Pile déjà triée | aucune instruction |
| Pile triée « à une rotation près » (ex. `3 4 5 1 2`) | amener le plus petit au sommet, par le sens de rotation le plus court |
| 2 nombres | un `sa` si nécessaire |
| 3 nombres | mettre le plus grand au fond (`ra` ou `rra`), puis `sa` si besoin : 2 coups au maximum |
| 4 et 5 nombres | envoyer les plus petits dans `b`, trier les 3 restants, les ramener |
| 6 nombres et plus | tri par base (radix sort) en binaire |

### Tri par base (radix sort)

C'est un tri **non comparatif**, comme le suggère le sujet.

1. **Normalisation** : chaque nombre est remplacé par son rang (le nombre de valeurs plus petites que lui). `4 67 3 87 23` devient `1 3 0 4 2`. L'ordre est conservé, et on obtient des entiers positifs et petits.
2. **Une passe par bit**, du bit le moins significatif au plus significatif. Pour chaque nombre de `a` : si son bit vaut 0, `pb` ; sinon, `ra`. Puis tous les nombres de `b` reviennent avec `pa`.

Chaque passe range les nombres selon un bit sans casser l'ordre des passes précédentes. Après la dernière passe, la pile est triée.

### Optimisations

- **Pile triée à une rotation près** : détectée en comptant les « descentes » en faisant le tour de la pile (une seule descente = triée en rond).
- **Arrêt anticipé** : le tri par base s'arrête dès que la pile est triée.
- **Nettoyage du journal** : les coups qui s'annulent (`pa` puis `pb`, `ra` puis `rra`, etc.) sont supprimés avant l'affichage. Sur 100 nombres, cela économise environ 7 % des coups.

### Nombre de coups (mesures indicatives)

| Taille | Coups |
|---|---|
| 3 nombres | 2 au maximum |
| 5 nombres | une dizaine au maximum |
| 100 nombres | environ 1 000 |
| 500 nombres | environ 6 700 |

## Tests

```bash
go test ./...
```

Les tests unitaires couvrent :

- les piles et l'opération `Retirer` (y compris sur pile vide) ;
- l'exemple pas à pas du sujet ;
- le refus des instructions inconnues ou mal formatées ;
- le tri des cas limites (une valeur, déjà triée, triée en rond, nombres négatifs, limites 32 bits) ;
- 20 listes de 100 nombres aléatoires, en rejouant les coups affichés comme le ferait `checker` ;
- le nettoyage du journal ;
- la lecture des arguments valides et invalides.

Tests de bout en bout avec 100 nombres aléatoires :

```bash
ARG=$(shuf -i 1-1000 -n 100 | tr '\n' ' ')
./push-swap "$ARG" | ./checker "$ARG"
./push-swap "$ARG" | wc -l
```