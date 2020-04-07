package main

import(
	"unicode"
)

type caracteres struct {
	alphabet [] rune
	punctuation [] rune
	uppercaseAlphabet [] rune
}

func (c *caracteres) GenerateUppercaseAlphabet(){
	c.uppercaseAlphabet = []rune{}
	for _, letter := range c.alphabet{
		upperLetter := unicode.ToUpper(letter)
		c.uppercaseAlphabet = append(c.uppercaseAlphabet, upperLetter)
	}
}

//TODO factoriser GetLetterIndex et GetUpperCaseIndex
func (c *caracteres) GetLetterIndex(caracteresToPlace rune) int {
	index := 0
	for i := range c.alphabet {
		if c.alphabet[i] == caracteresToPlace {
			index = i
			break
		}
	}

	return index
}

func (c *caracteres) GetUppercaseIndex(caracteresToPlace rune) int {
	index := 0
	for i := range c.uppercaseAlphabet {
		if c.uppercaseAlphabet[i] == caracteresToPlace {
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

func (c *caracteres) GetUppercase(index int) rune {
	newIndex := 0
	if index <= 25 {
		newIndex = index
	} else {
		newIndex = index % 25
	}
	result := c.uppercaseAlphabet[newIndex]
	return result
}

//TODO facotriser IsPunctuation et IsUpperCase et IsNormalLetter
func (c *caracteres) IsPunctuation (element rune) bool{
	result := false
	for _, punc := range c.punctuation{
		if element == punc{
			result = true;
		}
	}
	return result;
}

func (c *caracteres) IsUpperCase (element rune) bool{
	result := false
	for _, upperLetter := range c.uppercaseAlphabet{
		if element == upperLetter{
			result = true;
		}
	}
	return result
}

func (c *caracteres) IsNormalLetter (element rune) bool{
	result := false
	for _, letter := range c.alphabet{
		if element == letter{
			result = true;
		}
	}
	return result
}
