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

func (e *encryption)transformWord(word string) string {
	e.lettersOfNewWord = []rune{}
	for _, letter := range word {
		//8217 => ' 10 => &amp;
		if(letter == 8217 || letter == 10){
			e.lettersOfNewWord = append(e.lettersOfNewWord, letter)
			continue
		}
		
		carac := e.caracter
		carac.DeterminesTypeCaracters(letter)	

		if(carac.isPunctuation == false && carac.isUpperCase == false && carac.isNormalLetter == false){
			fmt.Printf("Dans le mot %s il y a le caractère %s inconnu donc on arrête tout \n", word, string(letter))
			os.Exit(1)
		}

		e.GetLetterCiphered(letter, carac)
	}
	newWord := string(e.lettersOfNewWord)
	return newWord
}

func (e *encryption) GetLetterCiphered(letter rune, carac caracteres){
	if(carac.isPunctuation == false){
		if(carac.isUpperCase == false){
			e.lettersOfNewWord = e.DetermineNewLetter(e.caracter.alphabet, letter)
		}else {
			e.lettersOfNewWord = e.DetermineNewLetter(e.caracter.uppercaseAlphabet, letter)
		}
	} else {
		e.lettersOfNewWord = append(e.lettersOfNewWord, letter)
	}
}

func (e encryption) DetermineNewLetter(alphabet []rune, letter rune) []rune{
	indexLetter := e.caracter.GetSpecificIndex(letter, alphabet)
	newIndex :=  e.crypt.GetNewIndex(indexLetter)
	newLetter := e.caracter.GetSpecificCaracters(newIndex, alphabet)
	lettersCiphered := append(e.lettersOfNewWord, newLetter)
	return lettersCiphered
}
