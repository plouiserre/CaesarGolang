package main 

import(
	"fmt"
	"os"
)

type workflow struct{
	fileReading os.FileInfo
	folderCrossing folder
}

//TODO structer cette méthode
func (w workflow) ManipulateFile(){
	fileName := w.fileReading.Name()
		fileToCipher := file{
			pathReadFile : w.folderCrossing.pathRead+"/"+fileName,
			pathWriteFile : w.folderCrossing.pathWrite+"/"+fileName,
		}

		content := fileToCipher.readFile()

		caesarEncryption := encryption{
			text : content,
			caracter : w.folderCrossing.allCaracteres,
			crypt : w.folderCrossing.msgToCipher,
		}

		sentenceCiphered := caesarEncryption.transformSentence()

		fmt.Printf("%s chiffre donne \n %s \n", fileName, sentenceCiphered)

		fileToCipher.writeCipherMessage(sentenceCiphered)

		caesarEncryption.setIcipher(w.folderCrossing.msgToDecipher)
		caesarEncryption.setText(sentenceCiphered)

		oldSentence := caesarEncryption.transformSentence()

		fmt.Printf("Voici l'ancienne phrase dechiffre %s \n", oldSentence)
}