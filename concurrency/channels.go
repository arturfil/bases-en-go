package concurrency

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func ChannelsMain() {
    // unbufferedChannels() 
    // bufferedChannels()
    channelsSelect()
}

func unbufferedChannels() {
    c := make(chan int)
    start := time.Now()

    processed := []int{}

    for w := 1; w <= 6; w++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for val := range c {
                time.Sleep(time.Millisecond * 1000)
                // emulating proccesses
                fmt.Println("value -> ", val)
                val *= val
                val += 1
                processed = append(processed, val)
            }
        }()
    }
    
    // blocking code
    for i := 1; i <= 10; i++ {
        c <- i
    }

    close(c)
    wg.Wait()

    end := time.Since(start)
    fmt.Println("processed", processed)
    fmt.Println("time taken:", end)
}

func bufferedChannels() {
    c := make(chan int, 2)
    start := time.Now()

    // wg.Add(1)
    for w := 1; w <= 5; w++ {
        wg.Add(1)
        go func() {
            defer wg.Done()

            for val := range c {
                fmt.Println("Read\t", val, "\tfrom channel")
                time.Sleep(time.Millisecond * 2000)
            }
        }()

    }
    
    for i := 1; i <= 10; i++ {
        c <- i
        fmt.Println("Wrote\t", i, "\tto channel")
    }

    close(c) 
    wg.Wait()

    end := time.Since(start)
    fmt.Println("time taken:", end)

}

func channelsSelect() {
    c1 := make(chan string)
    c2 := make(chan string)
    c3 := make(chan string)

    time.Now()

    go func() {
        for {
            time.Sleep(time.Millisecond * 500)
            c1 <- "Every 500ms"
        }
    }()

    go func() {
        for {
            time.Sleep(time.Millisecond * 2000)
            c2 <- "Every 2 seconds"
        }
    }()

    go func() {
        for {
            time.Sleep(time.Second * 8)
            c3 <- "All proccesses finished"
        }
    }()


    for {
        select {
        case ms1 := <- c1:
            fmt.Println(ms1)
        case ms2 := <- c2:
            fmt.Println(ms2)
        case ms3 := <- c3:
            fmt.Println(ms3)
            os.Exit(0)
        }
    }
}
