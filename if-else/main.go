package main

import "fmt"


func main(){
	x:=10;

	if x >= 4{
		fmt.Println("X is greater than equal 4")
	}else{
		fmt.Println("X is less than 14")
	}

	z := 4
	if z >= 10{
		fmt.Println("Z is greater than equal 10")
	}else if z > 5{
		fmt.Println("Z is greater than 5")
	}else{
		fmt.Println("Z is less than 5")
	}

	y := 10 
	
	
	if x > 5 && y > 5{
		fmt.Println("hey")
	}else{
		fmt.Println("X and Y are less than 5")
	}
}