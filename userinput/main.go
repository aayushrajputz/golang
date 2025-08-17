package main

import ("fmt"
         "bufio"
		 "os"
)
func main() {
	fmt.Println("Enter your name")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Hello mr.", name);
	// fmt.Println("Enter your age")
	// var age int
	// fmt.Scan(&age)
	// fmt.Println("Your age is", age)
	// fmt.Println("You will be 100 years old in", time.Now().Year() + 100 - age)

reader := bufio.NewReader(os.Stdin)
name, _ := reader.ReadString('\n')
fmt.Println("Hello mr.",(name))
}
