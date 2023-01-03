package main

import (
	"sort"
	"testing"
)

func TestWithBufferedChannels(t *testing.T) {
	ch1 := make(chan interface{}, 5)
	ch2 := make(chan interface{}, 5)
	ch3 := make(chan interface{}, 5)
	for i := 0; i < cap(ch1); i++ {
		ch1 <- i
		ch2 <- i
		ch3 <- i
	}
	l := len(ch1) + len(ch2) + len(ch3)
	expected := []int{
		0, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4,
	}
	outputChan := multiplex(ch1, ch2, ch3)
	result := make([]int, 0)

	for i := 0; i < l; i++ {
		value, ok := <-outputChan
		if !ok {
			break
		} else {
			result = append(result, value.(int))
		}
	}
	if len(result) != len(expected) {
		t.Errorf("Wrong result")
	}
	sort.Ints(result)
	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected and result mismatch")
		}
	}
}
