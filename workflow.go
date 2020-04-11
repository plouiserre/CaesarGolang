package main

import(	
	"strings"
	"fmt"
	"os"
)

type workflow struct {
	text string
	caracter caracteres
	crypt icipher 
	lettersOfNewWord []rune
}

func (w *workflow) setIcipher(cipher icipher){
	w.crypt = cipher
}

func(w *workflow) setText(newText string){
	w.text = newText
}

func (w *workflow)transformSentence() string {
	words := strings.Split(w.text, " ")
	var newSentence string
	for _, word := range words {
		newWord := w.transformWord(word)
		newSentence = newSentence + " " + newWord
	}
	return newSentence
}

//TODO réécrire cette méthode (étape 8)
func (w *workflow)transformWord(word string) string {
	w.lettersOfNewWord = []rune{}
	for _, letter := range word {
		isUpperCase := false 
		isNormalLetter := false
		isPunctuation := w.caracter.IsSpecificCaracters(letter, w.caracter.punctuation)
		if(isPunctuation == false){
			isUpperCase =  w.caracter.IsSpecificCaracters(letter, w.caracter.uppercaseAlphabet)
			if(isUpperCase == false){
				isNormalLetter = w.caracter.IsSpecificCaracters(letter, w.caracter.alphabet)
			}
		}
		
		//mettre une étape pour faire un continue si c'est ' et l'ajouter dans le truc a chiffrer
		//TODO améliorer et le mettre en méthode
		if(letter == 8217 || letter == 10){
			w.lettersOfNewWord = append(w.lettersOfNewWord, letter)
			continue
		}

		if(isPunctuation == false && isUpperCase == false && isNormalLetter == false){
			fmt.Printf("Dans le mot %s il y a le caractère %s inconnu donc on arrête tout \n", word, string(letter))
			os.Exit(1)
		}

		if(isPunctuation == false){
			if(isUpperCase == false){
				indexLetter := w.caracter.GetSpecificIndex(letter, w.caracter.alphabet)
				newIndex :=  w.crypt.GetNewIndex(indexLetter)
				newLetter := w.caracter.GetSpecificCaracters(newIndex, w.caracter.alphabet)
				w.lettersOfNewWord = append(w.lettersOfNewWord, newLetter)
			}else {
				indexLetter := w.caracter.GetSpecificIndex(letter, w.caracter.uppercaseAlphabet)
				newIndex :=  w.crypt.GetNewIndex(indexLetter)
				newLetter := w.caracter.GetSpecificCaracters(newIndex, w.caracter.uppercaseAlphabet)
				w.lettersOfNewWord = append(w.lettersOfNewWord, newLetter)
			}
		} else {
			w.lettersOfNewWord = append(w.lettersOfNewWord, letter)
		}
	}
	newWord := string(w.lettersOfNewWord)
	return newWord
}
