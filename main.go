//Use 'go run main.go' to run this code.
package main

import (
	"fmt"
)

func hello(finished chan bool) {
	fmt.Println("Hello GoRoutine")
	finished <- true
}

func main() {
	finished := make(chan bool)
	go hello(finished)
	<-finished
	fmt.Printf("Main")
}
