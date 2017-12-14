// +build OMIT

package main

import (
	"fmt"
	"time"
)

func sleepAndTalk(secs time.Duration, msg string, c chan string) {
	time.Sleep(secs * time.Second)
	c <- msg
}

func main() {
	c := make(chan string)

	go sleepAndTalk(0, "Salut!", c)
	go sleepAndTalk(1, "Ati", c)
	go sleepAndTalk(2, "completat", c)
	go sleepAndTalk(3, "Intervals?", c)

	for i := 0; i < 4; i++ {
		fmt.Printf("%v ", <-c)
	}
}
