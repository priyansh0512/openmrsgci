package main

import ( 
	"fmt"
	"os"
	"path/filepath"
)	

func main() {
	// create a TestDir directory on current working directory
	os.Mkdir("." + string(filepath.Separator) + "TestDir" ,0777)
	fmt.Println("Hello World")
}
