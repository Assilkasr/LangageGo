package main

import (
	"fmt"
	"strings"
)

type Produit struct {
	ID        int
	Nom       string
	Marque    string
	Prix      float64
	Stock     int
	Categorie string
	Actif     bool
}

type Catalogue struct {
	produits []Produit
}

func (c *Catalogue) AjouterProduit(p Produit) error {
	for _, existant := range c.produits {
		if existant.ID == p.ID {
			return fmt.Errorf("produit avec l'ID %d existe déjà", p.ID)
		}
	}
	c.produits = append(c.produits, p)
	return nil
}

func (c Catalogue) TrouverParID(id int) (Produit, error) {
	for _, p := range c.produits {
		if p.ID == id {
			return p, nil
		}
	}
	return Produit{}, fmt.Errorf("aucun produit trouvé avec l'ID %d", id)
}

func (c Catalogue) TrouverParCategorie(cat string) []Produit {
	var resultats []Produit
	for _, p := range c.produits {
		if strings.EqualFold(p.Categorie, cat) {
			resultats = append(resultats, p)
		}
	}
	return resultats
}

func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	nb := 0
	for i, p := range c.produits {
		if strings.EqualFold(p.Categorie, categorie) {
			c.produits[i].Prix = p.Prix * (1 - pct/100)
			nb++
		}
	}
	return nb
}

func (c *Catalogue) Vendre(id int, qte int) error {
	for i, p := range c.produits {
		if p.ID == id {
			if p.Stock < qte {
				return fmt.Errorf("stock insuffisant : %d disponible(s), %d demandé(s)", p.Stock, qte)
			}
			c.produits[i].Stock -= qte
			return nil
		}
	}
	return fmt.Errorf("aucun produit trouvé avec l'ID %d", id)
}

func (c Catalogue) Rapport() string {
	valeurTotale := 0.0
	for _, p := range c.produits {
		valeurTotale += p.Prix * float64(p.Stock)
	}
	return fmt.Sprintf("Nombre de produits : %d\nValeur totale du stock : %.2f€", len(c.produits), valeurTotale)
}

func afficherProduit(p Produit) {
	fmt.Printf("  [%d] %s %s — %.2f€ | Stock: %d | Catégorie: %s\n",
		p.ID, p.Marque, p.Nom, p.Prix, p.Stock, p.Categorie)
}

func main() {
	catalogue := Catalogue{}

	produits := []Produit{
		{ID: 1, Nom: "iPhone 15 Pro", Marque: "Apple", Prix: 1229.00, Stock: 12, Categorie: "smartphone", Actif: true},
		{ID: 2, Nom: "MacBook Air M3", Marque: "Apple", Prix: 1499.00, Stock: 8, Categorie: "laptop", Actif: true},
		{ID: 3, Nom: "Galaxy S24 Ultra", Marque: "Samsung", Prix: 1099.00, Stock: 15, Categorie: "smartphone", Actif: true},
		{ID: 4, Nom: "ROG Zephyrus G14", Marque: "Asus", Prix: 1799.00, Stock: 5, Categorie: "laptop", Actif: true},
		{ID: 5, Nom: "AirPods Pro 2", Marque: "Apple", Prix: 279.00, Stock: 30, Categorie: "audio", Actif: true},
	}

	for _, p := range produits {
		catalogue.AjouterProduit(p)
	}

	for {
		fmt.Println("\n=== TECHSHOP — CATALOGUE ===")
		fmt.Println("[1] Ajouter un produit")
		fmt.Println("[2] Chercher un produit")
		fmt.Println("[3] Appliquer des soldes")
		fmt.Println("[4] Vendre un produit")
		fmt.Println("[5] Rapport du catalogue")
		fmt.Println("[0] Quitter")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 0:
			fmt.Println("À bientôt !")
			return

		case 1:
			var p Produit
			fmt.Print("ID : ")
			fmt.Scan(&p.ID)
			fmt.Print("Nom : ")
			fmt.Scan(&p.Nom)
			fmt.Print("Marque : ")
			fmt.Scan(&p.Marque)
			fmt.Print("Prix : ")
			fmt.Scan(&p.Prix)
			fmt.Print("Stock : ")
			fmt.Scan(&p.Stock)
			fmt.Print("Catégorie : ")
			fmt.Scan(&p.Categorie)
			p.Actif = true
			if err := catalogue.AjouterProduit(p); err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Produit ajouté avec succès.")
			}

		case 2:
			fmt.Println("[1] Par ID   [2] Par catégorie")
			fmt.Print("Choix : ")
			var sous int
			fmt.Scan(&sous)
			if sous == 1 {
				fmt.Print("ID : ")
				var id int
				fmt.Scan(&id)
				p, err := catalogue.TrouverParID(id)
				if err != nil {
					fmt.Println("Erreur :", err)
				} else {
					afficherProduit(p)
				}
			} else if sous == 2 {
				fmt.Print("Catégorie : ")
				var cat string
				fmt.Scan(&cat)
				resultats := catalogue.TrouverParCategorie(cat)
				if len(resultats) == 0 {
					fmt.Println("Aucun produit trouvé dans cette catégorie.")
				} else {
					for _, p := range resultats {
						afficherProduit(p)
					}
				}
			} else {
				fmt.Println("Choix invalide.")
			}

		case 3:
			fmt.Print("Catégorie à solder : ")
			var cat string
			fmt.Scan(&cat)
			fmt.Print("Réduction (%) : ")
			var pct float64
			fmt.Scan(&pct)
			if pct <= 0 || pct >= 100 {
				fmt.Println("Erreur : le pourcentage doit être entre 1 et 99.")
			} else {
				nb := catalogue.AppliquerReduction(cat, pct)
				fmt.Printf("%.0f%% appliqués sur %d produit(s) de la catégorie \"%s\".\n", pct, nb, cat)
			}

		case 4:
			fmt.Print("ID du produit : ")
			var id int
			fmt.Scan(&id)
			fmt.Print("Quantité : ")
			var qte int
			fmt.Scan(&qte)
			if qte <= 0 {
				fmt.Println("Erreur : la quantité doit être positive.")
			} else if err := catalogue.Vendre(id, qte); err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Printf("Vente enregistrée : %d exemplaire(s) vendus.\n", qte)
			}

		case 5:
			fmt.Println("\n--- RAPPORT ---")
			fmt.Println(catalogue.Rapport())

		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}
