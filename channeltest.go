package main

import (
	"fmt"
	"time"
)

type ChannelTest struct{}

func (c *ChannelTest) Do() {
	msg := make(chan string)
	go func() {
		msg <- "Hello, world!"
	}()
	fmt.Println(<-msg)

	channelsync()

	channelDir()
}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true

}

func channelDir() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func channelsync() {

	done := make(chan bool, 1)

	go worker(done)

	<-done

}
