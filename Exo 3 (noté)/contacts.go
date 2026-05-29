package main

import "fmt"

// ─────────────────────────────────────────
// 1. STRUCT PERSONNE
// ─────────────────────────────────────────

type Personne struct {
	Prenom string
	Nom    string
	Age    int
	Email  string
}

func (p Personne) NomComplet() string {
	return fmt.Sprintf("%s %s", p.Prenom, p.Nom)
}

func (p Personne) Presentation() string {
	return fmt.Sprintf("%s, %d ans — %s", p.NomComplet(), p.Age, p.Email)
}

// ─────────────────────────────────────────
// 2. STRUCT ADRESSE
// ─────────────────────────────────────────

type Adresse struct {
	Rue        string
	Ville      string
	CodePostal string
}

func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

// ─────────────────────────────────────────
// 3. STRUCT EMPLOYE
// ─────────────────────────────────────────

type Employe struct {
	Personne
	Adresse
	Poste   string
	Salaire float64
}

func (e Employe) FicheEmploye() string {
	return fmt.Sprintf(
		"👤 %s\n   Poste   : %s\n   Adresse : %s\n   Email   : %s\n   Âge     : %d ans\n   Salaire : %.2f €",
		e.NomComplet(),
		e.Poste,
		e.Format(),
		e.Email,
		e.Age,
		e.Salaire,
	)
}

func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire += e.Salaire * pct / 100
}

// ─────────────────────────────────────────
// 4. STRUCT ETUDIANT
// ─────────────────────────────────────────

type Etudiant struct {
	Personne
	Promo   string
	Moyenne float64
}

func (et Etudiant) MentionObtenue() string {
	switch {
	case et.Moyenne >= 16:
		return "Très Bien"
	case et.Moyenne >= 14:
		return "Bien"
	case et.Moyenne >= 12:
		return "Assez Bien"
	case et.Moyenne >= 10:
		return "Passable"
	default:
		return "Ajourné"
	}
}

func (et Etudiant) FicheEtudiant() string {
	return fmt.Sprintf(
		"🎓 %s\n   Promo   : %s\n   Email   : %s\n   Âge     : %d ans\n   Moyenne : %.2f/20 — %s",
		et.NomComplet(),
		et.Promo,
		et.Email,
		et.Age,
		et.Moyenne,
		et.MentionObtenue(),
	)
}

// ─────────────────────────────────────────
// 5. MAIN
// ─────────────────────────────────────────

func main() {
	// Deux employés
	e1 := Employe{
		Personne: Personne{"Alice", "Martin", 34, "alice.martin@empresa.fr"},
		Adresse:  Adresse{"12 rue de la Paix", "Paris", "75001"},
		Poste:    "Développeuse Go",
		Salaire:  3800.00,
	}
	e2 := Employe{
		Personne: Personne{"Karim", "Benali", 41, "k.benali@empresa.fr"},
		Adresse:  Adresse{"7 avenue Foch", "Lyon", "69006"},
		Poste:    "Chef de projet",
		Salaire:  4500.00,
	}

	// Augmentation de salaire
	e1.AugmenterSalaire(10)
	e2.AugmenterSalaire(5)

	// Deux étudiants
	et1 := Etudiant{
		Personne: Personne{"Sofia", "Dupont", 20, "sofia.dupont@univ.fr"},
		Promo:    "BUT Informatique 2026",
		Moyenne:  15.8,
	}
	et2 := Etudiant{
		Personne: Personne{"Noah", "Lefevre", 22, "noah.lefevre@univ.fr"},
		Promo:    "Licence Pro Dev Web 2026",
		Moyenne:  11.2,
	}

	// Affichage des fiches
	fmt.Println("════════════════════════════════")
	fmt.Println("        ANNUAIRE ENTREPRISE     ")
	fmt.Println("════════════════════════════════")

	fmt.Println("\n── Employés ──")
	fmt.Println(e1.FicheEmploye())
	fmt.Println()
	fmt.Println(e2.FicheEmploye())

	fmt.Println("\n── Étudiants ──")
	fmt.Println(et1.FicheEtudiant())
	fmt.Println()
	fmt.Println(et2.FicheEtudiant())
}
