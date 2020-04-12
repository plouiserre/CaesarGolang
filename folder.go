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

func (f folder)CipherFiles(){
	
	for _, fileReading := range f.files{

		f.CipherFile(fileReading)
	}
}

func (f folder) CipherFile(fileReading os.FileInfo){
	workflowEncryption := workflow {
		fileReading : fileReading,
		folderCrossing : f,
	}

	workflowEncryption.ManipulateFile()
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