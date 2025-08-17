package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 100
	fmt.Println(num)
	fmt.Printf("type of num iS ,%T", num)

	var data float64 = float64(num)
	data = data + 0.5
	fmt.Println(data)
	fmt.Printf("type of data is %T", data)

	num = 123
	str := strconv.Itoa(num)
	str = str + "456"
	fmt.Println(str)
	fmt.Printf("type of str is %T", str)

	num_str := "1234"
	num_int, _ := strconv.Atoi(num_str)
	num_int = num_int + 1
	fmt.Println(num_int)
	fmt.Printf("type of num_int is %T", num_int)

	number_str := "1234.56"
	num_float, _ := strconv.ParseFloat(number_str, 64)
	fmt.Println("num_float is", num_float)
	fmt.Printf("type of num_float is %T", num_float)

}
