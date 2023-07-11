package concurrency

// func ConcurrencyMain() {
//     workersWithChannels()
// }

// UNCOMMENT BELOW & AT THE END OF FILE
/*

func weightGroups() {
    c := make(chan string)
    go count("sheep", c)

    msg := <- c
    fmt.Println(msg)
}

func channels() {
	c := make(chan string, 2)
	c <- "hello"
	c <- "I'm arturo"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)

}

func selectLogic() {
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
			c2 <- "Every 2 seconds"
			time.Sleep(time.Millisecond * 2000)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func workersWithChannels() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    go worker(jobs, results)
    go worker(jobs, results)
    go worker(jobs, results)
    go worker(jobs, results)

    for i := 0; i < 100; i++ {
        jobs <- i
    }
    close(jobs)

    for j := 0; j < 100; j++ {
        fmt.Println(<- results)
    }

}

func worker(jobs chan int, results chan int) {
    for n := range jobs {
        results <- fib(n)
    }
}

func fib(n int) int {
    if n <= 1 {
        return n
    }
    return fib(n - 1) + fib(n - 2)
}

func count(item string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- item
		time.Sleep(time.Millisecond + 500)
	}

	close(c)
}

*/
