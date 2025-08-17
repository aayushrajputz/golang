package main

import "fmt"

func main() {
	for i:=1; i<=1; i++ {
		fmt.Println(i)
	}
	counter := 0
       for{
		fmt.Println("loops")
		counter++
		if counter == 1{
			break;
		}
	   }
	   number := []int{11, 42, 83, 14, 75}
	   for index , value := range number{
		fmt.Println(index, value)
	   }
	   data := "hello world"
	   for index, char := range data{
		fmt.Println(index, string(char))
	   }
	   
}
