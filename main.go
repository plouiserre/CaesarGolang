package main

import (
	"fmt"
)

//TODO
// 1 - pouvoir parcourir les lettres de chiffrement OK
// 2 - faire l'alphabet OK
// 3 - faire la clé + trouver la bonne lettre OK
// 4 - pouvoir chiffrer un mot OK
// 5 - Afficher la phrase chiffré OK
// 6 - mettre le tout dans des structs OK
type letters []rune

func main() {
	myLetters := letters{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	text := "je suis le texte a chiffrer"
	fmt.Println(text)

	key := 5

	msg := message{
		key : key, 
		text : text,
	}

	newSentence := msg.cypherSentence(myLetters)

	fmt.Printf("Voici la nouvelle phrase chiffre %s \n", newSentence)
}
