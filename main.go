package main


func main() {
	allCaracteres := caracteres{
		alphabet : []rune {'a', 'à','ä','â','b', 'c', 'ç', 'd', 'e','é','è','ê','ë','f', 'g', 'h', 'i','î', 'j', 'k', 'l', 'm', 'n', 'o','œ','ö','ô', 'p', 'q', 'r', 's', 't', 'u','ù','û','v', 'w', 'x', 'y', 'z'},
		punctuation: []rune{'.',',','?',';',':','/','&','\'','(','§','!',')','-','_','+','=','*','…'},
	 } 

	data := file{
		pathReadFile:"data/file.txt",
		pathWriteFile:"data/cipheredText.txt",
	}

	allCaracteres.GenerateUppercaseAlphabet()

	data.deleteExistCipheredText()

	//text := data.readFile()

	//fmt.Println(text)

	key := 5

	msgToCipher := cipher{
		key : key, 
		length : len(allCaracteres.alphabet),
	}

	// caesarWorkflow := workflow{
	// 	text : text,
	// 	caracter : allCaracteres,
	// 	crypt : msgToCipher,
	// }

	// newSentence := caesarWorkflow.transformSentence()

	// fmt.Printf("Voici la nouvelle phrase chiffre %s \n", newSentence)
	
	// data.writeCipherMessage(newSentence)

	// msgToDecipher := decipher{
	// 	key : key,
	// 	length : len(allCaracteres.alphabet),
	// }

	// caesarWorkflow.setIcipher(msgToDecipher)
	// caesarWorkflow.setText(newSentence)

	// oldSentence := caesarWorkflow.transformSentence()

	// fmt.Printf("Voici l'ancienne phrase dechiffre %s \n", oldSentence)

	folderToWork := folder{
		pathRead: "data/input/",
		pathWrite: "data/output/",
		msgToCipher :msgToCipher,
		allCaracteres :allCaracteres,
	}

	folderToWork.WorkflowFolder()

}
