package main

import "fmt"

const (
	Nom          = "Adil"
	IMCMaigreur  = 18.5
	IMCNormal    = 25.0
	IMCSurpoids  = 30.0
)

func main() {
	var poids, taille float64

	fmt.Printf("Bonjour %s ! Calculons ton IMC.\n", Nom)
	fmt.Print("Entre ton poids (kg) : ")
	fmt.Scan(&poids)
	fmt.Print("Entre ta taille (m)  : ")
	fmt.Scan(&taille)

	imc := poids / (taille * taille)

	fmt.Printf("\nTon IMC est : %.2f\n", imc)

	var categorie string
	if imc < IMCMaigreur {
		categorie = "Maigreur"
	} else if imc < IMCNormal {
		categorie = "Normal"
	} else if imc < IMCSurpoids {
		categorie = "Surpoids"
	} else {
		categorie = "Obésité"
	}

	fmt.Printf("Catégorie     : %s\n", categorie)
}
