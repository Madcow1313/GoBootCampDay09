package main

import (
	"fmt"
	"sync"
)

func multiplex(arg ...chan interface{}) chan interface{} {
	output := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(len(arg))
	for _, ch := range arg {
		go func(ch chan interface{}) {
			for v := range ch {
				output <- v
			}
			wg.Done()
		}(ch)
	}
	go func() {
		wg.Wait()
		defer close(output)
	}()
	return output
}

func fillCh(ch chan interface{}, number int) {
	ch <- number
}

func main() {
	//channelLength := 10
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})
	ch4 := make(chan interface{})
	ch5 := make(chan interface{})
	// defer close(ch1)
	// defer close(ch2)
	// defer close(ch3)
	// defer close(ch4)
	// defer close(ch5)
	for i := 0; i < 10; i++ {
		n := i
		go fillCh(ch1, n)
		go fillCh(ch2, n)
		go fillCh(ch3, n)
		go fillCh(ch4, n)
	}
	for i := 0; i < 10; i++ {
		go func() {
			ch5 <- "Hooray"
		}()
	}
	out := multiplex(ch1, ch2, ch3, ch4, ch5)
	for i := 0; i < 50; i++ {
		val, ok := <-out
		if !ok {
			fmt.Println("Yep, it's closed")
			break
		} else {
			fmt.Println(val)
		}
	}
}
