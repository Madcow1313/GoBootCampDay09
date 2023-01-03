package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
)

var MAX_GOROUTINES int = 8

func crawlWeb(ctx context.Context, inputChannel chan string) chan *string {
	outputChannel := make(chan *string, len(inputChannel))
	guard := make(chan struct{}, MAX_GOROUTINES)
	l := len(inputChannel)
	for i := 0; i < l; i++ {
		select {
		case <-ctx.Done():
			defer close(outputChannel)
			return outputChannel
		default:
			guard <- struct{}{}
			go func(ctx context.Context) {
				fmt.Println("Start")
				resp, err := http.Get(<-inputChannel)
				if err == nil {
					body, err := io.ReadAll(resp.Body)
					defer resp.Body.Close()
					if err == nil {
						str := string(body)
						outputChannel <- &str
					} else {
						fmt.Println(err.Error())
						str := string("")
						outputChannel <- &str
					}
				} else {
					fmt.Println(err)
					str := string(err.Error())
					outputChannel <- &str
				}
				fmt.Println("finished")
				<-guard
			}(ctx)
		}
	}
	defer close(inputChannel)
	return outputChannel
}

func main() {
	urls := []string{
		"http://www.google.com",
		"http://www.google.com",
		"http://www.google.com",
		"http://www.google.com",
		"http://www.google.com",
		"http://www.google.com",
		"http://www.google.com",
		"http://www.google.com",
		"http://www.google.com",
		"http://www.google.com",
		"http://www.google.com",
		"http://www.ya.ru",
		"http://www.vk.com",
		"http://www.yahoo.com",
		"http://www.twitch.tv",
		"http://github.com",
		"http://leetcode.com",
		"http://www.youtube.com",
		"http://www.youtube.com",
		"http://www.youtube.com",
		"http://www.google.com",
		"http://www.google.com",
		"http://www.google.com",
		"http://www.google.com",
		"http://www.google.com",
		"http://www.google.com",
		"http://www.google.com",
		"http://www.google.com",
		"http://www.google.com",
		"http://www.google.com",
		"http://www.ya.ru",
		"http://www.vk.com",
		"http://www.yahoo.com",
		"http://www.twitch.tv",
		"http://github.com",
		"http://leetcode.com",
		"http://www.youtube.com",
		"http://www.youtube.com",
		"http://www.youtube.com",
		"http://shouldnothang",
	}
	l := len(urls)
	inputChannel := make(chan string, l)
	fmt.Println(l)
	for i := 0; i < l; i++ {
		inputChannel <- urls[i]
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ctxChannel := make(chan os.Signal, 1)
	signal.Notify(ctxChannel, os.Interrupt)
	defer func() {
		signal.Stop(ctxChannel)
		cancel()
	}()
	go func() {
		select {
		case <-ctxChannel:
			fmt.Println("ctrl-c")
			cancel()
			//os.Exit(0) antipattern
		case <-ctx.Done():
		}
	}()
	out := crawlWeb(ctx, inputChannel)
	for i := 0; i < l; i++ {
		val, ok := <-out
		if !ok {
			fmt.Println("Channel closed")
			break
		} else {
			fmt.Println(val)
			fmt.Println("processed", i)
		}
	}
	fmt.Println("Done!")
}
