package main

import "fmt"
import "math"
import "time"
import "errors"

func operer(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("erreur : division par zéro impossible")
		}
		return a / b, nil
	default:
		return 0, errors.New("erreur : opérateur inconnu")
	}
}
func creerOperation(op string) func(float64, float64) (float64, error) {
	return func(a float64, b float64) (float64, error) {
		return operer(a, b, op)
	}
}
func main() {
	maintenant := time.Now()
    fmt.Println("Langage installé : Hello, World!")
	fmt.Println("Le carré de 64 est :", math.Pow(64,2))
	fmt.Println("la date du jour est :", maintenant.Format("02/01/2006"))
	dateNaissance := time.Date(2002, time.December, 03, 0, 0, 0, 0, time.Local)
	age := maintenant.Year() - dateNaissance.Year()
	if maintenant.Month() < dateNaissance.Month() || (maintenant.Month() == dateNaissance.Month() && maintenant.Day() < dateNaissance.Day()) {
		age--
	}
	fmt.Printf("L'âge de la personne née le %s est : %d ans\n", dateNaissance.Format("02/01/2006"), age)



	// EXERCICE NOTE 1 : Déclaration de la constante Nom avec ton prénom
	const Nom = "TonPrénom"

	var poids float64 = 70.5
	var taille float64 = 1.75

	const IMCMaigreur = 18.5
	const IMCNormal = 25.0
	const IMCSurpoids = 30.0

	imc := poids / math.Pow(taille, 2)
	fmt.Printf("Bonjour %s, voici les résultats de ton bilan :\n", Nom)

	fmt.Printf("Ton IMC est de : %.2f\n", imc)

	fmt.Print("Catégorie : ")
	if imc < IMCMaigreur {
		fmt.Println("Maigreur")
	} else if imc < IMCNormal {
		fmt.Println("Normal")
	} else if imc < IMCSurpoids {
		fmt.Println("Surpoids")
	} else {
		fmt.Println("Obésité")
	}


	// EXERCICE NOTE 2 : Calculatrice avec closures et validation
	fmt.Println("Entrez vos calculs sous la forme : [nombre1] [nombre2] [opérateur]")
	fmt.Println("Exemple : 10 5 +")
	fmt.Println("Tapez '0 0 quit' pour sortir.")

	for {
		var a, b float64
		var op string

		print("> ")
		_, err := fmt.Scan(&a, &b, &op)
		if err != nil {
			fmt.Println("Erreur de saisie, réessayez.")
			var abandon string
			fmt.Scanln(&abandon)
			continue
		}

		if op == "quit" {
			fmt.Println("Fin du programme. Au revoir !")
			break
		}

		calculateur := creerOperation(op)
		resultat, errCalcul := calculateur(a, b)

		if errCalcul != nil {
			fmt.Println(errCalcul)
		} else {
			fmt.Printf("Résultat : %.2f\n", resultat)
		}
	}


}
