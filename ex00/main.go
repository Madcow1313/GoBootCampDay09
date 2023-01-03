package main

import (
	"fmt"
	"time"
)

// func sleepSec(t int, number int, ch *chan int) {
// 	//defer wg.Done()
// }

func sleepSort(arr []int) chan int {
	ch := make(chan int, len(arr))
	for i := 0; i < len(arr); i++ {
		number := arr[i]
		t := arr[i]
		go func() {
			time.Sleep(time.Duration(t) * time.Second)
			ch <- number
		}()
	}
	return ch
}

func main() {
	arr := []int{5, 9, 8, 7, 4, 3, 6, 2, 1}
	ch := sleepSort(arr)
	for i := 0; i < cap(ch); i++ {
		val, ok := <-ch
		if !ok {
			break
		} else {
			fmt.Println(val)
		}
	}
}
