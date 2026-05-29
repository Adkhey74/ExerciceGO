package main

import "fmt"

// Matières scolaires numérotées automatiquement avec iota
const (
	Maths = iota
	Francais
	Histoire
	Anglais
	SVT
)

// Nom lisible d'une matière
func nomMatiere(m int) string {
	switch m {
	case Maths:
		return "Mathématiques"
	case Francais:
		return "Français"
	case Histoire:
		return "Histoire"
	case Anglais:
		return "Anglais"
	case SVT:
		return "SVT"
	default:
		return "Inconnue"
	}
}

func main() {
	fmt.Println("=== Bulletin scolaire ===\n")

	notes := []float64{14.5, 12.0, 16.0, 9.5, 11.0}

	var total float64
	for i, note := range notes {
		fmt.Printf("%-20s : %.1f/20\n", nomMatiere(i), note)
		total += note
	}

	notes = append(notes, 13.5)
	fmt.Printf("%-20s : %.1f/20  ← ajout en cours d'année\n", "Sport", notes[len(notes)-1])
	total += notes[len(notes)-1]

	moyenne := total / float64(len(notes))
	fmt.Printf("\nMoyenne générale : %.2f/20\n", moyenne)

	switch {
	case moyenne >= 16:
		fmt.Println("Mention : Très Bien")
	case moyenne >= 14:
		fmt.Println("Mention : Bien")
	case moyenne >= 12:
		fmt.Println("Mention : Assez Bien")
	case moyenne >= 10:
		fmt.Println("Mention : Admis")
	default:
		fmt.Println("Mention : Ajourné")
	}
}
