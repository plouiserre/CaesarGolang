package main

import "fmt"

//TODO
// 1 - pouvoir parcourir les lettres de chiffrement OK
// 2 - faire l'alphabet
// 3 - faire la clé
// 4 - trouver la bonne lettre
// 5 - Gérer les espaces
// 6 - Afficher la phrase chiffré
// 7 - mettre le tout dans des structs
// 8 - commiter

func main() {
	text := "je suis le texte a chiffrer"
	displayLetter(text)
}

func displayLetter(sentence string) {
	for _, word := range sentence {
		fmt.Printf("%c \n", word)
	}
}
