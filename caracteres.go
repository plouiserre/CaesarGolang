package main

type caracteres struct {
	alphabet [] rune
	punctuation [] rune
}

func (c *caracteres) GetIndex(caracteresToPlace rune) int {
	index := 0
	for i := range c.alphabet {
		if c.alphabet[i] == caracteresToPlace {
			index = i
			break
		}
	}

	return index
}

func (c *caracteres) GetLetter(index int) rune {
	newIndex := 0
	if index <= 25 {
		newIndex = index
	} else {
		newIndex = index % 25
	}
	result := c.alphabet[newIndex]
	return result
}

func (c *caracteres) IsPunctuation (element rune) bool{
	result := false;
	for _, punc := range c.punctuation{
		if element == punc{
			result = true;
		}
	}
	return result;
}
