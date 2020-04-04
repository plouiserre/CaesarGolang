package main

type cipher struct {
	key int
}

func (m cipher)GetNewIndex(indexLetter int) int {
	newIndex := (indexLetter + m.key) % 25
	return newIndex
} 
