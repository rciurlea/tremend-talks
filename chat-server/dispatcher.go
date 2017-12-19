package main

import "sync"

type dispatcher struct {
	input   chan string
	outputs []chan string
	lock    sync.Mutex
}

func newDispatcher() *dispatcher {
	d := &dispatcher{}
	d.input = make(chan string)
	d.outputs = make([]chan string, 0)
	return d
}

func (d *dispatcher) register(c chan string) {
	d.lock.Lock()
	d.outputs = append(d.outputs, c)
	d.lock.Unlock()
}

func (d *dispatcher) dispatch(s string) {
	d.lock.Lock()
	for _, c := range d.outputs {
		c <- s
	}
	d.lock.Unlock()
}
