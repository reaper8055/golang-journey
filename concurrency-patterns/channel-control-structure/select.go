package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
* The select statement procides another way to handle multiple channels.
* It's like a switch, but each case is a communication:
* - All channels are evaluated.
* - Selection blocks until one communication can procee, which then does.
* - If multiple can proceed, select chooses pseudo-randomly.
* - A dafault clause, if present, executes immediately if no channel is ready.

example:

select {
  case v1 := <-c1:
    fmt.Printf("received %v from c1\n", v1)
  case v2 := <-c2:
    fmt.Printf("received %v from c2\n", v2)
  case c3 <- 23:
    fmt.Printlf("sent %v to c3\n", 23)
  default:
    fmt.Printf("no ne was ready to communicate\n")
}
**/

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring, I'm leaving.")
}
