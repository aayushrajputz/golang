package main

import (
	"fmt"
	"strings"
)

func main(){
	// num := []int{1,2,3,4,5,6,7,8,9,10}
	// num = append(num,11,12,13)
	// fmt.Println(num)
	// fmt.Printf("number data type is %T\n",num)
	// fmt.Println(len(num))
	// number := make([]int ,3, 6)
	// number = append(number,7,97,6)
	// fmt.Println(number)
	// fmt.Println(len(number))
	// fmt.Println(cap(number))
  
     data := "apple, orange , banana , grapes"
	 parts:= strings.Split(data,".")
	 fmt.Println(parts)

	 str := "one two three four two two two five"
	 count := strings.Count(str,"five")
	 fmt.Println(count)
   
	  str = "     hello babe   "
	  trimmed := strings.TrimSpace(str)
	  fmt.Println(trimmed)

	  str1:= "rajput"
	  str2 := "aayush"
	  result := strings.Join([]string{str1,"MR.",str2} ," ")
	  fmt.Println(result)
	
}