package main
import(
	"fmt" 
	"io/ioutil"
	"os"
)

type file struct{
	pathReadFile string
	pathWriteFile string
}

func (f file) readFile() string{
	bs, err := ioutil.ReadFile(f.pathReadFile)
	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	content := string(bs)
	return content
}

func (f file) writeCipherMessage(msg string){
	err := ioutil.WriteFile(f.pathWriteFile, []byte(msg), 0666)
	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func (f file)deleteExistCipheredText(){
	fileExist := f.fileExist()
	if fileExist{
		err := os.Remove(f.pathWriteFile)
		if err != nil{
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}
}

func (f file) fileExist() bool {
	info, err := os.Stat(f.pathWriteFile)
	if os.IsNotExist(err){
		return false
	}
	return !info.IsDir()
}
