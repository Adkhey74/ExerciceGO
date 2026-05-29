package main

import (
	"fmt"
	"math"
	"time"
)

func puissance(x, z int) int {
	return int(math.Pow(float64(x), float64(z)))
}

func age(anneeNaissance, moisNaissance, jourNaissance int) int {
	maintenant := time.Now()
	age := maintenant.Year() - anneeNaissance
	if maintenant.Month() < time.Month(moisNaissance) ||
		(maintenant.Month() == time.Month(moisNaissance) && maintenant.Day() < jourNaissance) {
		age--
	}
	return age
}

func main() {
	annee, mois, jour := 2003, 4, 4

	fmt.Println(puissance(64, 8))
	fmt.Println(time.Now().Format("02/01/2006 15:04:05"))
	fmt.Printf("Je suis né le %02d/%02d/%d, j'ai donc %d ans.\n", jour, mois, annee, age(annee, mois, jour))
}