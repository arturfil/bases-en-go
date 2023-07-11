package concurrency

import (
	"fmt"
	"time"
)

func ChannelMain() {
	/*
	   c := make(chan string, 2)
	   c <- "hello"
	   c <- "world"

	   msg := <- c
	   fmt.Println(msg)

	   msg = <- c
	   fmt.Println(msg)

	   c = make(chan string)
	   go countCh("sheep", c)
	*/

	/*for {
	    msg, open := <- c
	    if !open { break }
	    fmt.Println(msg)
	}*/
	// no need to check for open!
	/*for msg := range c {
	    fmt.Println(msg)
	}*/

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
        for {
            c2 <- "Every two seconds"
            time.Sleep(time.Second * 2)
        }
    }()

    // here we are blocking, waiting for the second one to finish
    for {
       fmt.Println(<- c1) 
       fmt.Println(<- c2) 
    }
}

func countCh(item string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- item
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
