package main

import "fmt"


 func add(a, b int) int {
	return a + b
 }
func main() {
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
	defer fmt.Println( "data is: ",add(1, 2))
}