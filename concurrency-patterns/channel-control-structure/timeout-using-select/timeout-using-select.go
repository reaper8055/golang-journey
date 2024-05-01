package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
* The time.After function returns a channel that blocks for the specified
* duration. After the interval, the channel delivers the current time, once.
**/

func main() {
	timeoutForWholeConversation()
}

func timeoutAfterOneSeccond() {
	c := boring("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("You're too slow.")
			return
		}
	}
}

func timeoutForWholeConversation() {
	c := boring("Joe")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You talk too much.")
			return
		}
	}
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
