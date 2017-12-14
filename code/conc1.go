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
	sleepAndTalk(0*time.Second, "Salut!")
	sleepAndTalk(1*time.Second, "Ati")
	sleepAndTalk(2*time.Second, "completat")
	sleepAndTalk(3*time.Second, "Intervals?")
}
