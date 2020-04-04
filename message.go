package main

import(	
	"strings"
)

type message struct {
	text string
	words []string
	key int
}

func (m *message)cypherSentence(alphabet letters) string {
	words := strings.Split(m.text, " ")
	var newSentence string
	for _, word := range words {
		newWord := m.cypherWord(word, alphabet)
		newSentence = newSentence + " " + newWord
	}
	return newSentence
}

func (m *message)cypherWord(word string, alphabet letters) string {
	lettersOfNewWord := []rune{}
	for _, letter := range word {
		indexLetter := alphabet.GetIndex(letter)
		newIndex := (indexLetter + m.key) % 25
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
