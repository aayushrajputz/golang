package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Hello World")
	myUrl := "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"
	fmt.Printf("type of url is %T\n", myUrl)

parsedUrl, err	:=url.Parse(myUrl)
   if err != nil{
	fmt.Println("Error parsing url",err)
	return
   }
   fmt.Printf("type of url:%T\n", parsedUrl)
   fmt.Println(parsedUrl.Scheme)
   fmt.Println(parsedUrl.Host)
   fmt.Println(parsedUrl.Path)
   fmt.Println(parsedUrl.Port())
   fmt.Println(parsedUrl.RawQuery)
    //   modifying the url comp 
    parsedUrl.Path = "/newPath"
	parsedUrl.RawQuery = "username=imaayush"

	// constrating a url string from a url obj

	newUrl := parsedUrl.String()
	fmt.Println(newUrl)
	

}