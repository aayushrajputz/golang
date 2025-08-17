package main 

import "fmt"

func main() {
	fmt.Println("we are learning about arrays ")

	var name[5]string
	name[0] = "aayush"
	name[1] = "singh"
	name[2] = "is"
	name[3] = "a"
	name[4] = "good"
	fmt.Println(name)

	var num = [19]int{1, 2, 3, 4 ,5}
	fmt.Println(num)
	fmt.Println(len(num))

	fmt.Println(name[1])

	var price =  [5]string{"aayush",  "singh"}

	fmt.Printf("price array is %q\n", price)
	

}