package main

import (
	"fmt"
	"log"
	"net"
)

const listenAddr = "localhost:4000"

func main() {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println("Starting the Loop")
		c, err := l.Accept()
		fmt.Println("Connection established!!")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(c, "Hello yo!")
		c.Close()
	}
}
