package main

type decipher struct{
	key int
	length int
}

func (m decipher)GetNewIndex(indexLetter int) int {
	newIndex := 0
	if indexLetter -m.key>=0{ 
		newIndex = indexLetter - m.key
	}else {
		key := m.key
		for indexLetter > 0{
			indexLetter --
			key --
		}
		newIndex = m.length - key
	}
	return newIndex
}