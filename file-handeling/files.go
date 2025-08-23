package main

import (
	"fmt"
	"io"
	// "io/ioutil"
	"os"
)

func main() {
	file, err := os.Create("exe.txt")
	if err != nil {
		fmt.Println("error while creating file:",err)
	
		return
	}
	defer file.Close()
	content := "hey aayush"
	bytes, errors := io.WriteString(file, content)
	fmt.Println("bytes written:", bytes)
	if errors != nil {
		fmt.Println("error while writing files:", err)
		return
	}
	

	fmt.Println("file created successfully")

	file, err = os.Open("exe.txt")
	if err != nil {
		fmt.Println("error while creating file:", err)
		return
	}
	defer file.Close()
      
	// create buffer  to read the file

	buffer := make([]byte, 1024)

	// read the file content into buffer

	for{
      n, err :=  file.Read(buffer)
	  if err == io.EOF{
		break
	  }
	  if err != nil{
		fmt.Println("error while reading file",err)
		return
	  }
	  fmt.Println(string(buffer[:n]))
	}
// 	// content ,err := ioutil.ReadFile(("exe.txt"))
// 	// if err !=nil{
// 	// 	fmt.Println("error while reading file",err)
// 	// 	return
// 	// }
// 	// fmt.Println(string(content))
}