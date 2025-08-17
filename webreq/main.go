package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("learning webreq service")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error getting", err)
		return
	}
	defer res.Body.Close()
	// fmt.Printf("type of res is: %T\n", res)
	// fmt.Println(res)
       
	//  read the response body
	 
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading", err)
		return
	}
	fmt.Println("response is: ", string(data))

}
