// +build OMIT

package main

import (
	"fmt"
	"time"
)

func sleepAndTalk(t time.Duration, msg string) {
	time.Sleep(t)
	fmt.Printf("%v ", msg)
}

func main() {
	go sleepAndTalk(0*time.Second, "Salut!")
	go sleepAndTalk(1*time.Second, "Ati")
	go sleepAndTalk(2*time.Second, "completat")
	go sleepAndTalk(3*time.Second, "Intervals?")
	time.Sleep(4 * time.Second)
}
