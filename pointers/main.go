package main

import "fmt"

    func modifyReference(num *int) {
		*num = *num * 3
	}     

	func main() {
		// var num int 
		// num = 10
		// var ptr *int
		// ptr = &num
		// fmt.Println("Value of num is", num)
		// fmt.Println("Address of num is", &num)
		// fmt.Println("Value of ptr is", ptr)
		// fmt.Println("Address of ptr is", &ptr)
		// fmt.Println("Value of num using ptr is", *ptr)
		num :=10
		ptr:= &num
		fmt.Println("Value of num is", num)
		fmt.Println("Address of num is", &num)
		fmt.Println("Value of ptr is", ptr)
		fmt.Println("Address of ptr is", &ptr)
		fmt.Println("Value of num using ptr is", *ptr)
		

		var pointer *int
		nums := 101
		pointer = &nums
		if pointer == nil {
			fmt.Println("Pointer is nil")
		}else{
			fmt.Println("Pointer is not nil" ,*pointer)
		}


		value :=79
		modifyReference(&value)
		fmt.Println("Value is", value)
	}
