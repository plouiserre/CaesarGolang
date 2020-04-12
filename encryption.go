package main

import(	
	"strings"
	"fmt"
	"os"
)

type encryption struct {
	text string
	caracter caracteres
	crypt icipher 
	lettersOfNewWord []rune
}

func (e *encryption) setIcipher(cipher icipher){
	e.crypt = cipher
}

func(e *encryption) setText(newText string){
	e.text = newText
}

func (e *encryption)transformSentence() string {
	words := strings.Split(e.text, " ")
	var newSentence string
	for _, word := range words {
		newWord := e.transformWord(word)
		newSentence = newSentence + " " + newWord
	}
	return newSentence
}

//TODO réécrire cette méthode (étape 8)
func (e *encryption)transformWord(word string) string {
	e.lettersOfNewWord = []rune{}
	for _, letter := range word {
		isUpperCase := false 
		isNormalLetter := false
		isPunctuation := e.caracter.IsSpecificCaracters(letter, e.caracter.punctuation)
		if(isPunctuation == false){
			isUpperCase =  e.caracter.IsSpecificCaracters(letter, e.caracter.uppercaseAlphabet)
			if(isUpperCase == false){
				isNormalLetter = e.caracter.IsSpecificCaracters(letter, e.caracter.alphabet)
			}
		}
		
		//mettre une étape pour faire un continue si c'est ' et l'ajouter dans le truc a chiffrer
		//TODO améliorer et le mettre en méthode
		if(letter == 8217 || letter == 10){
			e.lettersOfNewWord = append(e.lettersOfNewWord, letter)
			continue
		}

		if(isPunctuation == false && isUpperCase == false && isNormalLetter == false){
			fmt.Printf("Dans le mot %s il y a le caractère %s inconnu donc on arrête tout \n", word, string(letter))
			os.Exit(1)
		}

		if(isPunctuation == false){
			if(isUpperCase == false){
				indexLetter := e.caracter.GetSpecificIndex(letter, e.caracter.alphabet)
				newIndex :=  e.crypt.GetNewIndex(indexLetter)
				newLetter := e.caracter.GetSpecificCaracters(newIndex, e.caracter.alphabet)
				e.lettersOfNewWord = append(e.lettersOfNewWord, newLetter)
			}else {
				indexLetter := e.caracter.GetSpecificIndex(letter, e.caracter.uppercaseAlphabet)
				newIndex :=  e.crypt.GetNewIndex(indexLetter)
				newLetter := e.caracter.GetSpecificCaracters(newIndex, e.caracter.uppercaseAlphabet)
				e.lettersOfNewWord = append(e.lettersOfNewWord, newLetter)
			}
		} else {
			e.lettersOfNewWord = append(e.lettersOfNewWord, letter)
		}
	}
	newWord := string(e.lettersOfNewWord)
	return newWord
}
