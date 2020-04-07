package main

import (
	"fmt"
)

func main() {
	allCaracteres := caracteres{
		alphabet : []rune {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'},
		punctuation: []rune{'.',',','?',';',':','/','&','\'','(','§','!',')','-','_','+','=','*'},
	 } 

	data := file{
		pathReadFile:"data/file.txt",
		pathWriteFile:"data/cipheredText.txt",
	}

	data.deleteExistCipheredText()

	text := data.readFile()

	fmt.Println(text)

	key := 5

	msgToCipher := cipher{
		key : key, 
	}

	caesarWorkflow := workflow{
		text : text,
		caracter : allCaracteres,
		crypt : msgToCipher,
	}

	newSentence := caesarWorkflow.transformSentence()

	fmt.Printf("Voici la nouvelle phrase chiffre %s \n", newSentence)
	
	data.writeCipherMessage(newSentence)

	msgToDecipher := decipher{
		key : key,
	}

	caesarWorkflow.setIcipher(msgToDecipher)
	caesarWorkflow.setText(newSentence)

	oldSentence := caesarWorkflow.transformSentence()

	fmt.Printf("Voici l'ancienne phrase dechiffre %s \n", oldSentence)

}
