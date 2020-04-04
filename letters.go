package main

type letters struct {
	alphabet [] rune
}

func (l *letters) GetIndex(letterToPlace rune) int {
	index := 0
	for i := range l.alphabet {
		if l.alphabet[i] == letterToPlace {
			index = i
			break
		}
	}

	return index
}

func (l *letters) GetLetter(index int) rune {
	newIndex := 0
	if index <= 25 {
		newIndex = index
	} else {
		newIndex = index % 25
	}
	result := l.alphabet[newIndex]
	return result
}
