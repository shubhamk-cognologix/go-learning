package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan bool)
	go func() {
		fmt.Println("series 1")
		time.Sleep(2 * time.Second)
		fmt.Println("End of routine")
		done <- true
	}()
	fmt.Println("series 2")
	result := <-done
	fmt.Println("End of main ", result)

	go count("sheep")
	go count("fish")
	time.Sleep(2 * time.Second)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("Sheep")
		wg.Done()
	}()
	fmt.Printf("In between\n")
	wg.Add(1)
	go func() {
		count("Fish")
		wg.Done()
	}()
	fmt.Printf("In main\n")
	wg.Wait()
	fmt.Printf("End of main\n")
}

func count(thing string) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d : %s\n", i, thing)
		time.Sleep(500 * time.Millisecond)
	}
}
