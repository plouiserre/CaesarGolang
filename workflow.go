package main

import(	
	"strings"
)

type workflow struct {
	text string
	alphabet letters
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
		indexLetter := w.alphabet.GetIndex(letter)
		newIndex :=  w.crypt.GetNewIndex(indexLetter)
		newLetter := w.alphabet.GetLetter(newIndex)
		lettersOfNewWord = append(lettersOfNewWord, newLetter)
	}
	newWord := string(lettersOfNewWord)
	return newWord
}
