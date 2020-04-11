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

	key := 5

	msgToCipher := cipher{
		key : key, 
		length : len(allCaracteres.alphabet),
	}

	msgToDecipher := decipher{
		key : key,
		length : len(allCaracteres.alphabet),
	}

	folderToWork := folder{
		pathRead: "data/input/",
		pathWrite: "data/output/",
		msgToCipher :msgToCipher,
		msgToDecipher : msgToDecipher,
		allCaracteres :allCaracteres,
	}

	folderToWork.WorkflowFolder()

}
