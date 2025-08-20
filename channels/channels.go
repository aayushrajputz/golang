package main

import (
	"fmt"
	"time"
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
// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult
// }
//  go routine synchronizer
// func task(done chan bool) {

// 	defer func() { done <- true }()
// 	fmt.Println("process...")

// }

func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()

	

	
	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second)
	}
}
func main() {
      chan1 := make(chan int)
	  chan2 :=make(chan string) 

	   go func(){
        chan1 <- 10
	  } ()

	   go func(){
        chan2 <- "pong"
	  } ()

   for i := 0; i <2; i++{
	select{
	case chan1Val := <-chan1:
		fmt.Println("recevied data chan1", chan1Val)
	case chan2Val := <-chan2:
		fmt.Println("recevied data from chan2",chan2Val)
	}
   }

	   





	// emailChan := make(chan string, 100)

	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i := 0; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("done sending.")
	// // this is impo
	// close(emailChan)
	// <-done

	// emailChan <- "1@email.com"
	// emailChan <- "2@email.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	// done := make(chan bool)
	// go task(done)

	// <-done // block

	//   result := make( chan int )

	//    go sum(result, 6, 5)

	//    res := <-result // block
	//    fmt.Println(res)

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
