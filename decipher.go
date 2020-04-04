package main

import (
	"strings"
)

type decipher struct{
	text string
	key int
	alphabet letters
}

func (m *decipher)decipherSentence() string {
	words := strings.Split(m.text, " ")
	var newSentence string
	for _, word := range words {
		newWord := m.decipherWord(word)
		newSentence = newSentence + " " + newWord
	}
	return newSentence
}

func (m *decipher)decipherWord(word string) string {
	lettersOfNewWord := []rune{}
	for _, letter := range word {
		indexLetter := m.alphabet.GetIndex(letter)
		newIndex := m.GetNewIndex(indexLetter)
		newLetter := m.alphabet.GetLetter(newIndex)
		lettersOfNewWord = append(lettersOfNewWord, newLetter)
	}
	newWord := string(lettersOfNewWord)
	return newWord
}

func (m *decipher)GetNewIndex(indexLetter int) int {
	newIndex := 0
	if indexLetter -m.key>=0{ 
		newIndex = indexLetter - m.key
	}else {
		key := m.key
		for indexLetter > 0{
			indexLetter --
			key --
		}
		newIndex = 25 - key
	}
	return newIndex
}