package main

import "fmt"
/*EXO:
Crée une slice vide de nombres appelée carburant avec make (longueur 0, capacité 5).

Ajoute dedans les valeurs 45.5, 20.0 et 15.5 avec append.

Crée une slice valises ("Sac_Thomas", "Pack_Scientifique") et fusionne-la dans une autre slice fretOfficiel ("Satellite") avec l'opérateur ....

Découpe la slice carburant pour extraire uniquement les deux premiers éléments avec la syntaxe [:2] */

const (
	NiveauInvite = iota
	NiveauCodeur        
	NiveauAdmin         
)

func exoPerso() {
	autorisations := make([]int, 0, 3)

	autorisations = append(autorisations, NiveauInvite)
	autorisations = append(autorisations, NiveauCodeur, NiveauAdmin)

	fmt.Println("Liste des niveaux d'accès actifs :", autorisations) 

	accesStandard := autorisations[:2] 
	fmt.Println("Niveaux d'accès standards (Admin exclu) :", accesStandard) 
}