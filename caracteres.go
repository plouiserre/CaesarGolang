package main

import(
	"unicode"
)

type caracteres struct {
	alphabet [] rune
	punctuation [] rune
	uppercaseAlphabet [] rune
	isUpperCase bool
	isNormalLetter bool 
	isPunctuation bool
}

func (c *caracteres) GenerateUppercaseAlphabet(){
	c.uppercaseAlphabet = []rune{}
	for _, letter := range c.alphabet{
		upperLetter := unicode.ToUpper(letter)
		c.uppercaseAlphabet = append(c.uppercaseAlphabet, upperLetter)
	}
}
func (c *caracteres) GetSpecificIndex(caracteresToPlace rune, listCaracteres []rune) int {
	index := 0
	for i := range listCaracteres {
		if listCaracteres[i] == caracteresToPlace {
			index = i
			break
		}
	}

	return index
}

func (c *caracteres) GetSpecificCaracters(index int, listCaracteres []rune) rune {
	newIndex := 0
	max := len(listCaracteres)
	if index <= max {
		newIndex = index
	} else {
		newIndex = index % max
	}

	result := listCaracteres[newIndex]
	return result
}

func (c *caracteres) GetUppercase(index int) rune {
	newIndex := 0
	max := len(c.uppercaseAlphabet)
	if index <= max {
		newIndex = index
	} else {
		newIndex = index % max
	}
	result := c.uppercaseAlphabet[newIndex]
	return result
}


func (c *caracteres) DeterminesTypeCaracters(letter rune){
		c.isPunctuation =c.IsSpecificCaracters(letter, c.punctuation)
		if(c.isPunctuation == false){
			c.isUpperCase =  c.IsSpecificCaracters(letter, c.uppercaseAlphabet)
			if(c.isUpperCase == false){
				c.isNormalLetter = c.IsSpecificCaracters(letter, c.alphabet)
			}
		}
}

func (c *caracteres) IsSpecificCaracters (element rune, caracteres []rune) bool{
	result := false
	for _, punc := range caracteres{
		if element == punc{
			result = true;
		}
	}
	return result;
}