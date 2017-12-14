// +build OMIT

package main

import "fmt"

type Engine struct{}

func (e Engine) Start() {
	fmt.Println("Engine started")
}

func (e Engine) Stop() {
	fmt.Println("Engine stopped")
}

type Car struct {
	Engine // fără nume
}

func main() {
	var c Car

	c.Start()
	c.Stop()
}
