package main

import(	
	"strings"
)

type workflow struct {
	text string
	caracter caracteres
	crypt icipher
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

func (w *workflow)transformWord(word string) string {
	lettersOfNewWord := []rune{}
	for _, letter := range word {
			isPunctuation := w.caracter.IsPunctuation(letter)
			if(isPunctuation == false){
			indexLetter := w.caracter.GetIndex(letter)
			newIndex :=  w.crypt.GetNewIndex(indexLetter)
			newLetter := w.caracter.GetLetter(newIndex)
			lettersOfNewWord = append(lettersOfNewWord, newLetter)
		} else {
			lettersOfNewWord = append(lettersOfNewWord, letter)
		}
	}
	newWord := string(lettersOfNewWord)
	return newWord
}
