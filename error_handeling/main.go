package main

import (
	"fmt"
	
)

func divide(a, b float64) (float64, error ){
	if b ==0{
      	return 0, fmt.Errorf("cannot divide by zero")
	}
	return a/b, nil
}

func main() {
	fmt.Println("started error handeling in func")
	ans, _ := divide(12 , 4);
	fmt.Println(ans,);
}