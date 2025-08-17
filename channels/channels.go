package main

import (
	"fmt"
	// "math/rand"
	// "time"
)
// SENDING DATA TO CHANNEL
// func processNum(numChan chan int) {

// 	for num := range numChan {
// 		fmt.Println("processing number", num)
// 		time.Sleep(1 * time.Second)
// 	}

// }
// receiving data from channel
func sum(result chan int , num1 int, num2 int){
    numResult := num1 + num2
    result <- numResult
}







func main() {
    
  result := make( chan int )

   go sum(result, 6, 5)

   res := <-result // block
   fmt.Println(res)

	// numChan := make(chan int)
	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// messageChan := make(chan string) make channnel

	// messageChan <- "ping" // block

	// msg := <-messageChan  receved data from channel

	// fmt .Println(msg)
}
