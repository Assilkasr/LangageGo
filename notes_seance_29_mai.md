# Notes de séance — 29 mai 2026
### Cours : Langage Go

---

## Déclaration de variables

En Go, il existe deux façons de déclarer une variable :

### `var` — déclaration universelle
```go
var poids int = 70
var nom string = "Alice"
```
- Fonctionne **partout** : à l'intérieur comme à l'extérieur des fonctions
- À utiliser obligatoirement pour les déclarations **au niveau du package** (hors de toute fonction)

### `:=` — déclaration courte
```go
poids := 70
nom := "Alice"
```
- Fonctionne **uniquement à l'intérieur d'une fonction**
- Interdit en dehors d'une fonction (niveau package)

### Résumé : où déclarer quoi ?

| Endroit | `var` | `:=` |
|---|---|---|
| Niveau package (hors fonction) | oui | **non** |
| Dans `main()` ou toute autre fonction | oui | oui |

> `main` est bien une fonction — donc `:=` y est autorisé.

---

## Constantes

```go
const TVA = 0.20
```
- Se déclarent avec le mot-clé `const`
- À placer de préférence **en dehors des fonctions**, au niveau du package
- Ne peuvent pas être modifiées après leur déclaration

---

## Structure d'un fichier Go (rappel)

```go
package main

import "fmt"

const TVA = 0.20  // constante au niveau package

var pays string = "France"  // variable au niveau package

func main() {
    poids := 70       // := autorisé ici (dans une fonction)
    fmt.Println(poids)
}
```

---

## Switch / case

Go propose le `switch case` pour gérer plusieurs cas :

```go
switch note {
case "A":
    fmt.Println("Excellent")
case "B":
    fmt.Println("Bien")
default:
    fmt.Println("Autre")
}
```

> Il n'existe pas de mot-clé `match` en Go (contrairement à Rust par exemple).  
> Le `switch` en Go est l'équivalent — et il est plus souple : pas besoin de `break`, il ne tombe pas dans le cas suivant automatiquement.
