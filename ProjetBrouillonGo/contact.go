package main

import (
	"fmt"
)

type Personne struct {
	Prenom string
	Nom    string
	Age    int
	Email  string
}

func (p Personne) NomCompleto() string {
	return p.Prenom + " " + p.Nom
}

func (p Personne) Presentation() string {
	return fmt.Sprintf("Bonjour, je m'appelle %s et j'ai %d ans.", p.NomCompleto(), p.Age)
}

type Adresse struct {
	Rue        string
	Ville      string
	CodePostal string
}

func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

type Employe struct {
	Personne 
	Adresse  
	Poste    string
	Salaire  float64
}

func (e Employe) FicheEmploye() string {
	return fmt.Sprintf("EMPLOYÉ: %s\nPoste: %s | Salaire: %.2f€\nAdresse: %s\nContact: %s",
		e.NomCompleto(), e.Poste, e.Salaire, e.Format(), e.Email)
}

func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire = e.Salaire * (1 + pct/100)
}

type Etudiant struct {
	Personne
	Promo   string
	Moyenne float64
}

func (et Etudiant) MentionObtenue() string {
	switch {
	case et.Moyenne >= 16:
		return "Très Bien (TB)"
	case et.Moyenne >= 14:
		return "Bien (B)"
	case et.Moyenne >= 12:
		return "Assez Bien (AB)"
	case et.Moyenne >= 10:
		return "Passable (P)"
	default:
		return "Insuffisant"
	}
}

func contact() {
	fmt.Println("=== SYSTÈME DE CONTACTS ===")

	adr1 := Adresse{Rue: "10 rue de la Paix", Ville: "Paris", CodePostal: "75002"}
	adr2 := Adresse{Rue: "45 Avenue de la République", Ville: "Lyon", CodePostal: "69000"}

	emp1 := Employe{
		Personne: Personne{Prenom: "Kamil", Nom: "Garnier", Age: 34, Email: "kamil@company.com"},
		Adresse:  adr1,
		Poste:    "Lead Developer",
		Salaire:  4500.0,
	}
	emp2 := Employe{
		Personne: Personne{Prenom: "Alice", Nom: "Martin", Age: 28, Email: "alice@company.com"},
		Adresse:  adr2,
		Poste:    "Designer UI/UX",
		Salaire:  3200.0,
	}

	etu1 := Etudiant{
		Personne: Personne{Prenom: "Thomas", Nom: "Dubois", Age: 21, Email: "thomas@ecole.fr"},
		Promo:    "2026",
		Moyenne:  16.5,
	}
	etu2 := Etudiant{
		Personne: Personne{Prenom: "Lucas", Nom: "Robert", Age: 22, Email: "lucas@ecole.fr"},
		Promo:    "2026",
		Moyenne:  11.2,
	}

	emp1.AugmenterSalaire(10)

	fmt.Println("----------------------------------")
	fmt.Println(emp1.FicheEmploye())
	fmt.Println("----------------------------------")
	fmt.Println(emp2.FicheEmploye())
	fmt.Println("----------------------------------")
	fmt.Printf("ÉTUDIANT: %s\nPromo: %s | Moyenne: %.2f/20\nMention: %s\n", 
		etu1.NomCompleto(), etu1.Promo, etu1.Moyenne, etu1.MentionObtenue())
	fmt.Println("----------------------------------")
	fmt.Printf("ÉTUDIANT: %s\nPromo: %s | Moyenne: %.2f/20\nMention: %s\n", 
		etu2.NomCompleto(), etu2.Promo, etu2.Moyenne, etu2.MentionObtenue())
	fmt.Println("----------------------------------")
}