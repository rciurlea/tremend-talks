package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	d := newDispatcher()
	ids := idGenerator()
	s := make(chan os.Signal)
	signal.Notify(s, os.Interrupt)
	go func() {
		<-s
		log.Println("got ctrl-c, shutting down...")
		ln.Close()
	}()
	for {
		log.Println("waiting for connection")
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn, d, <-ids)
	}
}

func idGenerator() chan int {
	c := make(chan int)
	i := 0
	go func() {
		for {
			c <- i
			i++
		}
	}()
	return c
}

func handleConnection(conn net.Conn, d *dispatcher, id int) {
	nickname := fmt.Sprintf("user%d", id)
	done := make(chan bool)
	outgoing := connectionWriter(conn, done)
	outgoing <- fmt.Sprintf("Welcome, your nickname is %s, use /nick to change", nickname)
	d.register(outgoing)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "/nick") {
			words := strings.Split(line, " ")
			if len(words) != 2 {
				outgoing <- "wrong syntax. Supported commands /nick [nickname]"
			} else {
				nickname = words[1]
			}
		} else {
			d.dispatch(nickname + ": " + scanner.Text())
		}
	}
	close(done)
	log.Printf("client %v disconnected", conn.RemoteAddr())
}

func connectionWriter(conn net.Conn, done chan bool) chan string {
	c := make(chan string)
	go func() {
	Loop:
		for {
			select {
			case msg := <-c:
				fmt.Fprintf(conn, "%s\n", msg)
			case <-done:
				break Loop
			}
		}
	}()
	return c
}
