package main

import (
	"fmt"
	"strings"
)

//TODO
// 1 - pouvoir parcourir les lettres de chiffrement OK
// 2 - faire l'alphabet OK
// 3 - faire la clé + trouver la bonne lettre OK
// 4 - pouvoir chiffrer un mot OK
// 5 - Afficher la phrase chiffré OK
// 6 - pouvoir tout saisir par interface de commande
// 7 - mettre le tout dans des structs
type letters []rune

func main() {
	myLetters := letters{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	text := "je suis le texte a chiffrer"
	fmt.Println(text)

	key := 5

	newSentence := cypherSentence(text, myLetters, key)

	fmt.Printf("Voici la nouvelle phrase chiffre %s \n", newSentence)
}

func cypherSentence(sentence string, alphabet letters, key int) string {
	words := strings.Split(sentence, " ")
	var newSentence string
	for _, word := range words {
		newWord := cypherWord(word, alphabet, key)
		newSentence = newSentence + " " + newWord
	}
	return newSentence
}

func cypherWord(word string, alphabet letters, key int) string {
	lettersOfNewWord := []rune{}
	for _, letter := range word {
		indexLetter := alphabet.GetIndex(letter)
		newIndex := (indexLetter + key) % 25
		newLetter := alphabet.GetLetter(newIndex)
		lettersOfNewWord = append(lettersOfNewWord, newLetter)
	}
	newWord := string(lettersOfNewWord)
	return newWord
}

func (l letters) GetIndex(letterToPlace rune) int {
	index := 0
	for i := range l {
		if l[i] == letterToPlace {
			index = i
			break
		}
	}

	return index
}

func (l letters) GetLetter(index int) rune {
	newIndex := 0
	if index <= 25 {
		newIndex = index
	} else {
		newIndex = index % 25
	}
	result := l[newIndex]
	return result
}
