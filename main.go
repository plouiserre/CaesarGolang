package main

import "fmt"

//TODO
// 1 - pouvoir parcourir les lettres de chiffrement OK
// 2 - faire l'alphabet OK
// 3 - faire la clé + trouver la bonne lettre OK
// 4 - pouvoir chiffrer un mot
// 5 - Afficher la phrase chiffré
// 6 - mettre le tout dans des structs
type letters []string

func main() {
	myLetters := letters{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for _, letter := range myLetters {
		fmt.Printf(letter)
	}
	text := "je suis le texte a chiffrer"
	displayLetter(text)

	letter := myLetters.GetLetter(5)
	fmt.Println(letter)
}

func displayLetter(sentence string) {
	for _, word := range sentence {
		fmt.Printf("%c \n", word)
	}
}

func (l letters) GetLetter(key int) string {
	newKey := 0
	if key <= 25 {
		newKey = key
	} else {
		newKey = key % 26
	}
	result := l[newKey]
	return result
}
