package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isAdult"`
}

func main() {
	person := Person{Name: "Aayush", Age: 20, IsAdult: true}
	fmt.Println(person)

	// CONVERT STRUCT TO JSON
	// json.Marshal(person)

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))


	// decoding (unmarshalling)

	var person2 Person
	err = json.Unmarshal(jsonData, &person2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(person2)
}
