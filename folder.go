package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

type folder struct {
	pathRead string
	pathWrite string
	contentByFile map[string] string
	files []os.FileInfo
	msgToCipher cipher
	msgToDecipher decipher
	allCaracteres caracteres
}

func (f folder)WorkflowFolder (){
	f.contentByFile = make(map[string]string)

	f.ClearOutputDirectory()
	
	f.GetAllFiles()

	f.CipherFiles()
}

func (f *folder) GetAllFiles(){
	files, err := ioutil.ReadDir(f.pathRead)

	if(err != nil){
		fmt.Println(err);
	}

	f.files = files
}

//TODO structurer cette méthode
func (f folder)CipherFiles(){
	
	for _, fileReading := range f.files{
		fileName := fileReading.Name()
		fileToCipher := file{
			pathReadFile : f.pathRead+"/"+fileName,
			pathWriteFile : f.pathWrite+"/"+fileName,
		}

		content := fileToCipher.readFile()

		caesarWorkflow := workflow{
			text : content,
			caracter : f.allCaracteres,
			crypt : f.msgToCipher,
		}

		sentenceCiphered := caesarWorkflow.transformSentence()

		fmt.Printf("%s chiffre donne \n %s \n", fileName, sentenceCiphered)

		fileToCipher.writeCipherMessage(sentenceCiphered)

		caesarWorkflow.setIcipher(f.msgToDecipher)
		caesarWorkflow.setText(sentenceCiphered)

		oldSentence := caesarWorkflow.transformSentence()

		fmt.Printf("Voici l'ancienne phrase dechiffre %s \n", oldSentence)
	}
}

func (f folder) ClearOutputDirectory(){
	files, err := ioutil.ReadDir(f.pathWrite)

	if(err != nil){
		fmt.Println(err);
	}

	for _, fileToManage := range files{
		fileToDelete := file{
			pathWriteFile : f.pathWrite+"/"+fileToManage.Name(),
		}
		fileToDelete.deleteExistCipheredText()
	}
}