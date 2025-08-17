package main 
import (
	"fmt"
)

func main() {
	month := "janub"
	switch month{
	case "janu","feb","mar":
		fmt.Println("this is the first quarter")
	case "apr","may","jun":
		fmt.Println("this is the second quarter")
	case "jul","aug","sep":
		fmt.Println("this is the third quarter")
	case "oct","nov","dec":
		fmt.Println("this is the fourth quarter")
		default:
		fmt.Println("this is not a month")
	}
}