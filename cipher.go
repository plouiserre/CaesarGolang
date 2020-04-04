package main

import(	
	"strings"
)

type cipher struct {
	text string
	key int
	alphabet letters
}

func (m *cipher)cipherSentence() string {
	words := strings.Split(m.text, " ")
	var newSentence string
	for _, word := range words {
		newWord := m.cipherWord(word)
		newSentence = newSentence + " " + newWord
	}
	return newSentence
}

func (m *cipher)cipherWord(word string) string {
	lettersOfNewWord := []rune{}
	for _, letter := range word {
		indexLetter := m.alphabet.GetIndex(letter)
		newIndex := (indexLetter + m.key) % 25
		newLetter := m.alphabet.GetLetter(newIndex)
		lettersOfNewWord = append(lettersOfNewWord, newLetter)
	}
	newWord := string(lettersOfNewWord)
	return newWord
}
