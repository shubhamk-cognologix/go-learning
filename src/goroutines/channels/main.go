package main

import (
	"fmt"
	"time"
)

func main() {
	// c := make(chan string)
	// go count("Sheep", c)
	// for msg := range c {
	// 	fmt.Println(msg)
	// }

	/*
		No of send should be equal to no of receive
	*/
	// c := make(chan string, 2)
	// c <- "Hello"
	// c <- "world"

	// msg := <-c
	// fmt.Println(msg)
	// msg = <-c
	// fmt.Println(msg)

	/*
		Select statement
	*/

	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		for {
			ch1 <- "Every 500 ms"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "Every 2 seconds"
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}

	/*
		worker
	*/
	// jobs := make(chan int, 100)
	// results := make(chan int, 100)

	// go workers(jobs, results)
	// go workers(jobs, results)
	// go workers(jobs, results)
	// go workers(jobs, results)

	// for i := 0; i < 100; i++ {
	// 	jobs <- i
	// }
	// close(jobs)
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(<-results)
	// }

}
func workers(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func count(thing string, c chan string) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d : %s\n", i, thing)
		c <- thing
		time.Sleep(500 * time.Millisecond)
	}
	close(c)
}
