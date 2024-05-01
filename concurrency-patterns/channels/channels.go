package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
* Declaring and initializing:
* var c chan Int
* c = make(chan int)
* or c := make(chan int)
*
* Sending on a channel:
* c <- 1
*
* Receiving from a channel
* The "arrow" indicates the direction of data flow.
* value = <-c
**/

// Using channels

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		// Expression to be sent can be any suitable value.
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		// Receive expression is just a value.
		fmt.Printf("you say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}
