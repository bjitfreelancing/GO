/*
communication between goroutines (using share memory by communicating [channels])
channels can also be used for synchronizing goroutines
*/

package main

import (
	"fmt"
	"time"
)

// consumer
func main() {
	/*
		var ch chan int
		ch = make(chan int)
	*/
	ch := make(chan int)
	fmt.Println("main started")
	go add(100, 200, ch)
	result := <-ch
	fmt.Println("result :", result)
	fmt.Println("main completed")
}

// producer
func add(x, y int, ch chan<- int) {
	time.Sleep(5 * time.Second)
	result := x + y
	ch <- result // send the result through the channel
}
