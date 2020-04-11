package main

type cipher struct {
	key int
	length int
}

func (m cipher)GetNewIndex(indexLetter int) int {
	newIndex := (indexLetter + m.key) % m.length
	return newIndex
} 
