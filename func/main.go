package main

import (
	"fmt"
	
)
  func simplefunc(){
	fmt.Println("dear")
}
 
   func multi(a , b int) int{
	return a * b
}

func main() {
	fmt.Println("Hello World")
	simplefunc()
	
	ans := multi(4,8)
	fmt.Println(ans)
}