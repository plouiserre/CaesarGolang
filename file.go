package main
import(
	"fmt" 
	"io/ioutil"
	"os"
)

type file struct{
	path string
}

func (f file) readFile() string{
	bs, err := ioutil.ReadFile(f.path)
	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	content := string(bs)
	return content
}