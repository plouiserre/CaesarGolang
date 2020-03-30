package main

import "fmt"

//TODO
// 1 - pouvoir parcourir les lettres de chiffrement OK
// 2 - faire l'alphabet OK
// 3 - faire la clé
// 4 - trouver la bonne lettre
// 5 - Gérer les espaces
// 6 - Afficher la phrase chiffré
// 7 - mettre le tout dans des structs

type letters []string

func main() {
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for _, letter := range letters {
		fmt.Printf(letter)
	}
	text := "je suis le texte a chiffrer"
	displayLetter(text)
}

func displayLetter(sentence string) {
	for _, word := range sentence {
		fmt.Printf("%c \n", word)
	}
}
