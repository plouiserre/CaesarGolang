package main

import (
	"fmt"
)



func main() {
	myLetters := letters{
		alphabet : []rune {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'},
	 } 

	text := "je suis le texte a chiffrer"
	fmt.Println(text)

	key := 5

	msgToCipher := cipher{
		key : key, 
		text : text,
		alphabet : myLetters,
	}

	newSentence := msgToCipher.cipherSentence()

	fmt.Printf("Voici la nouvelle phrase chiffre %s \n", newSentence)

	msgToDecipher := decipher{
		key : key, 
		text : newSentence,
		alphabet : myLetters,
	}

	oldSentence := msgToDecipher.decipherSentence()

	fmt.Printf("Voici l'ancienne phrase dechiffre %s \n", oldSentence)
}
