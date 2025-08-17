package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"net/http"
)

type todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/23")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("error: status code:", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(data))

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(result)
	var todo todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(todo)
}

func performPostRequest() {
	var todo todo = todo{
		UserId:    12,
		Id:        13,
		Title:     "aayush",
		Completed: true,
	}
	// convert todo to json

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	//    conver json to data string
	jsonString := string(jsonData)

	//  conver to string data to reader

	jsonReader := strings.NewReader(jsonString)
	myUrl := "https://jsonplaceholder.typicode.com/todos/1"
	//  seND POST REQUEST
	res, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)

	fmt.Println("response:", string(data))
}

func performUpdateRequest() {
	todo := todo{
		UserId:    1254,
		Id:        13,
		Title:     "aayush golang",
		Completed: false,
	}
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	jsonString := string(jsonData)

	jsonReader := strings.NewReader(jsonString)

	const myUrl = "https://jsonplaceholder.typicode.com/todos/2"
	//   create put req
	req, err := http.NewRequest(http.MethodPut, myUrl, jsonReader)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	//  send request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)

	fmt.Println("response:", string(data))
	fmt.Println("status code:", res.Status)
}

func performDeleteRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/todos/2"

	// create delete request
	req, err := http.NewRequest(http.MethodDelete, myUrl, nil)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("status code:", res.Status)

}

func main() {
	fmt.Println("Hello World")
	// performGetRequest()
	// performPostRequest()
	// performUpdateRequest()
	performDeleteRequest()
}
