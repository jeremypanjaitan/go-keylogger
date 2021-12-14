package main

import (
	"log"
	"net"
	"os"
)

func NewConnection() net.Conn {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return conn
}
