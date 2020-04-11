package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

type folder struct {
	path string
	contentByFile map[string] string
	files []os.FileInfo
	msgToCipher cipher
	allCaracteres caracteres
}

func (f folder)WorkflowFolder (){
	f.contentByFile = make(map[string]string)
	
	f.GetAllFiles()

	f.CipherFiles()
	
}

func (f *folder) GetAllFiles(){
	//lister les fichiers 
	files, err := ioutil.ReadDir(f.path)

	if(err != nil){
		fmt.Println(err);
	}

	f.files = files
}

func (f folder)CipherFiles(){
	
	for _, fileReading := range f.files{
		fileName := fileReading.Name()
		fileName = f.GetFilesPath(fileName)
		fileToCipher := file{
			pathReadFile : fileName,
		}

		//TODO ne pas passer par deux étapes mais directement par une seul
		content := fileToCipher.readFile()

		content = fileToCipher.manageApostrophe(content)

		caesarWorkflow := workflow{
			text : content,
			caracter : f.allCaracteres,
			crypt : f.msgToCipher,
		}

		sentenceCiphered := caesarWorkflow.transformSentence()

		fmt.Printf("%s chiffre donne \n %s \n", fileName, sentenceCiphered)
	}
}

func (f folder) WriteCipherResult(){

}

func (f folder) GetFilesPath(fileName string )string{
	return f.path+"/"+fileName
}